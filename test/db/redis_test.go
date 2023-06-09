package db

import (
	"fmt"
	"github.com/daobin/gotools/internal"
	"os"
	"path/filepath"
	"testing"
)

func TestRedis(t *testing.T) {
	t.Run("Redis Option Test >>>", func(t *testing.T) {
		ymlFile, err := filepath.Abs("./redis.yml")
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

		err = internal.DB.Redis.Init(ymlFile)
		if err != nil {
			t.Error(err)
			return
		}

		conn, err := internal.DB.Redis.Get("dev")
		if err != nil {
			t.Error(err)
			return
		}

		err = conn.Set("RT Name", "ABC123", 0).Err()
		fmt.Println("Set RT Name: ABC123")
		if err != nil {
			t.Error(err)
			return
		}

		val := conn.Get("RT Name").Val()
		str := conn.Get("RT Name").String()
		i, err := conn.Get("RT Name").Int()
		fmt.Println("RT Name GET Val: ", val, ", String: ", str, ", Int: ", i)
		if err != nil {
			fmt.Println("RT Name GET Err: ", err.Error())
		}
		conn.Del("RT Name")

		err = conn.Incr("TreeAge").Err()
		if err != nil {
			t.Error(err)
			return
		}
		ta, err := conn.Get("TreeAge").Int()
		fmt.Println("TreeAge: ", ta)
		if err != nil {
			fmt.Println("TreeAge GET Err: ", err.Error())
		}
	})
}
