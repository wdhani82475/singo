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

func (service *ReasonService) Get() serializer.Response {
	var res model.Reason
	if err := model.DB.Where("reason = ?","测试一下" ).First(&res).Error; err != nil {
		return serializer.ParamErr("创建失败", err)
	}
	//fmt.Println(reflect.TypeOf(res))
	return serializer.BuildReasonResponse(res)
}