package internal

import (
	"github.com/daobin/gotools/internal/db"
	"github.com/daobin/gotools/internal/db/mongo"
)

var (
	DB    = db.Tool{Mongo: new(mongo.DB)}
	Slice = new(sliceTool)
	File  = new(fileTool)
)
