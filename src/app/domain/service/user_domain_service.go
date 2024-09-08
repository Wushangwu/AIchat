package service

import (
	"aichat/src/app/domain/entity"
	"aichat/src/app/domain/repository"
)

type IUserDomain interface {
	CreateUser(usr entity.User)
}

type IUserServiceImpl struct {
	usrRepos repository.IReposUser
}

func (s *IUserServiceImpl) CreateUser(usr entity.User) {
	s.usrRepos.CreateAccount(usr)
}

func NewIUserServiceImpl(usrRepos repository.IReposUser) IUserDomain {
	return &IUserServiceImpl{
		usrRepos: usrRepos,
	}
}
