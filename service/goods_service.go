package service

import (
	"github.com/go-redis/redis"
	"singo/model"
	"singo/serializer"
)

// UserLoginService 管理用户登录的服务
type GoodsService struct {
	//Id    int `form:"id" json:"id" binding:"required"`
	Stock int `form:"stock" json:"stock" binding:"required"`
	Price int `form:"price" json:"price" binding:"required"`
}

func (service *GoodsService) Create() serializer.Response {
	//参数
	goods := model.GoodsModel{
	//	Id:    service.Id,
		Stock: service.Stock,
		Price: service.Price,
	}
	// 创建原因
	if err := model.DB.Create(&goods).Error; err != nil {
		return serializer.ParamErr("创建商品失败失败", err)
	}
	return serializer.Response{}
}

func (service *GoodsService) OpRedis() *serializer.Response  {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
	})

	if err := client.Do("Set", "qqq", 100);err != nil {
		return &serializer.Response{
			Code: 400,
			Msg: "写入数据失败",
		}
	}
	 data := client.Do("Get", "qqq")


	return &serializer.Response{
		Code: 200,
		Data: data,
	}
}