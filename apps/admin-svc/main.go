package main

import (
	"github.com/hros-aio/apis/apps/admin-svc/app"
	"github.com/hros-aio/apis/libs/factory"
)

func main() {
	server := factory.Create(app.NewModule, "admin-api")

	server.Listen(3001)
}
