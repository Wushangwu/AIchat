package presentation

import (
	"github.com/gin-gonic/gin"
	"gobaseline/middleware"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	//r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	//r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{

		// 用户登录
		v1.POST("user/generateUser", CreateUser)
	}
	return r
}
