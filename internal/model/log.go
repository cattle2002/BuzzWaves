package model

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func WriteRemoteLog(msg string) {
	//position := pkg.GetLogFilePosition()
	file, err := os.OpenFile("D:\\gocode\\BuzzWaves\\internal\\model\\log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("无法创建日志文件：", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(errors.New("关闭文件失败"))
		}
	}(file)
	// 设置日志输出到文件
	log.SetOutput(file)
	log.SetPrefix("Info:")
	// 设置日志前缀和日志标志
	log.SetFlags(log.Ldate | log.Ltime)

	// 记录日志
	log.Println(msg)
}
func WriteLoginLogInInfo(ip string, port int, username string) {
	file, err := os.OpenFile("D:\\gocode\\BuzzWaves\\internal\\model\\info.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("无法创建日志文件：", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(errors.New("关闭文件失败"))
		}
	}(file)
	// 设置日志输出到文件
	log.SetOutput(file)
	log.SetPrefix("Info:")
	// 设置日志前缀和日志标志
	log.SetFlags(log.Ldate | log.Ltime)
	msg := fmt.Sprintf("远程IP地址：%s ， 远程端口地址：%d ,登录的用户名:%s", ip, port, username)
	// 记录日志
	log.Println(msg)
}
func WriteErrorLog(errs string) {
	file, err := os.OpenFile("D:\\gocode\\BuzzWaves\\internal\\model\\error.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("无法创建日志文件：", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(errors.New("关闭文件失败"))
		}
	}(file)
	// 设置日志输出到文件
	log.SetOutput(file)
	log.SetPrefix("Error:")
	// 设置日志前缀和日志标志
	log.SetFlags(log.Ldate | log.Ltime)
	//msg := fmt.Sprintf("应用出现巨大错误:%s", err.Error())
	//fmt.Println(msg)
	//s := err.Error()
	// 记录日志
	log.Println(errs)
}
func WriteSqlError(msg string) {
	file, err := os.OpenFile("D:\\gocode\\BuzzWaves\\internal\\model\\sqlerror.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("无法创建日志文件：", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(errors.New("关闭文件失败"))
		}
	}(file)
	// 设置日志输出到文件
	log.SetOutput(file)
	log.SetPrefix("Error:")
	// 设置日志前缀和日志标志
	log.SetFlags(log.Ldate | log.Ltime)
	//msg := fmt.Sprintf("应用出现巨大错误:%s", err.Error())
	//fmt.Println(msg)
	//s := err.Error()
	// 记录日志
	log.Println(msg)
}
