package conf

import (
	"github.com/joho/godotenv"
	"meryl/cache"
	"meryl/model"
	"meryl/util"
	"os"
)

func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	//连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}
