package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/service"
)

func CreateReason(c *gin.Context) {
	var service service.ReasonService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func GetReason(c *gin.Context){
	var service service.ReasonService
	res :=service.Get()
	c.JSON(200,res)
}


func GetReasonV1(c *gin.Context) {
	var res model.Reason
	if err := model.DB.Where("reason = ?","测试一下" ).First(&res).Error; err != nil {
		fmt.Println("获取数据失败")
	}
	//fmt.Println(res)
	c.JSON(200, res)
}