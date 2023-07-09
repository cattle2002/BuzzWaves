package main

import (
	"BuzzWaves/internal/model/buzzMinio"
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
	//model.GetCLi()
	//endpoint := pkkg.GetMinioEndpoint()
	//fmt.Println(endpoint)
	//id := pkkg.GetMinioAccessKeyID()
	//fmt.Println(id)
	//
	//key := pkkg.GetMinioSecretAccessKey()
	//fmt.Println(key)
	//
	//ssl := pkkg.GetMinioSecretUseSSL()
	//fmt.Println(ssl)
	//
	//name := pkkg.GetMinioSecretBucketName()
	//fmt.Println(name)
	buzzMinio.GetMinioClient()
	//func UploadFile(uploadLocalPosition string, BucketName string, contentType string, objectName string) {

	//upload.UploadFile("C:\\Users\\gongzhaowei\\Documents\\TencentMeeting\\2023-07-08 15.55.08 gonegone的快速会议 706578818\\meeting_01.mp4",
	//	"hello",
	//	"video/mp4",
	//	"gonevideo2",
	//)
	//download.DownloadFile("hello", "/gonevideo", "D:\\gocode\\BuzzWaves\\internal\\model\\buzzMinio\\download\\test2.mp4")
	wbsocket.WebSocketConns = make(map[string]*websocket.Conn, 5)
	buzz := server.NewServer()
	buzz.Run()
}
