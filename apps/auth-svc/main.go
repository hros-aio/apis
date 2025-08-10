package main

import (
	"github.com/hros-aio/apis/apps/auth-svc/app"
	"github.com/hros-aio/apis/libs/factory"
	"github.com/hros-aio/apis/libs/factory/shared"
	"github.com/tinh-tinh/config/v2"
	"github.com/tinh-tinh/tinhtinh/microservices/kafka"
)

func main() {
	server := factory.Create(app.NewModule, "auth-api")
	cfg := config.Inject[shared.Config](server.Module)

	if cfg.Kafka.Enable {
		server.ConnectMicroservice(kafka.Open(kafka.Options{
			Options: kafka.Config{
				Brokers: []string{"127.0.0.1:9092"},
			},
			GroupID: "admin--app",
		}))
		server.StartAllMicroservices()
	}

	server.Listen(cfg.Port)
}
