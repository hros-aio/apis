package main

import (
	"github.com/hros-aio/apis/apps/time-svc/app"
	"github.com/hros-aio/apis/libs/factory"
	"github.com/hros-aio/apis/libs/factory/shared"
	"github.com/tinh-tinh/config/v2"
	"github.com/tinh-tinh/tinhtinh/microservices/kafka"
)

func main() {
	server := factory.Create(app.NewModule, "time-api")
	cfg := config.Inject[shared.Config](server.Module)

	if cfg.Kafka.Enable {
		server.ConnectMicroservice(kafka.Open(kafka.Options{
			Options: kafka.Config{
				Brokers: cfg.Kafka.Brokers,
			},
			GroupID: cfg.Kafka.GroupID,
		}))
		server.StartAllMicroservices()
	}

	server.Listen(cfg.Port)
}
