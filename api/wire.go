//go:build wireinject
// +build wireinject

package api

import (
	"github.com/google/wire"
	"msa-auth/cache"
	"msa-auth/database"
	"msa-auth/members"
)

func InitializeMemberHandler() members.Controller {
	wire.Build(
		members.NewHandler,
		members.NewService,
		members.NewCommand,
		members.NewQuery,
		database.MysqlConnection,
		database.MysqlNewDSN,
		cache.RedisConnection,
		cache.NewRedisContext,
		members.NewRedis,
	)

	return members.Controller{}
}
