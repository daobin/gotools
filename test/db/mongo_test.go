package db

import (
	"fmt"
	"github.com/daobin/gotools/db"
	"gopkg.in/mgo.v2/bson"
	"os"
	"path/filepath"
	"testing"
)

func TestMongoDB(t *testing.T) {
	type testData struct {
		Id     bson.ObjectId `json:"id" bson:"_id"`
		DiskId string        `json:"diskId" bson:"diskId"` // 磁盘ID
		Name   string        `json:"name" bson:"name"`     // 磁盘名称
		Type   string        `json:"type" bson:"type"`     // 磁盘类型
		Size   int           `json:"size" bson:"size"`     // 磁盘大小，单位GB
	}

	t.Run("MongoDB Option Test >>>", func(t *testing.T) {
		ymlFile, err := filepath.Abs("./mongo.yml")
		fmt.Println("YML File: ", ymlFile)
		if err != nil {
			t.Error(err)
			return
		}

		_, err = os.Stat(ymlFile)
		if err != nil {
			if os.IsNotExist(err) {
				t.Error(err)
				return
			}
		}

		err = db.Mongo.Init(ymlFile)
		if err != nil {
			t.Error(err)
			return
		}

		conn, err := db.Mongo.GetConn("dev")
		if err != nil {
			t.Error(err)
			return
		}
		defer db.Mongo.CloseCurrentConn(conn)

		dataList := make([]testData, 0)
		err = conn.C("ecs_disk").Find(bson.M{}).All(&dataList)
		if err != nil {
			t.Error(err)
		}

		fmt.Printf("%#v", dataList)
	})
}
