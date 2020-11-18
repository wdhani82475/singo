package like

type ArticleModel struct {
	Id             int
	UserId         int
	ArticleName    string
	ArticleContent string
	TotalLikeCount int
	Status         int
}

func (ArticleModel) tableName() string {
	return "article"
}
