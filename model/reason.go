package model

import (
	"github.com/jinzhu/gorm"
)


// User 用户模型
type Reason struct {
	gorm.Model
	Type    string
	Status  string
	Avatar  string `gorm:"size:1000"`
	Reason  string
}
