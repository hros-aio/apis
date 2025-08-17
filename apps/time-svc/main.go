
package main

import (
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"time-svc/app"
)

func main() {
	server := core.CreateFactory(app.NewModule)

	server.Listen(3000)
}
