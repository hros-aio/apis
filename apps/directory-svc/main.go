package main

import (
	"github.com/hros-aio/apis/apps/directory-svc/app"
	"github.com/hros-aio/apis/libs/factory"
)

func main() {
	server := factory.Create(app.NewModule, "directory-api")

	server.Listen(3002)
}
