package pkkg

import (
	config "BuzzWaves/configs"
	"errors"
	"fmt"
)

func GetBuzzWavesIP() string {
	m, ok := config.Config.BuzzWaves["IP"].(string)
	if !ok {
		panic(errors.New("获取BuzzWaves信息失败"))
	} else {
		return m
	}
}
func GetBuzzWavesPort() int {
	m, ok := config.Config.BuzzWaves["Port"].(int)
	if !ok {
		panic(errors.New("获取BuzzWaves信息失败"))
	} else {
		return m
	}
}

// GetAuthorInfo GetAuthInfo 获取作者姓名
func GetAuthorInfo() string {
	fmt.Println(config.Config)
	m, ok := config.Config.Author["Author"].(string)
	if !ok {
		panic(errors.New("获取Author信息失败"))
	} else {
		return m
	}
}

// GetAgeInfo GetAgeInfo 获取作者年龄
func GetAgeInfo() int {
	m, ok := config.Config.Author["Age"].(int)
	if !ok {
		panic(errors.New("获取Age信息失败"))
	} else {
		return m
	}
}

//GetEmailInfo GetEmailInfo 获取邮箱
func GetEmailInfo() string {
	m, ok := config.Config.Author["Email"].(string)
	if !ok {
		panic(errors.New("获取Email信息失败"))
	} else {
		return m
	}
}

//GetWebsocketOwnerIPInfo GetWebsocketOwnerIPInfo 获取Websocket通信地址
func GetWebsocketOwnerIPInfo() string {
	m, ok := config.Config.Websocket["OwnerIP"].(string)
	if !ok {
		panic(errors.New("获取配置信息失败"))
	} else {
		return m
	}
}

//GetWebsocketPORTInfo GetWebsocketPORTInfo 获取websocket通信端口
func GetWebsocketPORTInfo() int {
	m, ok := config.Config.Websocket["PORT"].(int)
	if !ok {
		panic(errors.New("获取配置信息失败"))
	} else {
		return m
	}
}

//GetMysqlIPInfo  获取Mysql地址
func GetMysqlIPInfo() string {
	m, ok := config.Config.Mysql["IP"].(string)
	if !ok {
		panic(errors.New("获取配置信息失败"))
	} else {
		return m
	}
}

// GetMysqlPORTInfo 获取mysql端口号
func GetMysqlPORTInfo() int {
	m, ok := config.Config.Mysql["PORT"].(int)
	if !ok {
		panic(errors.New("获取配置信息失败"))
	} else {
		return m
	}
}

//GetMysqlUserInfo  获取Mysql 用户
func GetMysqlUserInfo() string {
	m, ok := config.Config.Mysql["User"].(string)
	if !ok {
		panic(errors.New("获取配置信息失败"))
	} else {
		return m
	}
}

// GetMysqlPasswordInfo 获取密码
func GetMysqlPasswordInfo() int {
	m, ok := config.Config.Mysql["Password"].(int)
	if !ok {
		panic(errors.New("获取配置信息失败"))
	} else {
		return m
	}
}

func GetRedisIpInfo() string {
	m, ok := config.Config.Redis["IP"].(string)
	if !ok {
		panic(errors.New("获取配置信息失败"))
	} else {
		return m
	}
}
func GetRedisPortInfo() int {
	m, ok := config.Config.Redis["PORT"].(int)
	if !ok {
		panic(errors.New("获取配置信息失败"))
	} else {
		return m
	}
}

func GetAesKey() string {
	m, ok := config.Config.AesKey["Key"].(string)
	if !ok {
		panic(errors.New("获取配置信息失败"))
	} else {
		return m
	}
}

func GetJwtKey() string {
	m, ok := config.Config.Jwt["Key"].(string)
	if !ok {
		panic(errors.New("获取配置信息失败"))
	} else {
		return m
	}
}

func GetJwtKeyExpire() int {
	m, ok := config.Config.Jwt["Expire"].(int)
	if !ok {
		panic(errors.New("获取配置信息失败"))
	} else {
		return m
	}
}

//RabbitMq:
//IP: 127.0.0.1
//PORT: 5672
//User: admin
//Password: 123
func GetRabbitmqIP() string {
	m, ok := config.Config.Rabbitmq["IP"].(string)
	if !ok {
		panic(errors.New("获取配置信息失败"))
	} else {
		return m
	}
}

func GetRabbitmqPort() int {
	m, ok := config.Config.Rabbitmq["PORT"].(int)
	if !ok {
		panic(errors.New("获取配置信息失败"))
	} else {
		return m
	}
}

func GetRabbitmqUser() string {
	m, ok := config.Config.Rabbitmq["User"].(string)
	if !ok {
		panic(errors.New("获取配置信息失败"))
	} else {
		return m
	}
}

func GetRabbitmqPassword() int {
	m, ok := config.Config.Rabbitmq["Password"].(int)
	if !ok {
		panic(errors.New("获取配置信息失败"))
	} else {
		return m
	}
}

// //# SqlLite
// //SqlLite:
// //IP: localhost
// //PORT: 3306
// //User: root
// //Password: 123456
// //Application: 客户管理程序 数据监管程序 证书监管程序
// func GetSqlLiteIPInfo() string {
// 	m, ok := config.Config["SqlLite"].(map[string]interface{})
// 	if !ok {
// 		panic(errors.New("获取配置信息失败"))
// 	} else {
// 		return m["IP"].(string)
// 	}
// }
// func GetSqlLitePORTInfo() int {
// 	m, ok := config.Config["SqlLite"].(map[string]interface{})
// 	if !ok {
// 		panic(errors.New("获取配置信息失败"))
// 	} else {
// 		return m["PORT"].(int)
// 	}
// }
// func GetSqlLiteUserInfo() string {
// 	m, ok := config.Config["SqlLite"].(map[string]interface{})
// 	if !ok {
// 		panic(errors.New("获取配置信息失败"))
// 	} else {
// 		return m["User"].(string)
// 	}
// }
// func GetSqlLitePasswordInfo() int {
// 	m, ok := config.Config["SqlLite"].(map[string]interface{})
// 	if !ok {
// 		panic(errors.New("获取配置信息失败"))
// 	} else {
// 		return m["Password"].(int)
// 	}
// }
// func GetSqlLiteApplicationInfo() string {
// 	m, ok := config.Config["SqlLite"].(map[string]interface{})
// 	if !ok {
// 		panic(errors.New("获取配置信息失败"))
// 	} else {
// 		return m["Application"].(string)
// 	}
// }

// func GetLogFilePosition() string {
// 	m, ok := config.Config["Log"].(map[string]interface{})
// 	if !ok {
// 		panic(errors.New("获取配置信息失败"))
// 	} else {
// 		return m["File"].(string)
// 	}
// }
