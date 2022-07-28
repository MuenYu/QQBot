package boot

// 处理配置文件相关逻辑
import (
	"QQBot/app/def"
	"github.com/spf13/viper"
	"log"
)

func Boot() {
	// 读取配置
	loadConfig()
	// 生成全局 OPQ 对象
	loadOPQ()
	// 配置功能转换为任务列表
	loadTasks()
}

// 读取配置文件
func loadConfig() {
	viper.AddConfigPath(def.Path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("ReadInConfig:", err)
	}
	if err := viper.Unmarshal(def.Conf); err != nil {
		log.Fatalln("Config Unmarshal:", err)
	}
}
