package sync

import (
	"github.com/tinh-tinh/queue/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule(module core.Module) core.Module {
	return module.New(core.NewModuleOptions{
		Imports:   []core.Modules{queue.Register(QUEUE_NAME)},
		Providers: []core.Providers{NewProcessor, NewHandler},
	})
}
