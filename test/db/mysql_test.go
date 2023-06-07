package db

import (
	"fmt"
	"github.com/daobin/gotools/internal"
	"os"
	"path/filepath"
	"testing"
)

func TestMySql(t *testing.T) {
	type testData struct {
		ZoneId    string `json:"zoneId" gorm:"zone_id"`
		AdapterId string `json:"adapterId" gorm:"adapter_id"`
		ZoneName  string `json:"zoneName" gorm:"zone_name"`
		Status    int    `json:"status" gorm:"status"`
		Provider  string `json:"provider" gorm:"provider"`
		RegionId  string `json:"regionId" gorm:"region_id"`
	}

	t.Run("MySql Option Test >>>", func(t *testing.T) {
		ymlFile, err := filepath.Abs("./mysql.yml")
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

		err = internal.DB.Mysql.Init(ymlFile)
		if err != nil {
			t.Error(err)
			return
		}

		conn, err := internal.DB.Mysql.Get("dev")
		if err != nil {
			t.Error(err)
			return
		}

		dataList := make([]testData, 0)
		err = conn.Table("ecs_zone").Find(&dataList).Error
		if err != nil {
			t.Error(err)
		}

		fmt.Printf("%#v", dataList)
	})
}
