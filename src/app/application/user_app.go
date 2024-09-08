package application

import (
	"aichat/src/app/domain/entity"
	"aichat/src/app/domain/service"
)

type Userapp struct {
	userSvr service.IUserDomain
}

func (u *Userapp) CreateAcount(usr entity.User) {
	u.userSvr.CreateUser(usr)
}
