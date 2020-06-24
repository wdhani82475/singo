package serializer

import "singo/model"

// User 用户序列化器
type Reason struct {
	Type   string `json:"type"`
	Status string `json:"status"`
	Id     int `json:"id"`
	Reason string `json:"reason"`
	//CreateDate int64  `json:"create_date"`
	//UpdateDate int64  `json:"update_date"`
}

// BuildUser 序列化原因
func BuildReason(res model.Reason) Reason {
	return Reason{
		Id:     res.Id,
		Reason: res.Reason,
		Type:   res.Type,
		Status: res.Status,
		//CreateDate: res.CreateDate.Unix(),
		//UpdateDate: res.UpdateDate.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildReasonResponse(res model.Reason) Response {
	return Response{
		Data: BuildReason(res),
	}
}
