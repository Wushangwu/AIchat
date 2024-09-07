package dbs

import (
	"aichat/app/domain/entity"
	"aichat/app/domain/repository"
	"aichat/configs"
)

type UserRepoImpl struct {
}

func NewUserRepoImpl() repository.IReposUser {
	return &UserRepoImpl{}
}

func (u UserRepoImpl) GetUser(account string) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepoImpl) CreateAccount(user entity.User) {
	configs.GetDatabase().Exec("INSERT INTO usr.users(username, user_account) VALUES(?, ?)", user.Name, user.Account)
}
