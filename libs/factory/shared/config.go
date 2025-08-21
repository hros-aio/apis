package shared

import "time"

type Config struct {
	Port                   int            `yaml:"port"`
	Postgres               PostgresConfig `yaml:"postgres"`
	Mongo                  MongoConfig    `yaml:"mongo"`
	Redis                  RedisConfig    `yaml:"redis"`
	Kafka                  KafkaConfig    `yaml:"kafka"`
	Nats                   NatsConfig     `yaml:"nats"`
	Http                   HttpConfig     `yaml:"http"`
	ApiKey                 string         `mapstructure:"API_KEY"`
	AccessTokenPrivateKey  string         `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string         `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration  `mapstructure:"ACCESS_TOKEN_EXPIRES_IN"`
	RefreshTokenPrivateKey string         `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string         `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	RefreshTokenExpiresIn  time.Duration  `mapstructure:"REFRESH_TOKEN_EXPIRES_IN"`
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type MongoConfig struct {
	Url string `yaml:"url"`
}

type RedisConfig struct {
	Addr string `yaml:"addr"`
	Pass string `yaml:"pass"`
	DB   int    `yaml:"db"`
}

type KafkaConfig struct {
	Enable  bool     `yaml:"enable"`
	Brokers []string `yaml:"brokers"`
	Topics  []string `yaml:"topics"`
	GroupID string   `yaml:"group_id"`
}

type NatsConfig struct {
	Addr string `yaml:"addr"`
}

type HttpConfig struct {
	AuthUrl    string `yaml:"auth_url"`
	SettingUrl string `yaml:"setting_url"`
}
