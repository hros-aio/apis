package shared

import "time"

type Config struct {
	Port                   int            `yaml:"port"`
	Postgres               PostgresConfig `yaml:"postgres"`
	Mongo                  MongoConfig    `yaml:"mongo"`
	Redis                  RedisConfig    `yaml:"redis"`
	Kafka                  KafkaConfig    `yaml:"kafka"`
	Nats                   NatsConfig     `yaml:"nats"`
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

type JwtSecret struct {
	AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRES_IN"`
	RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRES_IN"`
}

type ParamID struct {
	ID string `path:"id" example:"6cdad833-ba6d-49e3-889c-da23b764bb21"`
}

type QueryCompany struct {
	CompanyID string `query:"companyId" example:"6cdad833-ba6d-49e3-889c-da23b764bb21"`
}
