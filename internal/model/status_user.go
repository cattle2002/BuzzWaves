package model

import (
	"BuzzWaves/pkkg"
	"fmt"
)

// LoginStatusUE 通过up 记录用户登录状态
func LoginStatusUE(username string, email string) {
	u := User{}
	err := DB.Where("username = ? and email = ?", username, email).First(&u).Error
	if err != nil {
		fmt.Println(err)
	}
	ut := UserStatus{}
	ut.UserName = u.UserName
	ut.Email = u.Email
	ut.Status = "在线"
	err = DB.Create(&ut).Error
	fmt.Println(err)
}

//LoginStatusEP  通过ep 记录用户登录状态
func LoginStatusEP(email string, password string) {
	u := User{}
	err := DB.Where("email = ? and password = ?", email, pkkg.AesEncrypt(password)).First(&u).Error
	if err != nil {
		fmt.Println(err)
	}
	ut := UserStatus{}
	ut.UserName = u.UserName
	ut.Email = u.Email
	ut.Status = "在线"
	err = DB.Create(&ut).Error
	fmt.Println(err)
}
