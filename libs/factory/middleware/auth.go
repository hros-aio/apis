package middleware

import (
	"github.com/hros-aio/apis/libs/factory/keys"
	"github.com/hros-aio/apis/libs/factory/shared"
	"github.com/tinh-tinh/auth/v2"
	"github.com/tinh-tinh/cacher/v2"
	"github.com/tinh-tinh/config/v2"
	"github.com/tinh-tinh/tinhtinh/v2/common/exception"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

type UserContext struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	TenantId string `json:"tenantId"`
}

const IS_PUBLIC = "IsPublicKey"

func IsPublic() *core.Metadata {
	return core.SetMetadata(IS_PUBLIC, true)
}

func AuthN(ctx core.Ctx) error {
	// Check metadata
	isPublic, ok := ctx.GetMetadata(IS_PUBLIC).(bool)
	if ok && isPublic {
		return ctx.Next()
	}

	apiKey := ctx.Headers("X-API-Key")
	if apiKey != "" {
		cfg, ok := ctx.Ref(config.ENV).(*shared.Config)
		if !ok {
			return exception.InternalServer("Config not found")
		}
		if cfg.ApiKey != apiKey {
			return exception.Unauthorized("Invalid API key")
		}
		return ctx.Next()
	}

	// Inject providers
	jwtSvc, ok := ctx.Ref(auth.JWT_TOKEN).(auth.Jwt)
	if !ok {
		return exception.InternalServer("JWT service not found")
	}

	// Get context
	contextInfo := core.Execution[ContextInfo](APP_CONTEXT, ctx)
	if contextInfo == nil {
		return exception.InternalServer("Not context")
	}

	if contextInfo.Token == "" {
		return exception.Unauthorized("Empty token")
	}

	// Verify token
	_, err := jwtSvc.Verify(contextInfo.Token)
	if err != nil {
		return exception.Unauthorized(err.Error())
	}

	cacheManager, ok := ctx.Ref(cacher.REDIS).(*cacher.Config)
	if !ok || cacheManager == nil {
		return exception.Unauthorized("session invalid")
	}

	userCacher := cacher.NewSchema[UserContext](*cacheManager)
	user, err := userCacher.Get(keys.AuthSessionId(contextInfo.SessionId))
	if err != nil {
		return exception.Unauthorized(err.Error())
	}

	contextInfo.User = &user
	return ctx.Next()
}
