package shared

type Config struct {
	Port     int            `yaml:"port"`
	Postgres PostgresConfig `yaml:"postgres"`
	Mongo    MongoConfig    `yaml:"mongo"`
	Redis    RedisConfig    `yaml:"redis"`
	Kafka    KafkaConfig    `yaml:"kafka"`
	Nats     NatsConfig     `yaml:"nats"`
	Jwt      JwtSecret
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
	Brokers []string `yaml:"brokers"`
	Topics  []string `yaml:"topics"`
	GroupID string   `yaml:"groupId"`
}

type NatsConfig struct {
	Addr string `yaml:"addr"`
}

type JwtSecret struct {
	AccessTokenPrivateKey  string `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
}
