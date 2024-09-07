package application

import (
	"aichat/app/domain/entity"
	"aichat/app/domain/service"
)

type Userapp struct {
	userSvr service.IUserDomain
}

func (u *Userapp) CreateAcount(usr entity.User) {
	u.userSvr.CreateUser(usr)
}
