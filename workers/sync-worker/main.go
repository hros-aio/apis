package main

import (
	"github.com/hros-aio/apis/workers/sync-worker/app"
	"github.com/tinh-tinh/tinhtinh/microservices/kafka"
)

func main() {
	server := kafka.New(app.NewModule, kafka.Options{
		GroupID: "hros_sync_worker_group",
		Options: kafka.Config{
			Brokers: []string{"localhost:9092"},
		},
	})

	server.Listen()
}
