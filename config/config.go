package config

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-redis/redis/v8"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/yaml.v2"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

const (
	redisTypeCluster = "cluster"

	mongoTypeCluster = "cluster"
	mongoTypeReplica = "replica"

	dbAddr = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True"
)

type MysqlConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type MongoConf struct {
	DBName string `yaml:"db_name"`
	Addr   string `yaml:"addr"`
	Type   string `yaml:"type"`
}

type RedisConf struct {
	Addr []string `yaml:"addr"`
	DB   int      `yaml:"db"`
	Type string   `yaml:"type"`
}

type MCConf struct {
	Addr []string `yaml:"addr"`
}

type SvcConf struct {
	Bind string `yaml:"bind"`
	Addr string `yaml:"addr"`
	Name string `yaml:"name"`
}

type EtcdConf struct {
	Addr []string `yaml:"addr"`
	TTL  int      `yaml:"ttl"`
}

type KafkaConf struct {
	Addr []string `yaml:"addr"`
}

type LogConf struct {
	Api   string `yaml:"api"`
	Exc   string `yaml:"exc"`
	Debug string `yaml:"debug"`
}

func getIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("not real ip")
}

func LoadYaml(path string, v interface{}) error {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(bs, v)
	//y, err := ioutil.ReadFile(path)
	//if err != nil {
	//	return nil, err
	//}
	//err = yaml.Unmarshal(y, conf)
	//if err != nil {
	//	return nil, err
	//}
	//ip := getIP()
	//conf.Svc.Addr = ip + conf.Svc.Addr
	//conf.Svc.Bind = ip + conf.Svc.Bind
	//return conf, err
}

func (conf *MysqlConf) InitDB() (sqlbuilder.Database, error) {
	addr := fmt.Sprintf(dbAddr, conf.User, conf.Password, conf.Host, conf.Port, conf.DB)
	dsn, err := mysql.ParseURL(addr)
	if err != nil {
		return nil, err
	}
	return mysql.Open(dsn)
}

func (conf *MongoConf) InitMongo() (*mongo.Client, error) {
	opts := options.Client().ApplyURI(conf.Addr)
	switch conf.Type {
	case mongoTypeCluster:
	case mongoTypeReplica:
		opts = opts.SetReadPreference(readpref.SecondaryPreferred())
	default:
	}
	return mongo.Connect(context.TODO(), opts)
}

func (conf *MCConf) InitMC() *memcache.Client {
	return memcache.New(conf.Addr...)
}

func (conf *RedisConf) InitRedis() redis.Cmdable {
	switch conf.Type {
	case redisTypeCluster:
		return redis.NewClusterClient(&redis.ClusterOptions{
			Addrs: conf.Addr,
		})
	default:
		return redis.NewClient(&redis.Options{
			Addr: conf.Addr[0],
			DB:   conf.DB,
		})
	}
}

func (conf *EtcdConf) InitEtcd() (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints:   conf.Addr,
		DialTimeout: time.Duration(conf.TTL) * time.Second,
	})
}

func (conf *SvcConf) InitSvc() error {
	ip, err := getIP()
	if err != nil {
		return err
	}
	conf.Addr = ip + conf.Addr
	conf.Bind = ip + conf.Bind
	return nil
}

func InitLog(path string) (*log.Logger, error) {
	fp, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return log.New(fp, "", log.LstdFlags|log.Lshortfile), nil
}
