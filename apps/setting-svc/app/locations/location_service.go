package locations

import (
	"github.com/hros-aio/apis/libs/psql/common/location"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

type LocationService struct {
	locationRepo *location.Repository
	logger       *logger.Logger
}

func NewService(module core.Module) core.Provider {
	locationRepo := module.Ref(location.REPOSITORY).(*location.Repository)
	logger := logger.InjectLog(module)

	return module.NewProvider(&LocationService{
		locationRepo: locationRepo,
		logger:       logger,
	})
}
