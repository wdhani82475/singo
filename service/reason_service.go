package service

import (
	"singo/model"
	"singo/serializer"
)

// UserLoginService 管理用户登录的服务
type ReasonService struct {
	Reason string `form:"reason" json:"reason" binding:"required,min=2,max=30"`
	Status string `form:"status" json:"status" binding:"required,min=1"`
	Type   string `form:"type" json:"type" binding:"required,min=1"`
}

func (service *ReasonService) Create() serializer.Response {
	//参数
	reason := model.Reason{
		Type: service.Type,
		Reason:service.Reason,
		Status: service.Status,
	}

	// 创建原因
	if err := model.DB.Create(&reason).Error; err != nil {
		return serializer.ParamErr("创建失败", err)
	}
	return serializer.BuildReasonResponse(reason)
}
