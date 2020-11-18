package like

type LikeService struct {
	PostUser        string `form:"post_user_id" json:"post_user_id" binding:"required"`
	LikeUser        string `form:"like_user_id" json:"like_user_id" binding:"required"`
	ArticleId        string `form:"article_id" json:"article_id" binding:"required"`
}




func (service *LikeService) DoLikeArticle() string {

	return  "1"
	//return  serializer.Response{}
}


//使用协程同步数据库