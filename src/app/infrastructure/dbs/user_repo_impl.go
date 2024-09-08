package dbs

import (
	"aichat/src/app/domain/entity"
	"aichat/src/app/domain/repository"
	"aichat/src/configs"
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
