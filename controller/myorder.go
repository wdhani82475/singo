package controller

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func AddOrder(ctx *gin.Context) {
	var service  service.OrderServer
	if err := ctx.ShouldBind(&service); err == nil {
		res := service.CreateAndUpdateStock()
		ctx.JSON(200,res)
	}else{
		ctx.JSON(-1,ErrorResponse(err))
	}
}

