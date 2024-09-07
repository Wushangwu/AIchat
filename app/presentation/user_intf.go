package presentation

import (
	"github.com/gin-gonic/gin"
	"gobaseline/app/domain/entity"
	"gobaseline/app/domain/service"
	"gobaseline/app/infrastructure/dbs"
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
