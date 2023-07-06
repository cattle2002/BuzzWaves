package biz

import "C"
import (
	"BuzzWaves/internal/middleware"
	"BuzzWaves/pkkg"
	"fmt"
	"github.com/gin-gonic/gin"
)

func WebsocketMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取token

		token := c.Request.Header.Get("Authorization")
		_, err := pkkg.ValidateJWT(token)
		if err != nil {
			fmt.Println("--------1")
			c.Abort()
		} else {
			//todo 整个middleWare留点ep 过过期时间什么的 这里应该是用户登录之后的 去redis去查找用户是否存在
			//todo 如果存在则放行  如果不存在返回websocket错误
			username, password := middleware.GetClaimsUP(token)
			u, err := middleware.QueryUserNamePassword(username, password)
			if err != nil {
				fmt.Println(err)
				c.Abort()
			} else {
				//fmt.Println("--------2")
				c.Set("username", username)
				c.Set("email", u.Email)
				c.Next()
			}
			//fmt.Println("--------3")

			c.Abort()
		}
	}
}
