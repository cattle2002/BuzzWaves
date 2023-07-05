package middleware

import (
	v1 "BuzzWaves/api/v1"
	"BuzzWaves/pkkg"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

//SplitToken 返回token
func SplitToken(Authorization string) string {
	parts := strings.Split(Authorization, "Bearer ")
	if len(parts) == 2 {
		//bearer := parts[0]
		token := parts[1]
		return token
	} else {
		fmt.Println("Invalid token format")
		return ""
	}
}
func GetClaimsUP(token string) (string, string) {
	jwt, err := pkkg.ValidateJWT(token)
	if err != nil {
		return "", ""
	}
	//todo 用户登录有两种方式通过用户名和密码      通过密码和邮箱
	//todo 根据断言的成功与否来进行下面操作
	if password, ok := jwt["password"].(string); ok {
		if username, ok := jwt["username"].(string); ok {
			return username, password
		}
	} else {
		if email, ok := jwt["email"].(string); ok {
			return email, password
		}
	}
	return "", ""
}
func endsWithCom(str string) bool {
	return strings.HasSuffix(str, ".com")
}

// JwtMiddleWare todo 在设置token的过期时间的时候，我们统一把用户名作为key
// JwtMiddleWare 定义一个自定义的中间件
func JwtMiddleWare() gin.HandlerFunc {
	fmt.Println("-------")
	GetCLi()
	return func(c *gin.Context) {
		// 中间件逻辑处理
		Authorization := c.Request.Header.Get("Authorization")
		token := SplitToken(Authorization)
		if token != "" {
			//todo 验证token是否有效（token是否被篡改）
			_, err2 := pkkg.ValidateJWT(token)
			if err2 != nil {
				c.JSON(200, gin.H{
					"msg": "无效的token",
				})
				c.Abort()
			}
			//todo !!! 其实每次登录的时候都应该去数据查看用户是否存在  再进行续期等操作
			//todo 判断登录方式是up 还是 ep 根据url 来进行判断、
			if c.Request.URL.String() == "/api/v1/user/name/login" {
				//up
				username, password := GetClaimsUP(token)
				fmt.Println(username, password)
				//todo 从redis去查询token
				value := GetRedisValue(username)
				if value == token {
					//todo 没登陆一次重新设置token过期时间,以此进行续期
					SetRedisExpireTime(username, token)
					c.Next()
				} else {
					//todo 如果token过期 redis里面不存在信息  先去数据库里面查询用户是否存在，如果存在 则续期，如果不存在 则  Abort
					SetRedisExpireTime(username, token)
				}
			}
			if c.Request.URL.String() == "/api/v1/user/email/login" {
				//todo ep 进行登录 取出email 和 password  去数据查询出用户名，然后去redis去查询是否有该用户名的token 如果没有则设置redis token
				email, password := GetClaimsUP(token)
				fmt.Println("heeeeee", email, password)
				user, _ := QueryUserEmailPassword(email, password)
				value := GetRedisValue(user.UserName)
				if value == "" {
					SetRedisExpireTime(user.UserName, token)
				} else {
					//todo 可能过期
					SetRedisExpireTime(user.UserName, token)
				}
				c.Next()
			}
			//todo 拿到token之后去Redis里面查看token是否过期，如果token过期 （取数据里面查找用户是否存在，如果存在继续续期.否则完事）则继续续期
			RdbM.Get(context.Background(), "")
			c.Next()
		} else {
			//todo 用户第一次登录，没有token 需要去数据库里面查找是否用户是否存在 最后颁发token 判断登录方式
			url := c.Request.URL.String()
			fmt.Println(url)
			if url == "/api/v1/user/name/login" {
				//todo username password进行登录
				up := v1.UserLoginUPReq{}
				//ep := Ep{}
				err := c.ShouldBindJSON(&up)
				fmt.Println(err, up)
				password, err := QueryUserNamePassword(up.UserName, up.Password)
				fmt.Println(password)
				//todo 将token保存在Redis里面 过期时间为24h
				jwtup, _ := pkkg.GenerateJWTUP(up.UserName, up.Password)
				SetRedisExpireTime(up.UserName, jwtup)
				c.JSON(200, gin.H{
					"msg":   "登录初次登录成功,请保存好token",
					"token": jwtup,
				})
			} else {
				//todo 用户第一次登录通过email password进行登录
				ep := v1.UserLoginEPReq{}
				err := c.ShouldBindJSON(&ep)
				fmt.Println(err, ep)
				password, err := QueryUserEmailPassword(ep.Email, ep.Password)
				fmt.Println(password)
				//todo 将token保存在Redis里面 过期时间为24h
				jwtep, _ := pkkg.GenerateJWTEP(ep.Email, ep.Password)
				//todo
				SetRedisExpireTime(password.UserName, jwtep)
				c.JSON(200, gin.H{
					"msg":   "登录初次登录成功,请保存好token",
					"token": jwtep,
				})

			}
			//todo 通过邮箱和密码进行登录
			//todo 通过用户名和密码进行登录
			c.Abort()

		}

		//if {
		//	//1.数据库是否存在 如果不存在直接abort
		//	//2.数据如果存在,没有token 返回token next 设置token在redis的过期时间
		//	//3.数据存在，有 token
		//}
		//// 继续处理请求

	}
}
