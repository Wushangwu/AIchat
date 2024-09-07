package main

import (
	"gobaseline/app/presentation"
	"gobaseline/configs"
	"gorm.io/gorm"
	"os"

	"github.com/gin-gonic/gin"
)

var DB *gorm.DB

func main() {
	// 从配置文件读取配置
	configs.Init()

	// 装载路由
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := presentation.NewRouter()
	r.Run(":3000")
}
