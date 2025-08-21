package middleware

import (
	"github.com/google/uuid"
	"github.com/tinh-tinh/auth/v2"
	"github.com/tinh-tinh/tinhtinh/v2/common/exception"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

const APP_CONTEXT core.CtxKey = "AppContext"

type ContextInfo struct {
	IpAddress string    `json:"ipAddress"`
	UserAgent string    `json:"userAgent"`
	Referer   string    `json:"referer"`
	SessionId string    `json:"sessionId"`
	TenantID  string    `json:"tenantId"`
	CompanyID uuid.UUID `json:"companyId"`
	Token     string    `json:"token"`
	User      *UserContext
}

func SetContext(ctx core.Ctx) error {
	// Inject providers
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

	// Get tenant id from token
	authorization := ctx.Headers("Authorization")
	if authorization != "" {
		token := authorization[len("Bearer "):]
		contextInfo.Token = token
		claims, err := jwtSvc.Decode(token)
		if err != nil {
			return exception.Unauthorized("Invalid token")
		}
		if tenantId, ok := claims["tenantId"].(string); ok {
			contextInfo.TenantID = tenantId
		}
		if companyID, ok := claims["companyId"].(string); ok {
			contextInfo.CompanyID = uuid.MustParse(companyID)
		}
		if sessionId, ok := claims["sub"].(string); ok {
			contextInfo.SessionId = sessionId
		}
	}

	// get from query
	if contextInfo.TenantID == "" && ctx.Query("tenantId") != "" {
		contextInfo.TenantID = ctx.Query("tenantId")
	}

	if contextInfo.CompanyID == uuid.Nil && ctx.Query("companyId") != "" {
		contextInfo.CompanyID = uuid.MustParse(ctx.Query("companyId"))
	}

	ctx.Set(APP_CONTEXT, &contextInfo)
	return ctx.Next()
}
