package service

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"singo/model"
	"singo/serializer"
)

// UserLoginService 管理用户登录的服务

type OrderServer struct {
	GoodsId int `form:"goodsId" json:"goodsId" binding:"required"`
}

func (service *OrderServer) CreateAndUpdateStock() serializer.Response {
	//下订单
	client:=redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
		Password:"",
		DB:0,
	})

	pong,err:=client.Ping().Result()



	tx := model.DB.Begin()
	//如何开启一个事务

	uuid := 123

	order := model.Order{
		GoodsId: service.GoodsId,
		Uid:     uuid,
	}

	if err := tx.Create(&order).Error; err != nil {
		if err := tx.Rollback().Error; err != nil {
			_ = fmt.Errorf("Rollback should not raise error")
		}
		return serializer.ParamErr("用户下单失败", err)
	}
	//减少库存,按照商品id查找，并且物品的stock大于0时，才去更新库存
	var goods model.GoodsModel
	if err := tx.Where("id =? and stock>0", service.GoodsId).First(&goods).Update("stock", goods.Stock-1).Error; err != nil {
		if err := tx.Rollback().Error; err != nil {
			_ = fmt.Errorf("Rollback should not raise error")
		}
		return serializer.ParamErr("更新数据失败", err)
	}
	if err := tx.Commit().Error; err != nil {
		_ = fmt.Errorf("Commit should not raise error")
	}
	return serializer.Response{}

}
