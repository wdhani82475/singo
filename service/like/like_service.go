package like

import (
	"singo/model"
	"singo/model/like"
	"singo/serializer"
)

type LikeService struct {
	PostUser  int `form:"post_user_id" json:"post_user_id" binding:"required"`
	LikeUser  int `form:"like_user_id" json:"like_user_id" binding:"required"`
	ArticleId int `form:"article_id" json:"article_id" binding:"required"`
}

//1.检测用户是否已经点赞过该文章
//model.DB.Model(&like.UserLikeArticleModel{}).First(&userlike).Where("post_user_id = ? and like_user_id = ? and article_id = ?", service.PostUser, service.LikeUser, service.ArticleId).Count(&count)
func (service *LikeService) valid() *serializer.Response {
	count := 0
	model.DB.Model(&like.UserLikeArticleModel{}).Where("post_user_id = ? and like_user_id = ? and article_id = ? and is_like = ? ", service.PostUser, service.LikeUser, service.ArticleId, like.DO_LIKE).Count(&count)
	//fmt.Println(count)
	if count > 0 {
		return &serializer.Response{
			Code: 400,
			Msg:  "您已经点过赞",
		}
	}

	return nil
}

//2.文章点赞总数+1
func (service *LikeService) DoLikeArticle() *serializer.Response {
	//校验用户是否点过👍
	if err := service.valid(); err != nil {
		return err
	}

	userLikeArticle := like.UserLikeArticleModel{
		PostUserId: service.PostUser,
		LikeUserId: service.LikeUser,
		ArticleId:  service.ArticleId,
		IsLike:     like.DO_LIKE,
	}

	//var res model.Reason
	// model.DB.Model(&like.UserLikeArticleModel{})
	//if err := model.DB.Where("id = ?",id ).First(&res).Error; err != nil {}
	//TODO
	//用户点赞关联表+1
	//能查找到则更新状态
	var userlike like.UserLikeArticleModel
	if err := model.DB.Where("post_user_id = ? and like_user_id = ? and article_id = ?", service.PostUser, service.LikeUser, service.ArticleId).First(&userlike).Error; err != nil {
		return &serializer.Response{
			Code: 400,
			Msg:  "查询数据失败",
		}
	}
	//fmt.Printf("userLike:%#v\n",userlike)
	//找不到则插入
	if &userlike != nil {
		if err := model.DB.Where("post_user_id = ? and like_user_id = ? and article_id = ?", service.PostUser, service.LikeUser, service.ArticleId).First(&userlike).Update("is_like", like.DO_LIKE).Error; err != nil {
			return &serializer.Response{
				Code: 400,
				Msg:  "用户点赞失败1",
			}
		}
	} else {
		if err2 := model.DB.Model(&like.UserLikeArticleModel{}).Create(&userLikeArticle).Error; err2 != nil {
			return &serializer.Response{
				Code: 400,
				Msg:  "用户点赞失败2",
			}
		}
	}
	//更新数据
	var article like.ArticleModel
	if err4 := model.DB.Where("id = ?", service.ArticleId).First(&article).Update("total_like_count", article.TotalLikeCount+1).Error; err4 != nil {
		return &serializer.Response{
			Code: 400,
			Msg:  "更新数据失败",
		}
	}
	return &serializer.Response{
		Code: 200,
		Msg:  "操作成功",
		Data: "",
	}
}

//使用协程同步数据库V2

func (service *LikeService) isHasLike() *serializer.Response {
	var userLike like.UserLikeArticleModel
	if err := model.DB.Where("post_user_id = ? and like_user_id = ? and article_id = ? and is_like = ?", service.PostUser, service.LikeUser, service.ArticleId, like.DO_LIKE).First(&userLike).Error; err != nil {
		return &serializer.Response{
			Code: 400,
			Msg:  "获取数据失败",
		}
	}
	if &userLike != nil {
		return nil
	} else {
		return &serializer.Response{
			Code: 400,
			Msg:  "数据不存在",
		}
	}

}

func (service *LikeService) DisLikeArticle() *serializer.Response {
	//取消点赞 1.判断用户书否已经点过赞，如果没有点赞则不能取消点赞 是否有效
	if err := service.isHasLike(); err != nil {
		return err
	}
	tx := model.DB.Begin()
	//2. 用户取消点赞
	if err2 := tx.Model(like.UserLikeArticleModel{}).
		Where("post_user_id = ? and like_user_id = ? and article_id = ? ", service.PostUser, service.LikeUser, service.ArticleId).
		Update("is_like", like.DIS_LIKE).Error; err2 != nil {
		tx.Rollback()
		return &serializer.Response{
			Code: 400,
			Msg:  "用户取消点赞失败",
		}
	}
	//3.将对应的文章的点赞总数-1
	var article like.ArticleModel
	if err3 := tx.
		Where("id = ? and user_id = ? and total_like_count > 0", service.ArticleId, service.LikeUser).First(&article).
		Update("total_like_count", article.TotalLikeCount-1).Error; err3 != nil {
		tx.Rollback()
		return &serializer.Response{
			Code: 400,
			Msg:  "点赞数减一失败",
		}
	}
	tx.Commit()
	return &serializer.Response{
		Code: 200,
		Msg:  "取消点赞成功",
		Data: "",
	}
}
