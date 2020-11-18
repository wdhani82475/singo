package like

import (
	"fmt"
	"singo/model"
	"singo/model/like"
	"singo/serializer"
)

type LikeService struct {
	PostUser        int `form:"post_user_id" json:"post_user_id" binding:"required"`
	LikeUser        int `form:"like_user_id" json:"like_user_id" binding:"required"`
	ArticleId        int `form:"article_id" json:"article_id" binding:"required"`
}

//1.检测用户是否已经点赞过该文章
func(service *LikeService) valid() *serializer.Response{
	count := 0
	model.DB.Model(&like.UserLikeArticleModel{}).Where("post_user_id = ? and like_user_id = ? and article_id = ?  ", service.PostUser,service.LikeUser,service.ArticleId).Count(&count)
	fmt.Println(count)
	if count>0 {
		return &serializer.Response{
			Code: 404,
			Msg:  "您已点过赞",
		}
	}
	return nil
}



//2.文章点赞总数+1
func (service *LikeService) DoLikeArticle() *serializer.Response {
	//校验用户是否点过👍
	err := service.valid()
	if err != nil {
		return err
	}

	userLikeArticle := like.UserLikeArticleModel{
		PostUserId: service.PostUser,
		LikeUserId: service.LikeUser,
		ArticleId: service.ArticleId,
	}
	//用户点赞关联表+1
	err2 := model.DB.Model(&like.UserLikeArticleModel{}).Create(&userLikeArticle).Error
	if err2 != nil {
		return &serializer.Response{
			Code: 400,
			Msg:  "插入失败",
		}
	}

	//3.被点赞、点赞用户、文章+1
	//err = db.Model(&Toy{}).Where(&Toy{OwnerType: "Martian"}).Update("OwnerType", "Astronaut").Error
	model.DB.Model(&like.ArticleModel{}).Where("id = ?",service.ArticleId).Update("total_like_count")
	//更新逻辑
}


//使用协程同步数据库