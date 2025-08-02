package main

import (
	"github.com/hros-aio/apis/apps/setting-svc/app"
	"github.com/hros-aio/apis/libs/factory"
)

func main() {
	server := factory.Create(app.NewModule, "setting-api")

	server.Listen(3002)
}
