package like

import (
	"github.com/gin-gonic/gin"
	"singo/service/like"
)



//点赞
func LikeArticle(c *gin.Context) {
	//参数校验
	var service like.LikeService

	if err := c.ShouldBind(&service);err == nil {
		data := service.DoLikeArticle()
		c.JSON(200,data)
	}else{
		c.JSON(400,err)
	}
}

//取消点赞
func  DisLikeArticle(c*gin.Context)  {
	var service like.LikeService
	if err := c.ShouldBind(&service);err != nil {
		//data := service.DisLikeArticle()
		c.JSON(200,1)
	}else{
		c.JSON(400,err)
	}
}