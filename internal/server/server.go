package server

import (
	"BuzzWaves/internal/biz"
	"BuzzWaves/internal/middleware"
	"BuzzWaves/pkkg"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BuzzWaves struct {
	Engine      *gin.Engine
	IP          string
	Port        int
	Application string
}

func NewServer() *BuzzWaves {
	return &BuzzWaves{
		Engine:      gin.Default(),
		IP:          pkkg.GetBuzzWavesIP(),
		Port:        pkkg.GetBuzzWavesPort(),
		Application: "BuzzWaves Application",
	}
}
func (buzz *BuzzWaves) BuzzBind() {
	buzz.Engine.POST("/api/v1/user/register", biz.UserRegister)
	v1 := buzz.Engine.Group("/api/v1/user")
	{

		v1.Use(middleware.JwtMiddleWare())
		v1.POST("/name/login", biz.UserLoginUp)
		v1.POST("/email/login", biz.UserLoginEp)
	}

}
func (buzz *BuzzWaves) Run() {
	gin.SetMode(gin.ReleaseMode)
	fmt.Println("BuzzWaves 正在启动...")
	buzz.BuzzBind()
	err := buzz.Engine.Run(buzz.IP + ":" + strconv.Itoa(buzz.Port))
	if err != nil {
		//model.WriteErrorLog("BuzzWaves 启动失败" + err.Error())
		panic(err)
	}
}

//var BuzzWavesEngine *gin.Engine
//
//func NewBuzzWavesServer() {
//	gin.SetMode(gin.ReleaseMode)
//	BuzzWavesEngine = gin.Default()
//	BuzzBind()
//	fmt.Printf("BuzzWaves is running at %s : %d", pkkg.GetBuzzWavesIP(), pkkg.GetBuzzWavesPort())
//}
//func Run(){
//	fmt.Println("BuzzWaves 正在启动...")
//	err := BuzzWavesEngine.Run(pkkg.GetBuzzWavesIP() + ":" + strconv.Itoa(pkkg.GetBuzzWavesPort()))
//	if err != nil{
//		fmt.Println("BuzzWaves Run Failed")
//		model.WriteErrorLog("启动应用失败"+err.Error())
//	}
//}
