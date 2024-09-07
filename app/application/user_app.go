package application

import (
	"gobaseline/app/domain/entity"
	"gobaseline/app/domain/service"
)

type Userapp struct {
	userSvr service.IUserDomain
}

func (u *Userapp) CreateAcount(usr entity.User) {
	u.userSvr.CreateUser(usr)
}
