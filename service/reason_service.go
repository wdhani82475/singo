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

// 根据id获取详情
func (service *ReasonService) Get(id interface{}) serializer.Response {
	var res model.Reason
	if err := model.DB.Where("id = ?",id ).First(&res).Error; err != nil {
		return serializer.ParamErr("获取数据失败", err)
	}
	return serializer.BuildReasonResponse(res)
}

//删除某一行记录
func (service *ReasonService) DelReason(id interface{}) serializer.Response {
	var res model.Reason
	if err := model.DB.Delete(&model.Reason{},id).First(&res).Error; err != nil {
		return serializer.ParamErr("获取数据失败", err)
	}
	return serializer.BuildReasonResponse(res)
}


//更新
func (service * ReasonService) UpdateReason(reason interface{}) serializer.Response {
	var res model.Reason

	if err := model.DB.Where("reason =?",reason).First(&res).Update("reason","tt").Error;err != nil  {
		return serializer.ParamErr("获取数据失败", err)
	}
	return serializer.BuildReasonResponse(res)
}