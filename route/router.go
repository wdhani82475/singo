package route

import (
	"fmt"
	"os"
	"singo/api"
	"singo/controller"
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
		v1.GET("/", controller.Get)
		v1.GET("/check", controller.Check)
		v1.POST("/ping", controller.Ping)
		// 用户注册
		v1.POST("user/register", controller.UserRegister)

		// 用户登录
		v1.POST("user/login", controller.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", controller.UserMe)
			auth.DELETE("user/logout", controller.UserLogout)
		}
	}
	v2 := r.Group("/api/v2")
	{
		v2.POST("/add", controller.CreateReason)
		v2.GET("/get", controller.GetReason)
		v2.POST("/del", controller.DelReason)
		v2.POST("/update", controller.UpdateReason)
		//restful样式请求
		v2.POST("/post/:uuid", func(c *gin.Context){
			uuid := c.Param("uuid")
			fmt.Println(uuid)
			c.JSON(200,uuid)
		})
	}
	v3 := r.Group("/api/v3")
	{
		v3.POST("/add/orderv2", controller.AddOrder)
		v3.POST("/add/order", controller.AddOrderV1)
		v3.POST("/create/goods", controller.AddGoods)
	}
	return r
}
