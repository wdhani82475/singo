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
//根据id获取详情
func GetReason(c *gin.Context){
	//获取参数的方式
	////get
	id1 := c.Query("id") //查询请求URL后面的参数
	id2 := c.DefaultQuery("id", "100") //查询URL后面的参数，如果没有填写默认值
	id3 := c.PostForm("id") //从表单中查询参数
	fmt.Println(id1,id2,id3)
	//post
	name1 := c.Request.FormValue("name")
	name2 := c.Request.PostFormValue("name")
	name3 := c.DefaultPostForm("name","haha")
	name4,_ := c.GetPostForm("name")
	fmt.Println(name1,name2,name3,name4)

	fmt.Println("bing.....")

	id := c.Query("id")
	var service service.ReasonService
	res :=service.Get(id)
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