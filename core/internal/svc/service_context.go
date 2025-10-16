// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"cloudDisk/core/internal/config"
	"cloudDisk/core/models"

	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	Rdb    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.Init(c.MySQL.DataSource),
		Rdb:    models.InitRedis(c),
	}
}
