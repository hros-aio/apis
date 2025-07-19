package middleware

import (
	"github.com/tinh-tinh/auth/v2"
	"github.com/tinh-tinh/tinhtinh/v2/common/exception"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

const APP_CONTEXT core.CtxKey = "AppContext"

type ContextInfo struct {
	IpAddress string `json:"ipAddress"`
	UserAgent string `json:"userAgent"`
	Referer   string `json:"referer"`
	UserId    string `json:"userId"`
	TenantID  string `json:"tenantId"`
}

func SetContent(ctx core.Ctx) error {
	jwtSvc, ok := ctx.Ref(auth.JWT_TOKEN).(auth.Jwt)
	if !ok {
		return exception.InternalServer("JWT service not found")
	}

	ip := ctx.Headers("X-Forwarded-For")
	if ip == "" {
		ip = ctx.Req().RemoteAddr
	}
	contextInfo := ContextInfo{
		IpAddress: ip,
		UserAgent: ctx.Headers("User-Agent"),
		Referer:   ctx.Req().Referer(),
	}

	// Get tenant id
	authorization := ctx.Headers("Authorization")
	if authorization != "" {
		token := authorization[len("Bearer "):]
		claims, err := jwtSvc.Decode(token)
		if err != nil {
			return exception.Unauthorized("Invalid token")
		}
		if tenantId, ok := claims["tenantId"].(string); ok {
			contextInfo.TenantID = tenantId
		}
		if userId, ok := claims["userId"].(string); ok {
			contextInfo.UserId = userId
		}
	}

	ctx.Set(APP_CONTEXT, contextInfo)
	return ctx.Next()
}
