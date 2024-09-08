package repository

import (
	"aichat/src/app/domain/entity"
)

type IReposUser interface {
	// GetUser 获取用户
	GetUser(account string) (entity.User, error)
	// CreateAccount 创建账户
	CreateAccount(user entity.User)
}
