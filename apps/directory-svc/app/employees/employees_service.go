package employees

import (
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/psql/common/employee"
	"github.com/hros-aio/apis/libs/saga"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

type EmployeeService struct {
	employeeRepo   *employee.Repository
	logger         *logger.Logger
	eventPublisher *saga.EventPulisher
}

func NewService(module core.Module) core.Provider {
	employeeRepo := module.Ref(employee.REPOSITORY).(*employee.Repository)
	logger := logger.InjectLog(module)
	eventPublisher := module.Ref(saga.EVENT_PUBLISHER).(*saga.EventPulisher)

	return module.NewProvider(&EmployeeService{
		employeeRepo:   employeeRepo,
		logger:         logger,
		eventPublisher: eventPublisher,
	})
}

func (s *EmployeeService) Create(ctx middleware.ContextInfo, model *employee.EmployeeModel) (*employee.EmployeeModel, error) {
	createdEmployee, err := s.employeeRepo.Create(model)
	if err != nil {
		s.logger.Error("Failed to create employee", logger.Metadata{
			"error": err.Error(),
			"model": model,
		})
		return nil, err
	}
	return createdEmployee, nil
}
