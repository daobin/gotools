package db

import (
	"github.com/daobin/gotools/internal/db/mongo"
	"github.com/daobin/gotools/internal/db/mysql"
	"github.com/daobin/gotools/internal/db/redis"
)

// Tool 数据库工具
type Tool struct {
	Mongo *mongo.DB
	Mysql *mysql.DB
	Redis *redis.DB
}
