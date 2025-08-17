package main

import (
	"time-svc/app"

	"github.com/hros-aio/apis/libs/factory"
	"github.com/hros-aio/apis/libs/factory/shared"
	"github.com/tinh-tinh/config/v2"
)

func main() {
	server := factory.Create(app.NewModule, "time-api")
	cfg := config.Inject[shared.Config](server.Module)

	server.Listen(cfg.Port)
}
