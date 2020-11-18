package like

type UserModel struct {
	Id       int
	UserName string
	Status   int
}

func (UserModel) TableName() string {
	return "like_user"
}
