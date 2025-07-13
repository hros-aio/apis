package main

import (
	"github.com/hros-aio/apis/app"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func main() {
	server := core.CreateFactory(app.NewModule)

	server.Listen(3000)
}
