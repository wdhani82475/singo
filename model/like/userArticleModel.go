package like

type UserLikeArticleModel struct {
	LikeUserId int
	PostUserId int
	ArticleId  int
	IsLike     int
}

const DO_LIKE = 1
const DIS_LIKE = 0

func (UserLikeArticleModel) TableName() string {
	return "user_like_article"
}
