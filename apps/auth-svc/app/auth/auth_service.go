package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/factory/shared"
	"github.com/hros-aio/apis/libs/psql/common/user"
	"github.com/tinh-tinh/auth/v2"
	"github.com/tinh-tinh/cacher/v2"
	"github.com/tinh-tinh/config/v2"
	"github.com/tinh-tinh/tinhtinh/v2/common/exception"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

type AuthService struct {
	userRepo  *user.Repository
	jwtSvc    auth.Jwt
	config    *shared.Config
	userCache *cacher.Schema[middleware.UserContext]
}

func NewService(module core.Module) core.Provider {
	userRepo := module.Ref(user.REPOSITORY).(*user.Repository)
	jwtSvc := auth.InjectJwt(module)
	config := config.Inject[shared.Config](module)
	userCache := cacher.Inject[middleware.UserContext](module)

	return module.NewProvider(&AuthService{
		userRepo:  userRepo,
		jwtSvc:    jwtSvc,
		config:    config,
		userCache: userCache,
	})
}

func (s *AuthService) Login(ctx middleware.ContextInfo, input *LoginInput) (*TokenResponse, error) {
	foundUser, err := s.userRepo.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	if foundUser == nil {
		return nil, exception.NotFound("User not found")
	}

	isCompare := auth.VerifyHash(foundUser.Password, input.Password)
	if !isCompare {
		return nil, exception.BadRequest("Password incorrect")
	}

	tokenRes := &TokenResponse{}
	if input.RememberMe {
		refreshToken, err := s.jwtSvc.Generate(jwt.MapClaims{
			"sub":      foundUser.ID.String(),
			"tenantId": foundUser.TenantId,
		}, auth.GenOptions{
			Exp: s.config.RefreshTokenExpiresIn,
		})
		if err != nil {
			return nil, err
		}
		tokenRes.RefreshToken = refreshToken
	}

	sessionId := uuid.NewString()
	err = s.userCache.Set(sessionId, middleware.UserContext{
		ID:       foundUser.ID.String(),
		Email:    foundUser.Email,
		TenantId: foundUser.TenantId,
	})
	if err != nil {
		return nil, err
	}

	accessToken, err := s.jwtSvc.Generate(jwt.MapClaims{
		"sub":      sessionId,
		"tenantId": foundUser.TenantId,
	}, auth.GenOptions{
		Exp: s.config.AccessTokenExpiresIn,
	})
	if err != nil {
		return nil, err
	}
	tokenRes.AccessToken = accessToken

	return tokenRes, nil
}

func (s *AuthService) GetMe(ctx middleware.ContextInfo) {}
