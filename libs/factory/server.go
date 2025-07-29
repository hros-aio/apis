package factory

import (
	"github.com/tinh-tinh/swagger/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

func Create(module core.ModuleParam, prefix string) *core.App {
	app := core.CreateFactory(module)

	// Apply middleware
	app.SetGlobalPrefix(prefix)
	//app.EnableCors(cors.Options{})
	//	app.Use(compression.Handler())
	app.Use(logger.Handler(logger.MiddlewareOptions{
		Format:             logger.Common,
		SeparateBaseStatus: true,
		Rotate:             true,
		Max:                50,
	}))
	//app.Use(helmet.Handler(helmet.Options{}))

	// swagger
	swagger.SetUp("/docs", app, DefaultSwagger(prefix), swagger.Config{
		PersistAuthorization: true,
	})
	return app
}
