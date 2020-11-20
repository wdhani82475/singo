package server

import (
	"os"
	"singo/api"
	"singo/api/like"
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
		v1.POST("ping", api.Ping)

		// 用户登录
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
	v2 := r.Group("/api/v2/")
	{
		//测试
		v2.GET("ping",api.Ping)
		v2.GET("crontab",like.Sync)  //定时执行
		v2.POST("do-like",like.LikeArticle)  //点赞
		v2.POST("do-dislike",like.DisLikeArticle)  //取消点赞
		//v2.GET("user-like-article",like.UserLikeArticleList)  //用户点赞文章列表
		//v2.GET("article-like-count",like.ArticleLikeCount)  //文章点赞总数
		//v2.GET("article-like-user",like.ArticleLikeUserList) //点赞文章的用户列表
	}
	return r
}

