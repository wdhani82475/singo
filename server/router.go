package server

import (
	"fmt"
	"os"
	"singo/api"
	"singo/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		//测试一下
		v1.GET("/", api.Get)
		v1.GET("/check",api.Check)
		v1.POST("/ping", api.Ping)
		// 用户注册
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}
	}
	v2 := r.Group("/api/v2")
	{
		v2.POST("/add",api.CreateReason)
		v2.GET("/get",api.GetReason)
		v2.POST("/del",api.DelReason)
		v2.POST("/update", api.UpdateReason)
		//restful样式请求
		v2.POST("/post/:uuid", func(c *gin.Context){
			uuid := c.Param("uuid")
			fmt.Println(uuid)
			c.JSON(200,uuid)
		})
	}
	return r
}
