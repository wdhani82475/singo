package service

import (
	"singo/model"
	"singo/serializer"
)

// UserLoginService 管理用户登录的服务
type GoodsService struct {
	Id    int `form:"id" json:"id" binding:"required"`
	Stock int `form:"stock" json:"stock" binding:"required"`
}

func (service *GoodsService) Create() serializer.Response {
	//参数
	goods := model.Order{
		Id:    service.Id,
		Stock: service.Stock,
	}

	// 创建原因
	if err := model.DB.Create(&goods).Error; err != nil {
		return serializer.ParamErr("创建失败", err)
	}
	return serializer.Response{}
}

