package like

type UserLikeArticleModel struct {
	LikeUserId int
	PostUserId int
	ArticleId  int
	Status     int
}

func (UserLikeArticleModel) tableName() string {
	return "user_like_article"
}
