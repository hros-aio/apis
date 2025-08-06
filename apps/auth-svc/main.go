package main

import (
	"github.com/hros-aio/apis/apps/auth-svc/app"
	"github.com/hros-aio/apis/libs/factory"
)

func main() {
	server := factory.Create(app.NewModule, "auth-api")

	server.Listen(3003)
}
