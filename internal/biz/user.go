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
	Authorization := c.Request.Header.Get("Authorization")
	fmt.Println(Authorization)
	var buzzReq v1.UserLoginUPReq
	_ = c.ShouldBindJSON(&buzzReq)
	fmt.Println(buzzReq)
	c.JSON(200, gin.H{
		"msg": "登录成功111",
	})
}
func UserLoginEp(c *gin.Context) {

	c.JSON(200, gin.H{
		"msg": "登录成功",
	})
}
