package like

type ArticleModel struct {
	Id             int
	UserId         int
	ArticleName    string
	ArticleContent string
	TotalLikeCount int
}

func (ArticleModel) tableName() string {
	return "article"
}
