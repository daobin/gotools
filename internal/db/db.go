package db

import (
	"github.com/daobin/gotools/internal/db/mongo"
	"github.com/daobin/gotools/internal/db/mysql"
)

type Tool struct {
	Mongo *mongo.DB
	Mysql *mysql.DB
}
