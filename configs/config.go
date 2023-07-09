package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type ServerConfig struct {
	Author    map[string]interface{} `yaml:"Author"`
	Log       map[string]interface{} `yaml:"Log"`
	Mysql     map[string]interface{} `yaml:"Mysql"`
	Redis     map[string]interface{} `yaml:"Redis"`
	SqlLite   map[string]interface{} `yaml:"SqlLite"`
	Websocket map[string]interface{} `yaml:"Websocket"`
	BuzzWaves map[string]interface{} `yaml:"BuzzWaves"`
	AesKey    map[string]interface{} `yaml:"AesKey"`
	Jwt       map[string]interface{} `yaml:"Jwt"`
	Rabbitmq  map[string]interface{} `yaml:"Rabbitmq"`
	Minio     map[string]interface{} `yaml:"Minio"`
}

var Config ServerConfig

//var Config SernverConfig

//配置信息读取到Config 当中
func init() {
	// 打开 YAML 配置文件
	file, err := os.Open("../configs/config.yaml")
	if err != nil {
		log.Fatalf("无法打开配置文件: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// 创建解码器
	decoder := yaml.NewDecoder(file)

	// 解码 YAML 数据到结构体

	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatalf("解码 YAML 数据失败: %v", err)
	}
}
func GetAuthorInfo() string {
	//fmt.Println(config.Config)
	m, ok := Config.Author["Author"].(string)
	if !ok {
		panic(errors.New("获取Author信息失败"))
	} else {
		return m
	}
}
