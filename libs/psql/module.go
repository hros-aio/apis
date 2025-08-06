package psql

import (
	"fmt"

	"github.com/hros-aio/apis/libs/factory/shared"

	"github.com/tinh-tinh/config/v2"
	"github.com/tinh-tinh/sqlorm/v2"
	"github.com/tinh-tinh/tinhtinh/microservices"
	"github.com/tinh-tinh/tinhtinh/microservices/kafka"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"gorm.io/driver/postgres"
)

func Register(models ...any) core.Modules {
	isSync := len(models) > 0
	return func(module core.Module) core.Module {
		return module.New(core.NewModuleOptions{
			Imports: []core.Modules{
				sqlorm.ForRootFactory(func(ref core.RefProvider) sqlorm.Config {
					cfg := config.Inject[shared.Config](ref)
					dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
						cfg.Postgres.Host,
						cfg.Postgres.Port,
						cfg.Postgres.Username,
						cfg.Postgres.Password,
						cfg.Postgres.Database,
					)

					return sqlorm.Config{
						Dialect: postgres.Open(dns),
						Retry: &sqlorm.RetryOptions{
							MaxRetries: 5,
							Delay:      1000, // milliseconds
						},
						Sync:   isSync,
						Models: models,
					}
				}),
				microservices.RegisterClientFactory(func(ref core.RefProvider) []microservices.ClientOptions {
					cfg := config.Inject[shared.Config](ref)

					kafkaConn := kafka.NewClient(kafka.Options{
						Options: kafka.Config{
							Brokers: cfg.Kafka.Brokers,
							Topics:  cfg.Kafka.Topics,
						},
						GroupID: cfg.Kafka.GroupID,
					})
					return []microservices.ClientOptions{
						{
							Name:      microservices.KAFKA,
							Transport: kafkaConn,
						},
					}
				}),
			},
		})
	}
}
