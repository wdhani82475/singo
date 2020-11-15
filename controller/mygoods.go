package controller

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func AddGoods(c *gin.Context) {
	var service service.GoodsService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
