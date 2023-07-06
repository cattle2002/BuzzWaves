package model

import (
	v1 "BuzzWaves/api/v1"
	"errors"
	"fmt"
)

func CopyUserToBizUser(user *v1.UserRegisterReq) *BizUser {
	return &BizUser{
		UserName:  user.UserName,
		TellPhone: user.TellPhone,
		Email:     user.Email,
		Address:   user.Address,
		Captcha:   user.Captcha,
		Gender:    user.Gender,
	}
}

// AddUserToBizUser 将注册的用户添加到biz_user这张表 相比于user表少了一些隐私字段
func AddUserToBizUser(user *v1.UserRegisterReq) {
	bizUser := CopyUserToBizUser(user)
	err := DB.Create(bizUser).Error
	if err != nil {
		fmt.Println("用户注册添加到BizUser失败", err)
	}
}

// Friend 用户好友的存储表 作为BizUser的子表
//type Friend struct {
//	gorm.Model
//	UserName  string `json:"username" gorm:"type:varchar(50);column:username"`
//	TellPhone string `json:"telephone" gorm:"type:varchar(20);column:telephone"`
//	Email     string `json:"email" gorm:"type:varchar(50);column:email"`
//	Address   string `json:"address" gorm:"type:varchar(50);column:address"`
//	Captcha   string `json:"captcha" gorm:"type:varchar(50);column:captcha"`
//	Gender    string `json:"gender" gorm:"type:varchar(5);column:gender"`
//	Status    string `yaml:"status" gorm:"type:varchar(5);column:status"` //存储好友是否在线
//	BizUserID uint   `gorm:"foreignKey:UserID"`
//}
//todo 通过用户名和邮箱将好友查询出来然后进行添加
//todo 有个缺陷  我们强制去biz_user 里面去查询用户信息 然后进行添加好友信息吧 【也就是一定要用户登录了Websocket 才能添加好友】
func AddFriendBiz(ownername string, username string, email string) error {
	ou := BizUser{}
	u := User{}
	//todo 先将想要添加用户的用户id拿出来d 然后进行添加数据
	err := DB.Where("username = ?", ownername).First(&ou).Error
	if err != nil {
		return errors.New("前端传递的用户名不对")
	}
	//fmt.Println("用户数据", ou)
	//todo 查找好友的数据
	err = DB.Where("username = ? and  email = ?", username, email).First(&u).Error
	if err != nil {
		fmt.Println("通过username email 查询错误失败", err)
	}
	//todo 添加数据到biz friends 表 先将
	fmt.Println("friend", u)
	fridInfo := Friend{
		UserName:  u.UserName,
		TellPhone: u.TellPhone,
		Email:     u.Email,
		Address:   u.Address,
		Captcha:   u.Captcha,
		Gender:    u.Gender,
		BizUserID: ou.ID,
	}
	tx := DB.Create(&fridInfo).Error
	fmt.Println(tx)

	return nil
}
