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
		c.JSON(200,gin.H{
			"code":200,
			"msg":"ok",
			"data":data,
		})
	}else{
		c.JSON(200, gin.H{
			"code":200,
			"msg":"ok",
			"data":err,
		})
	}
}

//取消点赞
func  DisLikeArticle(c*gin.Context)  {

}