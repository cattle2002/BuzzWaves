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

	if err != nil {
		panic(err)
	}
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
