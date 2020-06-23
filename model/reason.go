package model

// Reason 原因模型
type Reason struct {
	//gorm.Model
	//CreateDate  time.Time
	//UpdateDate time.Time
	Type        string
	Status      string
	Reason      string
}
