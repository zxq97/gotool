package config

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/bradfitz/gomemcache/memcache"
	cluster "github.com/bsm/sarama-cluster"
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

	defaultDialTimeout  = 500 * time.Millisecond
	defaultReadTimeout  = 5 * time.Second
	defaultWriteTimeout = 5 * time.Second
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

type Conf struct {
	Mysql   MysqlConf `yaml:"mysql"`
	Tidb    MysqlConf `yaml:"tidb"`
	Mongo   MongoConf `yaml:"mongo"`
	Redis   RedisConf `yaml:"redis"`
	MC      MCConf    `yaml:"mc"`
	Svc     SvcConf   `yaml:"svc"`
	Etcd    EtcdConf  `yaml:"etcd"`
	Kafka   KafkaConf `yaml:"kafka"`
	LogPath LogConf   `yaml:"log_path"`
}

func getIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func LoadYaml(path string) (*Conf, error) {
	conf := &Conf{}
	y, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(y, conf)
	if err != nil {
		return nil, err
	}
	ip := getIP()
	conf.Svc.Addr = ip + conf.Svc.Addr
	conf.Svc.Bind = ip + conf.Svc.Bind
	return conf, err
}

func InitDB(conf *MysqlConf) (sqlbuilder.Database, error) {
	addr := fmt.Sprintf(dbAddr, conf.User, conf.Password, conf.Host, conf.Port, conf.DB)
	dsn, err := mysql.ParseURL(addr)
	if err != nil {
		return nil, err
	}
	return mysql.Open(dsn)
}

func InitMongo(conf *MongoConf) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(conf.Addr)
	switch conf.Type {
	case mongoTypeCluster:
	case mongoTypeReplica:
		opts = opts.SetReadPreference(readpref.SecondaryPreferred())
	default:
	}
	return mongo.Connect(context.TODO(), opts)
}

func InitMC(addr []string) *memcache.Client {
	return memcache.New(addr...)
}

func InitRedis(conf *RedisConf) redis.Cmdable {
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

func InitEtcd(conf *EtcdConf) (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints:   conf.Addr,
		DialTimeout: time.Duration(conf.TTL) * time.Second,
	})
}

func InitKafkaProducer(addr []string) (sarama.SyncProducer, error) {
	kfkConf := sarama.NewConfig()
	kfkConf.Producer.RequiredAcks = sarama.WaitForAll
	kfkConf.Producer.Retry.Max = 3
	kfkConf.Producer.Return.Successes = true
	kfkConf.Net.DialTimeout = defaultDialTimeout
	kfkConf.Net.ReadTimeout = defaultReadTimeout
	kfkConf.Net.WriteTimeout = defaultWriteTimeout
	return sarama.NewSyncProducer(addr, kfkConf)
}

func InitKafkaConsumer(broker, topics []string, group string) (*cluster.Consumer, error) {
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Group.Return.Notifications = true
	return cluster.NewConsumer(broker, group, topics, config)
}

func InitLog(path string) (*log.Logger, error) {
	fp, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	return log.New(fp, "", log.LstdFlags|log.Lshortfile), nil
}
