package routers

import (
	"loginimpl/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	//注册相关
	r.GET("/register", controller.RegisterIndexController)
	r.POST("/register", controller.PostController)
	//登录相关
	r.GET("/login", controller.LoginIndexController)
	r.POST("/login", controller.GetController)
	//修改密码相关
	r.GET("/modify", controller.ModifyIndexController)
	r.POST("/modify", controller.PutController)

	r.GET("/delete/:username", controller.DELETEController)
	return r
}
