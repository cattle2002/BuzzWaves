package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"testing"
)

type Config1 struct {
	Author    map[string]interface{} `yaml:"Author"`
	Log       map[string]interface{} `yaml:"Log"`
	Mysql     map[string]interface{} `yaml:"Mysql"`
	Redis     map[string]interface{} `yaml:"Redis"`
	SqlLite   map[string]interface{} `yaml:"SqlLite"`
	Websocket map[string]interface{} `yaml:"Websocket"`
}

func TestConfig(t *testing.T) {
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
	var config Config1
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("解码 YAML 数据失败: %v", err)
	}
	// 访问配置数据
	fmt.Println("Author:", config.Author)
	m, ok := Config.Author["Age"].(int)
	fmt.Println(m, ok)
	fmt.Println("Log:", config.Log)
	fmt.Println("Mysql:", config.Mysql)
	fmt.Println("Redis:", config.Redis)
	fmt.Println("SqlLite:", config.SqlLite)
	fmt.Println("Websocket:", config.Websocket)
}
