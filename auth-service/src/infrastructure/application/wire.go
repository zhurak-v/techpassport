//go:build wireinject
// +build wireinject

package application

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/zhurak-v/techpassport/auth-service/src/infrastructure/database"
	"github.com/zhurak-v/techpassport/auth-service/src/infrastructure/repository"
)

func InitApplication() (*gin.Engine, error) {
	wire.Build(
		database.InitDatabase,
		database.NewMigration,

		repository.NewRoleRepository,

		GinApplication,
	)
	return nil, nil
}
