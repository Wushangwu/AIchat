package configs

import (
	"github.com/spf13/viper"
	"gobaseline/configs/initstruct"
)

var Cfg Config

// Config 定义config结构体
type Config struct {
	Database initstruct.DataBaseConfig
}

// Init 初始化配置项
func Init() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := LoggerInit(initstruct.NewConfig())
	if err != nil {
		Logger.Error("logger init error：%d \n" + err.Error())
		return
	}

	if err := viper.ReadInConfig(); err != nil {
		Logger.Error("get properties error：%d \n" + err.Error())
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		Logger.Error("get properties struct error：%d \n" + err.Error())
	}

	dbInit(Cfg.Database)

}
