package conf

import (
	"github.com/koding/multiconfig"
	"gitlab.azbit.cn/web/facebook-spider/library/util"
)

type ConfigTOML struct {
	Spider struct {
		CsvPath          string `required:"true" flagUsage:"需要抓取facebook公众号的地址文件"`
		ArticleBaseDir   string `required:"true" flagUsage:"存储抓取文章存储的目录"`
		CrawlMaxSleep    int    `required:"true" default:"60" flagUsage:"需要最多隔多长时间抓取一次"`
		RepeatDays       int    `required:"true" default:"7" flagUsage:"需要抓几天前的"`
		StartHour        int    `default:"6" flagUsage:"每天的几小时开始抓取"`
		StartMin         int    `default:"0" flagUsage:"每天的几分开始抓取"`
		RecusiveMaxCount int    `default:"100" flagUsage:"最多递归多少次"`
	}

	FaceBook struct {
		Account  string `required:"true" flagUsage:"账号"`
		Password string `required:"true" flagUsage:"密码"`
	}

	Redis struct {
		Address  string `required:"true" flagUsage:"服务器地址"`
		Password string `required:"false" flagUsage:"redis的密码"`
		DB       int    `required:"true" flagUsage:"数据库"`
		Prefix   string `required:"true" flagUsage:"前缀"`
	}

	Log struct {
		Type  string `default:"json" flagUsage:"日志格式，json|raw"`
		Level int    `default:"5" flagUsage:"日志级别：0 CRITICAL, 1 ERROR, 2 WARNING, 3 NOTICE, 4 INFO, 5 DEBUG"`
	} `flagUsage:"服务日志配置"`
}

var Config *ConfigTOML

func Init(tomlPath, args string) {
	var loaders = []multiconfig.Loader{
		&multiconfig.TagLoader{},
		&multiconfig.TOMLLoader{Path: tomlPath},
		&multiconfig.EnvironmentLoader{},
	}
	m := multiconfig.DefaultLoader{
		Loader:    multiconfig.MultiLoader(loaders...),
		Validator: multiconfig.MultiValidator(&multiconfig.RequiredValidator{}),
	}
	Config = new(ConfigTOML)
	m.MustLoad(&Config)

	_ = util.PrettyPrint(Config)
}
