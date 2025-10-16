package models

import (
	"cloudDisk/core/internal/config"
	"log"

	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"

	_ "github.com/go-sql-driver/mysql"
)

func Init(datasource string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", datasource)
	if err != nil {
		log.Printf("new XORM Engine错误：%v", err)
		return nil
	}
	return engine
}
func InitRedis(c config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.Address,
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
}
