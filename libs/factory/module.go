package factory

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/hros-aio/apis/libs/factory/shared"
	redis_store "github.com/redis/go-redis/v9"
	"github.com/tinh-tinh/auth/v2"
	"github.com/tinh-tinh/cacher/storage/redis"
	"github.com/tinh-tinh/cacher/v2"
	"github.com/tinh-tinh/config/v2"
	"github.com/tinh-tinh/fetch/v2"
	"github.com/tinh-tinh/queue/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

func Register() core.Modules {
	return func(module core.Module) core.Module {
		return module.New(core.NewModuleOptions{
			Imports: []core.Modules{
				config.ForRoot[shared.Config]("./config/configuration.yaml", "./config/.env"),
				queue.ForRootFactory(func(ref core.RefProvider) *queue.Options {
					cfg := config.Inject[shared.Config](ref)
					return &queue.Options{
						Connect: &redis_store.Options{
							Addr:     cfg.Redis.Addr,
							Password: cfg.Redis.Pass,
							DB:       cfg.Redis.DB,
						},
						RetryFailures: 5,
						Workers:       10,
					}
				}),
				cacher.RegisterFactory(func(ref core.RefProvider) cacher.Config {
					cfg := config.Inject[shared.Config](ref)

					store := redis.New(redis.Options{
						Connect: &redis_store.Options{
							Addr:     cfg.Redis.Addr,
							Password: cfg.Redis.Pass,
							DB:       cfg.Redis.DB,
						},
					})
					return cacher.Config{
						Store: store,
					}
				}),
				auth.RegisterFactory(func(ref core.RefProvider) auth.JwtOptions {
					cfg := config.Inject[shared.Config](ref)
					return auth.JwtOptions{
						Alg:        jwt.SigningMethodRS256,
						PrivateKey: cfg.AccessTokenPrivateKey,
						PublicKey:  cfg.AccessTokenPublicKey,
					}
				}),
				fetch.Register(&fetch.Config{
					Timeout: 5000,
				}),
				logger.Module(logger.Options{
					Rotate: true,
				}),
			},
		})
	}
}
