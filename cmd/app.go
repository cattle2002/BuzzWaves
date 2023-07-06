package main

import (
	"BuzzWaves/internal/server"
	"BuzzWaves/internal/wbsocket"
	"github.com/gorilla/websocket"
)

func main() {
	// server.NewBuzzWavesServer()
	//info := config.GetAuthorInfo()
	//fmt.Println(info)
	// fmt.Println(name)
	//info := pkkg.GetMysqlPasswordInfo()
	//fmt.Println(info)
	//e := errors.New("巨大错误")
	//fmt.Println(e.Error())
	//model.WriteErrorLog(e.Error())
	//err := pkkg.RsaGenKey(2048)
	//key := pkkg.GetAesKey()
	//encrypt, err := pkkg.AesEncrypt("123456", []byte(key))
	//fmt.Println(err)
	//fmt.Println(string(encrypt))
	//decrypt, err := pkkg.AesDecrypt(encrypt, []byte(key))
	//fmt.Println(err)
	//fmt.Println(decrypt)
	//fmt.Println(key)
	//key := pkkg.GetJwtKey()
	//fmt.Println(key)
	//jwt, err := pkkg.GenerateJWT("miaowed", "128@qq.com", "123456")
	//fmt.Println(err)
	//fmt.Println(jwt)

	//validateJWT, err := pkkg.ValidateJWT(jwt)
	//fmt.Println(err)
	//fmt.Println(validateJWT)
	//middleware.GetCLi()
	//middleware.GetRedisValue("xiaowei1")
	//fmt.Println(value)
	//fmt.Println("---------")

	//return "", ""
	//fmt.Println("hh", up, s)
	wbsocket.WebSocketConns = make(map[string]*websocket.Conn, 5)
	buzz := server.NewServer()
	buzz.Run()
}
