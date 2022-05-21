package mongodb

import (
	"dev-local/lib-provider-go/pkg/v1/common"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/event"
)

const (
	defaultURI               = ""
	defaultHost              = "127.0.0.1:27017"
	defaultParameter         = ""
	defaultUser              = ""
	defaultPassword          = ""
	defaultDatabase          = "test"
	defaultTimeout           = 20
	defaultMaxPoolSize       = 16
	defaultMaxConnIdleTime   = 30
	defaultHeartbeatInterval = 15
	defualtMinPoolSize       = 1
)

type Config struct {
	URI               string
	Database          string
	Timeout           time.Duration
	MaxPoolSize       uint64
	MinPoolSize       uint64
	MaxConnIdleTime   time.Duration
	HeartbeatInterval time.Duration
	PoolMonitor       *event.PoolMonitor
	CommandMonitor    *event.CommandMonitor
}

func NewConfigFromEnv() *Config {
	v := viper.New()
	v.AutomaticEnv()
	common.LoadFromFile(v)

	v.SetDefault("MONGODB_URI", defaultURI)
	uri := v.GetString("MONGODB_URI")

	v.SetDefault("MONGODB_DATABASE", defaultDatabase)
	database := v.GetString("MONGODB_DATABASE")

	if uri == defaultURI {
		v.SetDefault("MONGODB_HOST", defaultHost)
		host := v.GetString("MONGODB_HOST")

		v.SetDefault("MONGODB_PARAMETER", defaultParameter)
		parameter := v.GetString("MONGODB_PARAMETER")

		v.SetDefault("MONGODB_USER", defaultUser)
		user := v.GetString("MONGODB_USER")

		v.SetDefault("MONGODB_PASSWORD", defaultPassword)
		password := v.GetString("MONGODB_PASSWORD")

		if parameter != "" && parameter[0] != '?' {
			parameter = "?" + parameter
		}

		mongoDBLogin := user
		if password != "" && mongoDBLogin != "" {
			mongoDBLogin = mongoDBLogin + ":" + password
		}
		if mongoDBLogin != "" {
			mongoDBLogin = mongoDBLogin + "@"
		}

		uri = "mongodb://" + mongoDBLogin + host + "/" + database + parameter
	}

	v.SetDefault("MONGODB_TIMEOUT", defaultTimeout)
	timeout := v.GetDuration("MONGODB_TIMEOUT") * time.Second

	v.SetDefault("MONGODB_MAX_POOL_SIZE", defaultMaxPoolSize)
	maxPoolSize := v.GetUint64("MONGODB_MAX_POOL_SIZE")

	v.SetDefault("MONGODB_MIN_POOL_SIZE", defualtMinPoolSize)
	minPoolSize := v.GetUint64("MONGODB_MIN_POOL_SIZE")

	v.SetDefault("MONGODB_MAX_CONN_IDLE_TIME", defaultMaxConnIdleTime)
	maxConnIdleTime := v.GetDuration("MONGODB_MAX_CONN_IDLE_TIME") * time.Second

	v.SetDefault("MONGODB_HEARTBEAT_INTERVAL", defaultHeartbeatInterval)
	heartbeatInterval := v.GetDuration("MONGODB_HEARTBEAT_INTERVAL") * time.Second

	return &Config{
		URI:               uri,
		Database:          database,
		Timeout:           timeout,
		MaxPoolSize:       maxPoolSize,
		MinPoolSize:       minPoolSize,
		MaxConnIdleTime:   maxConnIdleTime,
		HeartbeatInterval: heartbeatInterval,
	}
}
