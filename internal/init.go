package internal

import (
	"github.com/daobin/gotools/internal/db"
	"github.com/daobin/gotools/internal/db/mongo"
	"github.com/daobin/gotools/internal/db/mysql"
)

var (
	DB = db.Tool{
		Mongo: new(mongo.DB),
		Mysql: new(mysql.DB),
	}

	Slice = new(sliceTool)
	File  = new(fileTool)
)
