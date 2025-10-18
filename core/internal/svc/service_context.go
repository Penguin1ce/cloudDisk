// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"cloudDisk/core/internal/config"
	"cloudDisk/core/internal/middleware"
	"cloudDisk/core/models"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	Rdb    *redis.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.Init(c.MySQL.DataSource),
		Rdb:    models.InitRedis(c),
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
