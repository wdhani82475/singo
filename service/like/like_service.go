package like

import (
	"fmt"
	"singo/model"
	"singo/model/like"
)

type LikeService struct {
	PostUser  int `form:"post_user_id" json:"post_user_id" binding:"required"`
	LikeUser  int `form:"like_user_id" json:"like_user_id" binding:"required"`
	ArticleId int `form:"article_id" json:"article_id" binding:"required"`
}

//1.检测用户是否已经点赞过该文章
func (service *LikeService) valid() string {
	count := 0
	model.DB.Model(&like.UserLikeArticleModel{}).Where("post_user_id = ? and like_user_id = ? and article_id = ?  ", service.PostUser, service.LikeUser, service.ArticleId).Count(&count)
	fmt.Println(count)
	if count > 0 {
		return "您已点过赞"
	}
	return ""
}

//2.文章点赞总数+1
func (service *LikeService) DoLikeArticle() string {
	//校验用户是否点过👍
	err := service.valid()
	if err != "" {
		return err
	}

	userLikeArticle := like.UserLikeArticleModel{
		PostUserId: service.PostUser,
		LikeUserId: service.LikeUser,
		ArticleId:  service.ArticleId,
	}
	//用户点赞关联表+1
	err2 := model.DB.Model(&like.UserLikeArticleModel{}).Create(&userLikeArticle).Error
	if err2 != nil {
		return "插入数据失败"
	}
	//更新数据
	var article like.ArticleModel
	err4 := model.DB.Where("id = ?", service.ArticleId).First(&article).Update("total_like_count", article.TotalLikeCount+1).Error
	if err4 != nil {
		return "更新数据失败"
	}
	return ""
}
//使用协程同步数据库
