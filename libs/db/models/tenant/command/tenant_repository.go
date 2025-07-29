package command

import (
	"github.com/tinh-tinh/sqlorm/v2"
	"github.com/tinh-tinh/tinhtinh/microservices"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

func NewRepository(module core.Module) core.Provider {
	repo := sqlorm.InjectRepository[TenantDB](module)
	kafkaClient := microservices.InjectClient(module, microservices.KAFKA)
	logger := logger.InjectLog(module)

	return module.NewProvider(&TenantRepository{
		repo:   repo,
		client: kafkaClient,
		logger: logger,
	})
}

type TenantRepository struct {
	repo   *sqlorm.Repository[TenantDB]
	client microservices.ClientProxy
	logger *logger.Logger
}

func (t *TenantRepository) Create(val any) (*TenantDB, error) {
	data, err := t.repo.Create(val)
	if err != nil {
		return nil, err
	}
	go func() {
		if err := t.client.Publish(data.TableName(), data, microservices.Header{
			"action": "create",
		}); err != nil {
			t.logger.Errorf("Failed to publish create event: %v", err)
		}
	}()
	return data, nil
}

func (t *TenantRepository) Update(id string, val any) (*TenantDB, error) {
	data, err := t.repo.UpdateByID(id, val)
	if err != nil {
		return nil, err
	}
	go func() {
		if err := t.client.Publish(data.TableName(), data, microservices.Header{
			"action": "update",
		}); err != nil {
			t.logger.Errorf("Failed to publish update event: %v", err)
		}
	}()
	return data, nil
}

func (t *TenantRepository) Delete(id string) error {
	data, err := t.repo.FindByID(id)
	if err != nil {
		return err
	}
	err = t.repo.DeleteByID(id)
	if err != nil {
		return err
	}
	go func() {
		if err := t.client.Publish(data.TableName(), data, microservices.Header{
			"action": "delete",
		}); err != nil {
			t.logger.Errorf("Failed to publish delete event: %v", err)
		}
	}()
	return nil
}
