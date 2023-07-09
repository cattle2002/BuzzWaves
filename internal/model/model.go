package model

import (
	"BuzzWaves/pkkg"
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strconv"
)

// User 用户表 后续不会使用到 只在登录 注册的业务逻辑进行用到
type User struct {
	gorm.Model
	//UserName     string `json:"username" gorm:"username" `
	UserName     string `json:"username" gorm:"type:varchar(50);column:username"`
	PassWord     string `json:"password" gorm:"type:varchar(255);column:password"`
	RegisterTime string `json:"registered" gorm:"type:datetime;column:registered"`
	TellPhone    string `json:"telephone" gorm:"type:varchar(20);column:telephone"`
	Email        string `json:"email" gorm:"type:varchar(50);column:email"`
	Address      string `json:"address" gorm:"type:varchar(50);column:address"`
	Captcha      string `json:"captcha" gorm:"type:varchar(50);column:captcha"`
	Gender       string `json:"gender" gorm:"type:varchar(5);column:gender"`
}

// BizUser biz User 用来进行添加好友 作为friens的主表(将用户注册表的部分字段进行过滤)
type BizUser struct {
	gorm.Model
	UserName  string `json:"username" gorm:"type:varchar(50);column:username"`
	TellPhone string `json:"telephone" gorm:"type:varchar(20);column:telephone"`
	Email     string `json:"email" gorm:"type:varchar(50);column:email"`
	Address   string `json:"address" gorm:"type:varchar(50);column:address"`
	Captcha   string `json:"captcha" gorm:"type:varchar(50);column:captcha"`
	Gender    string `json:"gender" gorm:"type:varchar(5);column:gender"`
	Friends   []Friend
}

// Friend 用户好友的存储表 作为BizUser的子表
type Friend struct {
	gorm.Model
	UserName  string `json:"username" gorm:"type:varchar(50);column:username"`
	TellPhone string `json:"telephone" gorm:"type:varchar(20);column:telephone"`
	Email     string `json:"email" gorm:"type:varchar(50);column:email"`
	Address   string `json:"address" gorm:"type:varchar(50);column:address"`
	Captcha   string `json:"captcha" gorm:"type:varchar(50);column:captcha"`
	Gender    string `json:"gender" gorm:"type:varchar(5);column:gender"`
	//Status    string `yaml:"status" gorm:"type:varchar(5);column:status"` //存储好友是否在线
	BizUserID uint `gorm:"foreignKey:UserID"`
}
type UserStatus struct {
	UserName string `json:"username" gorm:"type:varchar(50);column:username"`
	Status   string `json:"status" gorm:"type:varchar(20);column:status"`
	Email    string `json:"email" gorm:"type:varchar(50);column:email"`
}
type Group struct {
	gorm.Model
	OwnerUserName string      `json:"ownerusername" gorm:"type:varchar(20);column:ownerusername"`
	Name          string      `json:"groupname" gorm:"type:varchar(20);column:groupname"`
	GroupUsers    []GroupUser `gorm:"foreignKey:GroupID"`
}

type GroupUser struct {
	gorm.Model
	UserName  string `json:"username" gorm:"type:varchar(50);column:username"`
	TellPhone string `json:"telephone" gorm:"type:varchar(20);column:telephone"`
	Email     string `json:"email" gorm:"type:varchar(50);column:email"`
	Address   string `json:"address" gorm:"type:varchar(50);column:address"`
	Captcha   string `json:"captcha" gorm:"type:varchar(50);column:captcha"`
	Gender    string `json:"gender" gorm:"type:varchar(5);column:gender"`
	GroupID   uint   `gorm:"column:GroupID"`
}

var DB *gorm.DB
var err error
var Rdb *redis.Client

func GetCLi() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)", pkkg.GetMysqlUserInfo(), strconv.Itoa(pkkg.GetMysqlPasswordInfo()), pkkg.GetMysqlIPInfo(), pkkg.GetMysqlPORTInfo())
	dsn = dsn + "/BuzzWaves?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&Group{}, &GroupUser{})
	// 创建 Redis 客户端实例
	Rdb = redis.NewClient(&redis.Options{
		Addr:     pkkg.GetRedisIpInfo() + ":" + strconv.Itoa(pkkg.GetRedisPortInfo()), // Redis 服务器地址和端口号
		Password: "",                                                                  // Redis 密码（如果有）
		DB:       0,                                                                   // Redis 数据库索引
	})

	// 检查是否成功连接到 Redis
	pong, err := Rdb.Ping().Result()
	if err != nil {
		fmt.Println("Failed to connect to Redis:", err)
		return
	}

	fmt.Println("Connected to Redis:", pong)
}

func (User) TableName() string {
	return "user_register_reqs"
}

//func (Groups) TableName() string {
//	return "groups"
//}
//func (Groups) GroupUser() string {
//	return "group_user"
//}
