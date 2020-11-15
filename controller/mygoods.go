package controller

import (
	"singo/service"
)

func AddGoods(c *Context) {
	//参数校验，绑定参数
	var service service.GoodsService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(CodeParamError, ErrorResponse(err))
	}
}
