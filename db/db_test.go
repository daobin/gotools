package db

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"os"
	"path/filepath"
	"testing"
)

func TestDB(t *testing.T) {
	type testData struct {
		Id           bson.ObjectId `json:"id" bson:"_id"`
		DiskId       string        `json:"diskId" bson:"diskId"`             // 磁盘ID
		Name         string        `json:"name" bson:"name"`                 // 磁盘名称
		Description  string        `json:"description" bson:"description"`   // 磁盘描述
		Type         string        `json:"type" bson:"type"`                 // 磁盘类型
		Size         int           `json:"size" bson:"size"`                 // 磁盘大小，单位GB
		Bandwidth    int           `json:"bandwidth" bson:"bandwidth"`       // 磁盘带宽，单位Mbps
		DeployStatus int           `json:"deployStatus" bson:"deployStatus"` // 部署状态
		RunStatus    string        `json:"runStatus" bson:"runStatus"`       // 运行状态
		VmId         string        `json:"vmId" bson:"vmId"`                 // 实例ID
		ZoneId       string        `json:"zoneId" bson:"zoneId"`             // 区域ID
		TenantId     int           `json:"tenantId" bson:"tenantId"`         // 租户ID
		UserId       int           `json:"userId" bson:"userId"`             // 用户ID
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

		err = Mongo.Init(ymlFile)
		if err != nil {
			t.Error(err)
			return
		}

		conn, err := Mongo.GetConn("dev")
		if err != nil {
			t.Error(err)
			return
		}
		defer Mongo.CloseCurrentConn(conn)

		dataList := make([]testData, 0)
		err = conn.C("ecs_disk").Find(bson.M{}).All(&dataList)
		if err != nil {
			t.Error(err)
		}

		fmt.Println(dataList)
	})
}
