package biz

import (
	v1 "BuzzWaves/api/v1"
	"BuzzWaves/internal/model"
	"BuzzWaves/internal/model/rbt"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func UserRegister(c *gin.Context) {
	var buzzReq v1.UserRegisterReq
	_ = c.ShouldBindJSON(&buzzReq)
	model.WriteRemoteLog(c.Request.RemoteAddr + "  " + "新用户注册")
	fmt.Println(buzzReq)
	model.RegisterUser(&buzzReq)
	c.JSON(200, gin.H{"message": "注册用户成功"})
}

func UserLoginUp(c *gin.Context) {

	//Authorization := c.Request.Header.Get("Authorization")
	//fmt.Println(Authorization)
	var buzzReq v1.UserLoginUPReq
	_ = c.ShouldBindJSON(&buzzReq)
	fmt.Println("req", buzzReq)
	//fmt.Println(buzzReq)
	//s, _ := c.Get("token")
	//s2 := s.(string)
	//username, _ := c.Get("username")
	//usernamestr := username.(string)
	//todo 每一次用户登录之后 向UserStatus这张表写入数据 专门记录用户是否登录 XXXX
	//model.LoginStatusUP(buzzReq.UserName, buzzReq.Password)
	//todo 这里应该 从队列里面取出离线消息 (用户登录之后先从biz_user这张表查询出自己的user_id 然后拿到所有好友的username)
	userid := model.GetOwnerUserDFromBizUser(buzzReq.UserName)
	fmt.Println(userid)
	more := model.FindOneToMore(userid)
	fmt.Println("-------------------------------------", more)
	//todo 先从好友查出来好友列表 然后依次取出消息
	userNameList := make([]string, 0)
	for i := range more {
		userNameList = append(userNameList, more[i].UserName)
		fmt.Println("我要获取消息了")
		ch, _ := rbt.ConsumeMessage(rbt.RabbitChannel, more[i].UserName+":"+buzzReq.UserName)
		go func(ch <-chan amqp.Delivery) {
			for delivery := range ch {
				fmt.Println(string(delivery.Body))
			}
		}(ch)
	}

	c.JSON(200, gin.H{
		"msg": "登录成功111",
		//"Token": s2,
	})
}
func UserLoginEp(c *gin.Context) {
	//todo 每一次用户登录之后 向UserStatus这张表写入数据 专门记录用户是否登录
	var buzzReq v1.UserLoginEPReq
	_ = c.ShouldBindJSON(&buzzReq)
	fmt.Println(buzzReq)
	//todo 每一次用户登录之后 向UserStatus这张表写入数据 专门记录用户是否登录 XXXX
	//model.LoginStatusEP(buzzReq.Email, buzzReq.Password)
	c.JSON(200, gin.H{
		"msg": "登录成功",
	})
}

// AddFriend AddUser 添加用户好友到 业务的表 friends里面
func AddFriend(c *gin.Context) {
	req := v1.FriendAddReq{}
	c.ShouldBindJSON(&req)
	fmt.Println(req)
	model.AddFriendBiz(req.OwnerName, req.Username, req.Email)
}

func AddGroup(c *gin.Context) {
	var gn v1.CreateGroupReq
	c.ShouldBindJSON(&gn)
	fmt.Println(gn)
}
func GetMessage(queueName []string, ownername string) {
	// 创建一个上下文对象和取消函数
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 启动协程
	go func() {
		// 模拟协程的工作
		for {
			select {
			case <-ctx.Done():
				// 上下文被取消，结束协程
				log.Println("Coroutine stopped")
				return
			default:
				// 协程的工作逻辑
				for i := range queueName {
					rbt.ConsumeMessage(rbt.RabbitChannel, queueName[i]+":"+ownername)

				}

			}
		}
	}()

	// 等待上下文超时或被取消
	select {
	case <-ctx.Done():
		log.Println("Context timeout or cancelled")
	}
}
