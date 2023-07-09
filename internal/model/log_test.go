package model

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	fmt.Println("---------------------")
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

	// 设置日志前缀和日志标志
	log.SetFlags(log.Ldate | log.Ltime)

	// 记录日志
	log.Println("hello,world")
}
