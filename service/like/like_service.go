package like

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"singo/cache"
	"singo/job"
	"singo/model"
	"singo/model/like"
	"singo/serializer"
	"strconv"
	"time"
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
func (service *LikeService)  sync() bool{
	//获取文章id、总数
	//user_id
	//article_id
	client := cache.RedisClient //client 对象
	data,err := client.Do("get",like.LIKE_ARTICLE_COUNT_PREFIX+strconv.Itoa(service.ArticleId)).Result()
	if err != nil {
		panic(err)
	}
	var article1 like.ArticleModel
	if err := gorm.DB.Where("id =?",service.ArticleId).First(&article1).Update("total_like_count",data).Error;err != nil {
		panic(err)
	}


}


//2.文章点赞总数+1
func (service *LikeService) DoLikeArticle() *serializer.Response {
	//校验用户是否点过👍

	res := service.sync()
	if res != true {
		panic(res)
	}
	//更新数据库


	if err := service.valid(); err != nil {
		return err
	}
	//写入redis缓存
	client := cache.RedisClient //client 对象
	//用户点赞文章列表
	if err := client.Do("sadd", like.USER_LIKE_ARTICLE_PREFIX+strconv.Itoa(service.PostUser), service.ArticleId).Err(); err != nil {
		panic(err)
	}

	//文章点赞总数+1
	if err1 := client.Do("Incr", like.LIKE_ARTICLE_COUNT_PREFIX+strconv.Itoa(service.ArticleId)).Err(); err1 != nil {
		panic(err1)
	}

	//文章被哪些用户点过赞
	if err3 := client.Do("sadd", like.LIKE_ARTICLE_HAS_USER_PREFIX+strconv.Itoa(service.ArticleId), service.PostUser).Err(); err3 != nil {
		panic(err3)
	}



	// v1:mysql
	userLikeArticle := like.UserLikeArticleModel{
		PostUserId: service.PostUser,
		LikeUserId: service.LikeUser,
		ArticleId:  service.ArticleId,
		IsLike:     like.DO_LIKE,
	}
	//开启事务
	tx := model.DB.Begin()

	//var res model.Reason
	// model.DB.Model(&like.UserLikeArticleModel{})
	//if err := model.DB.Where("id = ?",id ).First(&res).Error; err != nil {}
	//用户点赞关联表+1
	//能查找到则更新状态
	var userlike like.UserLikeArticleModel
	if err := tx.Where("post_user_id = ? and like_user_id = ? and article_id = ?", service.PostUser, service.LikeUser, service.ArticleId).First(&userlike).Error; err != nil {
		tx.Rollback()
		return &serializer.Response{
			Code: 400,
			Msg:  "查询数据失败",
		}
	}
	//fmt.Printf("userLike:%#v\n",userlike)
	//找不到则插入
	if &userlike != nil {
		if err := tx.Where("post_user_id = ? and like_user_id = ? and article_id = ?", service.PostUser, service.LikeUser, service.ArticleId).First(&userlike).Update("is_like", like.DO_LIKE).Error; err != nil {
			tx.Rollback()
			return &serializer.Response{
				Code: 400,
				Msg:  "用户点赞失败1",
			}
		}
	} else {
		//&Toy{Name: "Stuffed Animal", OwnerType: "Nobody"} 1.定义对象 2.直接构造对象
		if err2 := tx.Model(&like.UserLikeArticleModel{}).Create(&userLikeArticle).Error; err2 != nil {
			tx.Rollback()
			return &serializer.Response{
				Code: 400,
				Msg:  "用户点赞失败2",
			}
		}
	}
	//更新数据
	var article like.ArticleModel
	if err4 := tx.Where("id = ?", service.ArticleId).First(&article).Update("total_like_count", article.TotalLikeCount+1).Error; err4 != nil {
		tx.Rollback()
		return &serializer.Response{
			Code: 400,
			Msg:  "更新数据失败",
		}
	}
	tx.Commit()
	return &serializer.Response{
		Code: 200,
		Msg:  "操作成功",
		Data: "",
	}
}

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

func (service *LikeService) Sync() {
	log.Println("Starting......")
	c := job.Crontab
	c.AddFunc("* */3 * * * ?", func() {
		log.Println("Run sync data......")
		service.syncData2Mysql()
	})
	c.Start()  //启动
	t1 := time.NewTimer(time.Minute*1)  //新增定时器
	for { 								// for+select 阻塞select 等待channel
		select {
		case <-t1.C:
			t1.Reset(time.Minute*1)   // 重置定时器，重新计数
		}
	}
}



//同步数据到数据表
func (service *LikeService) syncData2Mysql() {
	//使用协程同步数据库V2
	fmt.Println(1234)
}
