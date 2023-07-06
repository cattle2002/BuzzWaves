package biz

import (
	v1 "BuzzWaves/api/v1"
	"BuzzWaves/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var buzzReq v1.UserRegisterReq
	_ = c.ShouldBindJSON(&buzzReq)
	model.WriteRemoteLog(c.Request.RemoteAddr + "  " + "新用户注册")
	fmt.Println(buzzReq)
	model.RegisterUser(&buzzReq)
	c.JSON(200, gin.H{"message": "注册用户成功"})
}

func UserLoginUp(c *gin.Context) {
	//Authorization := c.Request.Header.Get("Authorization")
	//fmt.Println(Authorization)
	var buzzReq v1.UserLoginUPReq
	_ = c.ShouldBindJSON(&buzzReq)
	//fmt.Println(buzzReq)
	//todo 每一次用户登录之后 向UserStatus这张表写入数据 专门记录用户是否登录 XXXX
	//model.LoginStatusUP(buzzReq.UserName, buzzReq.Password)
	c.JSON(200, gin.H{
		"msg": "登录成功111",
	})
}
func UserLoginEp(c *gin.Context) {
	//todo 每一次用户登录之后 向UserStatus这张表写入数据 专门记录用户是否登录
	var buzzReq v1.UserLoginEPReq
	_ = c.ShouldBindJSON(&buzzReq)
	fmt.Println(buzzReq)
	//todo 每一次用户登录之后 向UserStatus这张表写入数据 专门记录用户是否登录 XXXX
	//model.LoginStatusEP(buzzReq.Email, buzzReq.Password)
	c.JSON(200, gin.H{
		"msg": "登录成功",
	})
}

// AddFriend AddUser 添加用户好友到 业务的表 friends里面
func AddFriend(c *gin.Context) {
	req := v1.FriendAddReq{}
	c.ShouldBindJSON(&req)
	fmt.Println(req)
	model.AddFriendBiz(req.OwnerName, req.Username, req.Email)
}
