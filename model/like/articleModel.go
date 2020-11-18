package like

type ArticleModel struct {
	Id             int
	UserId         int
	ArticleName    string
	ArticleContent string
	TotalLikeCount int
}

func (ArticleModel) TableName() string {
	return "article"
}
