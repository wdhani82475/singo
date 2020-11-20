package conf

import (
	"os"
	"singo/cache"
	"singo/job"
	"singo/model"
	"singo/util"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	pwd := "/Users/wangdong/GolandProjects/src/singo/";
	// 从本地读取环境变量
	godotenv.Load(pwd+"env/.env.local")
	//godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales(pwd+"conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN")) //加载MySQL配置  使用  model.DB.Where()
	cache.Redis() //加载redis对象 cache.RedisClient.Incr()
	job.GetCron() //加载定时任务对象  job.Crontab.AddFunc()
}
