package middleware

import (
	"BuzzWaves/internal/model"
	"BuzzWaves/pkkg"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strconv"
	"time"
)

var DBM *gorm.DB
var err error
var RdbM *redis.Client

func GetCLi() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)", pkkg.GetMysqlUserInfo(), strconv.Itoa(pkkg.GetMysqlPasswordInfo()), pkkg.GetMysqlIPInfo(), pkkg.GetMysqlPORTInfo())
	dsn = dsn + "/BuzzWaves?charset=utf8&parseTime=True&loc=Local"
	DBM, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default})
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}
	// 创建 Redis 客户端实例

	RdbM = redis.NewClient(&redis.Options{
		Addr:     pkkg.GetRedisIpInfo() + ":" + strconv.Itoa(pkkg.GetRedisPortInfo()), // Redis 服务器地址和端口号
		Password: "",                                                                  // Redis 密码（如果有）
		DB:       0,                                                                   // Redis 数据库索引
	})

	// 检查是否成功连接到 Redis
	pong, err := RdbM.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Failed to connect to Redis:", err)
		return
	}

	fmt.Println("Connected to Redis:", pong)
}

type Up struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
type Ep struct {
	Email    string `yaml:"email"`
	password string `yaml:"password"`
}

func QueryUserNamePassword(username, password string) (*model.User, error) {
	ep := model.User{}
	passwd := pkkg.AesEncrypt(password)
	err := DBM.Where("username = ? and password = ?", username, passwd).First(&ep).Error
	//fmt.Println(ep)
	return &ep, err
}
func QueryUserEmailPassword(email, password string) (*model.User, error) {
	ep := model.User{}
	passwd := pkkg.AesEncrypt(password)
	err := DBM.Where("email = ? and password = ?", email, passwd).First(&ep).Error
	fmt.Println(ep)
	return &ep, err
}
func SetRedisExpireTime(username string, token string) {
	//expire :=time.Duration( pkkg.GetJwtKeyExpire())
	RdbM.Set(context.Background(), username, token, time.Duration(pkkg.GetJwtKeyExpire())*time.Second)
}
func GetRedisValue(username string) string {
	value, err := RdbM.Get(context.Background(), username).Result()
	if err == redis.Nil {
		fmt.Println("键不存在")
	} else if err != nil {
		fmt.Println("获取值失败:", err)
	} else {
		fmt.Println("键的值:", value)
	}
	return value
}
