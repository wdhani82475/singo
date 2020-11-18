package like

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Nickname       string
	Status         string
	Avatar         string `gorm:"size:1000"`
}

func (b *User) tableName() string{
	return "like_user"
}