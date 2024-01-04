package config

import "time"

type Config struct {
	Server   ServerConfig
	Mongo    MongoConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Cors     CorsConfig
	Logger   LoggerConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	InternalPort string
	ExternalPort string
	RunMode      string
	RequestLimit string
}

type LoggerConfig struct {
	FilePath string
	Encoding string
	Level    string
}

type MongoConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DbName          string
	SSLMode         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

type PostgresConfig struct {
	Host             string
	Port             string
	User             string
	Password         string
	DbName           string
	SSLMode          string
	MaxIdleConns     int
	MaxIdleConnsTime int
	MaxOpenConns     int
	ConnMaxLifetime  time.Duration
}

type RedisConfig struct {
	Host               string
	Port               string
	Password           string
	Db                 int
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	IdleCheckFrequency time.Duration
	PoolSize           int
	PoolTimeout        time.Duration
}

type CorsConfig struct {
	AllowOrigins string
}

type JWTConfig struct {
	AccessTokenExpireDuration  time.Duration
	RefreshTokenExpireDuration time.Duration
	Secret                     string
	RefreshSecret              string
}
