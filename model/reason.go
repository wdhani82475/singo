package model

// Reason 原因模型
type Reason struct {
	Type        string
	Status      string
	Reason      string
	Id 			int
}

func (Reason) TableName() string {
	return "ad_check_reasons"
}