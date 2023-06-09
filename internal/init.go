package internal

import (
	"github.com/daobin/gotools/internal/db"
	"github.com/daobin/gotools/internal/db/mongo"
	"github.com/daobin/gotools/internal/db/mysql"
	"github.com/daobin/gotools/internal/db/redis"
)

var (
	DB = db.Tool{
		Mongo: new(mongo.DB),
		Mysql: new(mysql.DB),
		Redis: new(redis.DB),
	}

	Slice = new(sliceTool)
	File  = new(fileTool)
)
