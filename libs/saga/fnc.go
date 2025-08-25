package saga

import (
	"reflect"

	"github.com/tinh-tinh/cacher/v2"
	"github.com/tinh-tinh/tinhtinh/microservices"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func SyncFnc(module core.Module, fnc microservices.FactoryFunc) microservices.FactoryFunc {
	eventPublisher := module.Ref(EVENT_PUBLISHER).(*EventPulisher)
	cacheSync := cacher.InjectSchemaByStore[bool](module, cacher.MEMORY)

	return func(ctx microservices.Ctx) error {
		sessionId := ctx.Headers("X-Sync-SessionId")
		session, _ := cacheSync.Get(sessionId)
		if !reflect.ValueOf(session).IsZero() {
			return nil
		}

		err := fnc(ctx)
		if err != nil {
			go eventPublisher.RetrySync(sessionId)
			return err
		}

		return cacheSync.Set(sessionId, true)
	}
}
