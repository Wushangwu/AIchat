package presentation

import (
	"aichat/app/domain/entity"
	"aichat/app/domain/service"
	"aichat/app/infrastructure/dbs"
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
