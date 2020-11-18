package like

type UserLikeArticleModel struct {
	LikeUserId int
	PostUserId int
	ArticleId  int
}

func (UserLikeArticleModel) TableName() string {
	return "user_like_article"
}
