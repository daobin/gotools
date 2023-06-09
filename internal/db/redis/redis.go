package redis

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"strings"
)

type DB struct {
	ymlFile string                   // yml配置文件
	conf    *koanf.Koanf             // 配置处理器
	Pools   map[string]*redis.Client // 连接池
	Tags    []string                 // 多库标识
}

func (d *DB) Init(ymlFile string) error {
	if len(d.Pools) > 0 {
		return nil
	}

	if ymlFile != "" {
		d.ymlFile = ymlFile
	}
	if d.ymlFile == "" {
		return errors.New("Redis Init Failed: YML File cannot e empty")
	}

	d.conf = koanf.New(".")
	err := d.conf.Load(file.Provider(ymlFile), yaml.Parser())
	if err != nil {
		d.conf = nil
		return errors.New("Redis Init Failed: " + err.Error())
	}

	tags := strings.Split(d.conf.String("redis.tags"), ",")
	if len(tags) > 0 {
		d.Pools = make(map[string]*redis.Client, len(tags))

		for _, tag := range tags {
			err = d.build(tag)
			if err != nil {
				return errors.New("Redis Init Failed: " + err.Error())
			}
		}
	}

	return nil
}

func (d *DB) build(tag string) error {
	host := d.conf.String("redis." + tag + ".host")
	if host == "" {
		return errors.New("Redis Build Failed: [" + tag + "] Host Invalid")
	}

	port := d.conf.Int("redis." + tag + ".port")
	if port <= 0 {
		return errors.New("Redis Build Failed: [" + tag + "] Port Invalid")
	}

	password := d.conf.String("redis." + tag + ".password")
	if password == "" {
		return errors.New("Redis Build Failed: [" + tag + "] Password Invalid")
	}

	dbName := d.conf.Int("redis." + tag + ".dbName")
	if dbName < 0 {
		return errors.New("Redis Build Failed: [" + tag + "] DB Name Invalid")
	}

	addr := fmt.Sprintf("%s:%d", host, port)
	options := &redis.Options{
		Addr:     addr,
		Password: password,
		DB:       dbName,
	}

	// 连接池支持
	connMax := d.conf.Int("redis." + tag + ".pool.max")
	if connMax > 1 {
		options.PoolSize = connMax
		//options.MinIdleConns
		//options.IdleTimeout
		//options.DialTimeout
	}

	client := redis.NewClient(options)
	if err := client.Ping().Err(); err != nil {
		return errors.New("Redis Build Failed: " + err.Error())
	}

	d.Pools[tag] = client
	d.Tags = append(d.Tags, tag)

	return nil
}

func (d *DB) check(tag string) error {
	if len(d.Pools) == 0 {
		err := d.Init("")
		if err != nil {
			return errors.New("Redis Check Failed: [" + tag + "] " + err.Error())
		}
	}

	if d.Pools[tag] == nil {
		err := d.build(tag)
		if err != nil {
			return errors.New("Redis Check Failed: [" + tag + "] " + err.Error())
		}
	}

	return nil
}

func (d *DB) Get(tag string) (*redis.Client, error) {
	if tag == "" {
		return nil, errors.New("Redis Get Failed: Tag Invalid")
	}

	err := d.check(tag)
	if err != nil {
		return nil, errors.New("Redis Get Failed: [" + tag + "] " + err.Error())
	}

	return d.Pools[tag], nil
}

func (d *DB) Close() {
	if len(d.Pools) == 0 {
		return
	}

	for tag := range d.Pools {
		d.Pools[tag].Close()
		delete(d.Pools, tag)
	}
}
