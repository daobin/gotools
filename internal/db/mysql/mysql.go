package mysql

import (
	"errors"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

type DB struct {
	ymlFile string              // yml配置文件
	conf    *koanf.Koanf        // 配置处理器
	Pools   map[string]*gorm.DB // 连接池
	Tags    []string            // 多库标识
}

func (d *DB) Init(ymlFile string) error {
	if len(d.Pools) > 0 {
		return nil
	}

	if ymlFile != "" {
		d.ymlFile = ymlFile
	}
	if d.ymlFile == "" {
		return errors.New("MySql Init Failed: YML File cannot e empty")
	}

	d.conf = koanf.New(".")
	err := d.conf.Load(file.Provider(ymlFile), yaml.Parser())
	if err != nil {
		d.conf = nil
		return errors.New("MySql Init Failed: " + err.Error())
	}

	tags := strings.Split(d.conf.String("mysql.tags"), ",")
	if len(tags) > 0 {
		d.Pools = make(map[string]*gorm.DB, len(tags))

		for _, tag := range tags {
			err = d.build(tag)
			if err != nil {
				return errors.New("MySql Init Failed: " + err.Error())
			}
		}
	}

	return nil
}

func (d *DB) build(tag string) error {
	url := d.conf.String("mysql." + tag + ".url")
	if url == "" {
		return errors.New("MySql Build Failed: [" + tag + "] Url Invalid")
	}

	conn, err := gorm.Open(mysql.Open(url))
	if err != nil {
		return errors.New("MySql Build Failed: " + err.Error())
	}

	sqlDB, _ := conn.DB()
	if err = sqlDB.Ping(); err != nil {
		_ = sqlDB.Close()
		return errors.New("MySql Build Failed: " + err.Error())
	}

	// 连接池支持
	connMax := d.conf.Int("mysql." + tag + ".pool.max")
	if connMax > 1 {
		sqlDB, _ := conn.DB()

		lifetimeMinute := d.conf.Int64("mysql." + tag + ".pool.lifetimeMinute")
		if lifetimeMinute <= 0 {
			lifetimeMinute = 60
		}
		sqlDB.SetConnMaxLifetime(time.Duration(lifetimeMinute) * time.Minute)

		timeoutSecond := d.conf.Int64("mysql." + tag + ".pool.timeoutSecond")
		if timeoutSecond <= 0 {
			timeoutSecond = 60
		}
		sqlDB.SetConnMaxIdleTime(time.Duration(timeoutSecond) * time.Second)

		sqlDB.SetMaxIdleConns(connMax * 2)
		sqlDB.SetMaxOpenConns(connMax)
	}

	d.Pools[tag] = conn
	d.Tags = append(d.Tags, tag)

	return nil
}

func (d *DB) check(tag string) error {
	if len(d.Pools) == 0 {
		err := d.Init("")
		if err != nil {
			return errors.New("MySql Check Failed: [" + tag + "] " + err.Error())
		}
	}

	if d.Pools[tag] == nil {
		err := d.build(tag)
		if err != nil {
			return errors.New("MySql Check Failed: [" + tag + "] " + err.Error())
		}
	}

	return nil
}

func (d *DB) Get(tag string) (*gorm.DB, error) {
	if tag == "" {
		return nil, errors.New("MySql Get Failed: Tag Invalid")
	}

	err := d.check(tag)
	if err != nil {
		return nil, errors.New("MySql Get Failed: [" + tag + "] " + err.Error())
	}

	return d.Pools[tag], nil
}

func (d *DB) Close() {
	if len(d.Pools) == 0 {
		return
	}

	for tag := range d.Pools {
		sqlDB, _ := d.Pools[tag].DB()
		_ = sqlDB.Close()
		delete(d.Pools, tag)
	}
}
