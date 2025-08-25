package saga

import (
	"github.com/hros-aio/apis/libs/factory/shared"
	"github.com/tinh-tinh/config/v2"
	"github.com/tinh-tinh/tinhtinh/microservices"
	"github.com/tinh-tinh/tinhtinh/microservices/kafka"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func Register() core.Modules {
	return func(module core.Module) core.Module {
		return module.New(core.NewModuleOptions{
			Imports: []core.Modules{
				microservices.Register(),
				microservices.RegisterClientFactory(
					func(ref core.RefProvider) []microservices.ClientOptions {
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
					},
				),
			},
			Providers: []core.Providers{NewProvider},
			Exports:   []core.Providers{NewProvider},
		})
	}
}
