package mongo

import (
	"errors"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"gopkg.in/mgo.v2"
	"strings"
)

type DB struct {
	ymlFile string           // yml配置文件
	conf    *koanf.Koanf     // 配置处理器
	Pools   map[string]*pool // 连接池
	Tags    []string         // 多库标识
}

type pool struct {
	conn   *mgo.Session // 会话连接
	dbName string       // 数据库名
}

func (d *DB) Init(ymlFile string) error {
	if len(d.Pools) > 0 {
		return nil
	}

	if ymlFile != "" {
		d.ymlFile = ymlFile
	}
	if d.ymlFile == "" {
		return errors.New("YML File cannot e empty")
	}

	d.conf = koanf.New(".")
	err := d.conf.Load(file.Provider(ymlFile), yaml.Parser())
	if err != nil {
		d.conf = nil
		return err
	}
	tags := strings.Split(d.conf.String("mongodb.tags"), ",")
	if len(tags) > 0 {
		d.Pools = make(map[string]*pool, len(tags))

		for _, tag := range tags {
			url := d.conf.String("mongodb." + tag + ".url")
			conn, err := mgo.Dial(url)
			if err != nil {
				return errors.New("MongoDB Dial Failed: " + err.Error())
			}
			dbName := d.conf.String("mongodb." + tag + ".dbName")
			if dbName == "" {
				return errors.New("MongoDB DB Name cannot be empty")
			}

			d.Pools[tag] = &pool{conn: conn}
			d.Tags = append(d.Tags, tag)

			// 连接池支持
			pLimit := d.conf.Int("mongodb." + tag + ".pool.limit")
			if pLimit > 1 {
				conn.SetPoolLimit(pLimit)
				conn.SetMode(mgo.Monotonic, true)
			}
		}
	}

	return nil
}

func (d *DB) GetConn(tag string) (*mgo.Database, error) {
	if tag == "" {
		return nil, errors.New("MongoDB Tag cannot be empty")
	}

	err := d.Init("")
	if err != nil {
		return nil, errors.New("MongoDB Init Failed: " + err.Error())
	}

	if d.Pools[tag] == nil {
		return nil, errors.New("MongoDB Conn Get Failed")
	}

	return d.Pools[tag].conn.Copy().DB(d.Pools[tag].dbName), nil
}

func (d *DB) CloseCurrentConn(conn *mgo.Database) {
	conn.Session.Close()
}

func (d *DB) Close() {
	if len(d.Pools) == 0 {
		return
	}

	for tag := range d.Pools {
		d.Pools[tag].conn.Close()
		delete(d.Pools, tag)
	}
}
