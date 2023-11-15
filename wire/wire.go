//go:build wireinject

package wire

import (
	"go_demo/wire/repository"
	"go_demo/wire/repository/dao"

	"github.com/google/wire"
)

func InitUserRepository() *repository.UserRepository {
	wire.Build(repository.NewUserRepository, dao.NewUserDAO, InitDB)
	return &repository.UserRepository{}
}
