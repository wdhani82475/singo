package like

type UserModel struct {
	Id       int
	UserName string
	Status   int
}

func (UserModel) tableName() string {
	return "like_user"
}
