package presentation

import (
	"aichat/src/app/domain/entity"
	"aichat/src/app/domain/service"
	"aichat/src/app/infrastructure/dbs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	usrRepo := dbs.NewUserRepoImpl()
	srv := service.NewIUserServiceImpl(usrRepo)
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	srv.CreateUser(user)
}
