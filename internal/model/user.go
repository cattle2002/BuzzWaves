package model

import (
	v1 "BuzzWaves/api/v1"
	"BuzzWaves/pkkg"
)

func CopyReq(u *v1.UserRegisterReq) User {
	newUser := User{
		UserName: u.UserName,
		PassWord: u.PassWord,
		//RegisterTime: "2023-07-04 12:34:56",
		RegisterTime: u.RegisterTime,
		TellPhone:    u.TellPhone,
		Email:        u.Email,
		Address:      u.Address,
		Captcha:      u.Captcha,
		Gender:       u.Gender,
	}
	return newUser
}
func init() {
	GetCLi()
}
func RegisterUser(user *v1.UserRegisterReq) {
	req := CopyReq(user)
	req.PassWord = pkkg.AesEncrypt(req.PassWord)
	tx := DB.Create(&req).Debug()
	if tx.Error != nil {
		WriteSqlError("插入用户失败" + tx.Error.Error())
	}
	WriteRemoteLog("访问数据用户表User,添加用户")
	//todo 用户注册的时候 需要往 业务用户表插入一条数据
	AddUserToBizUser(user)

}
func UserLoginUp(up *v1.UserLoginUPReq) {

}
