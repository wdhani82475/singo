package like

type UserLikeArticleModel struct {
	LikeUserId int
	PostUserId int
	ArticleId  int
	IsLike     int
}

const DO_LIKE = 1
const DIS_LIKE = 0

const USER_LIKE_ARTICLE_PREFIX = "user_like_article_prefix_" //用户点赞哪些文章

const LIKE_ARTICLE_COUNT_PREFIX = "like_article_count_prefix_" //文章点赞的总数

const LIKE_ARTICLE_HAS_USER_PREFIX = "like_article_user_prefix_" // 文章被那些用户点过赞

func (UserLikeArticleModel) TableName() string {
	return "user_like_article"
}
