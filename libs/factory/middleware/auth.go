package middleware

import (
	"github.com/tinh-tinh/auth/v2"
	"github.com/tinh-tinh/tinhtinh/v2/common/exception"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

const IS_PUBLIC = "IsPublicKey"

func IsPublic() *core.Metadata {
	return core.SetMetadata(IS_PUBLIC, true)
}

func AuthN(ctx core.Ctx) error {
	isPublic, ok := ctx.GetMetadata(IS_PUBLIC).(bool)
	if ok && isPublic {
		return ctx.Next()
	}

	jwtSvc, ok := ctx.Ref(auth.JWT_TOKEN).(auth.Jwt)
	if !ok {
		return exception.InternalServer("JWT service not found")
	}

	contextInfo := ctx.Get(APP_CONTEXT).(ContextInfo)
	if !ok {
		return exception.InternalServer("Not context")
	}
	if contextInfo.Token == "" {
		return exception.Unauthorized("Empty token")
	}

	_, err := jwtSvc.Verify(contextInfo.Token)
	if err != nil {
		return exception.Unauthorized(err.Error())
	}
	return ctx.Next()
}
