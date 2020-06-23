package model

import (
	"time"
)

// Reason 原因模型
type Reason struct {
	//gorm.Model
	CreateDate  time.Time
	UpdatedDate time.Time
	Type        string
	Status      string
	Reason      string
}
