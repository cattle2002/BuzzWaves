package wbsocket

import "C"
import (
	"BuzzWaves/internal/middleware"
	"BuzzWaves/internal/model"
	"BuzzWaves/pkkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有的跨域请求
		return true
	},
}

func WsHandlerDemo(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	_, err := pkkg.ValidateJWT(token)
	//if err != nil {
	//	c.JSON(404, gin.H{
	//		"msg": "token错误",
	//	})
	//}
	username, password := middleware.GetClaimsUP(token)
	_, err = middleware.QueryUserNamePassword(username, password)
	if err != nil {
		c.JSON(404, gin.H{
			"msg": "token错误",
		})
	}
	_, err = upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	get := c.Request.Header.Get("Authorization")
	fmt.Println(get)
	_, err2 := pkkg.ValidateJWT(get)
	if err2 != nil {
		fmt.Println(err2)
	}
	//todo 留个ep的接口出来
	up, s := middleware.GetClaimsUP(get)
	fmt.Println(up, s)

}

type OnlineMessageReq struct {
	FriendName string `yaml:"friendname"`
	Email      string `yaml:"email"`
}

//todo 做一个全局的WebsokcetConn
var WebSocketConns map[string]*websocket.Conn

func WsHandler(c *gin.Context) {

	username, exists := c.Get("username")
	if exists == true {
		fmt.Println("username", username)
	}
	email, e := c.Get("email")
	if e == true {
		fmt.Println("----------", email)
	}
	//todo 当用户每次登录之后 记录这条连接 key 为 username+email
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	//todo Websocket的string
	wname := username.(string) + ":" + email.(string)
	WebSocketConns[wname] = conn
	for key, _ := range WebSocketConns {
		// 打印键和值
		fmt.Println("Key:", key)
	}
	//todo 将登录的用户的状态改为在线
	model.LoginStatusUE(username.(string), email.(string))
}

// QueryUserNameEmail 通过查找记录是否存在 然后发送websocket消息
func QueryUserNameEmail(friendName string, email string) bool {
	us := model.UserStatus{}
	err := model.DB.Where("username = ? and email = ?", friendName, email).First(&us).Error
	if err != nil {
		fmt.Println("发送消息的好友不存在")
		return false
	}
	if us.Status == "在线" {
		return true
	} else {
		return false
	}
}

// WbSendMessageOnline  ddd
//todo  这个函数用来处理在线用户发送消息 前端携带在线好友的用户名 还是需要经过
func WbSendMessageOnline(c *gin.Context) {
	f := OnlineMessageReq{}
	c.ShouldBindJSON(&f)
	fmt.Println(f)
	email := QueryUserNameEmail(f.FriendName, f.Email)
	if email == true {
		WebSocketConns[f.FriendName+":"+f.Email].WriteMessage(websocket.TextMessage, []byte("hello,world"))
	} else {
		c.JSON(200, gin.H{
			"msg": "当前好友不在线",
		})
	}
}
