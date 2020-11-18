package like

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"singo/conf"
	"singo/serializer"
	"singo/service/like"
)



//点赞
func LikeArticle(c *gin.Context) {
	//参数校验
	var service like.LikeService
	if err := c.ShouldBind(&service);err != nil {
		res := service.DoLikeArticle()
		c.JSON(200,res)
	}else{
		c.JSON(200, ErrorResponse(err))
	}
	//返回结果值

}

//取消点赞
func  DisLikeArticle(c*gin.Context)  {

}



func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", e.Field))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return serializer.ParamErr(
				fmt.Sprintf("%s%s", field, tag),
				err,
			)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.ParamErr("JSON类型不匹配", err)
	}

	return serializer.ParamErr("参数错误", err)
}