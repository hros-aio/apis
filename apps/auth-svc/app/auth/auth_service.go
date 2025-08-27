package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/factory/keys"
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
	userCache := cacher.InjectSchemaByStore[middleware.UserContext](module, cacher.REDIS)

	return module.NewProvider(&AuthService{
		userRepo:  userRepo,
		jwtSvc:    jwtSvc,
		config:    config,
		userCache: userCache,
	})
}

func (s *AuthService) Login(ctx *middleware.ContextInfo, input *LoginInput) (*TokenResponse, error) {
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

	return s.generateToken(foundUser, input.RememberMe)
}

func (s *AuthService) GetMe(ctx *middleware.ContextInfo) (*user.UserModel, error) {
	if ctx.User == nil {
		return nil, exception.NotFound("missing context")
	}
	return s.userRepo.FindByID(ctx.User.ID)
}

func (s *AuthService) RefreshAccessToken(ctx *middleware.ContextInfo, input *RefreshAccessTokenInput) (*TokenResponse, error) {
	claims, err := s.jwtSvc.Verify(input.RefreshToken, auth.VerifyOptions{
		PublicKey: s.config.RefreshTokenPublicKey,
	})
	if err != nil {
		return nil, err
	}

	userId, ok := claims["sub"].(string)
	if !ok {
		return nil, exception.BadRequest("Token missing sub field")
	}
	user, err := s.userRepo.FindByID(userId)
	if err != nil {
		return nil, err
	}

	return s.generateToken(user, true)
}

func (s *AuthService) generateToken(foundUser *user.UserModel, isRefreshToken bool) (*TokenResponse, error) {
	tokenRes := &TokenResponse{}
	if isRefreshToken {
		refreshToken, err := s.jwtSvc.Generate(jwt.MapClaims{
			"sub":      foundUser.ID.String(),
			"tenantId": foundUser.TenantId,
		}, auth.GenOptions{
			PrivateKey: s.config.RefreshTokenPrivateKey,
			Exp:        s.config.RefreshTokenExpiresIn,
		})
		if err != nil {
			return nil, err
		}
		tokenRes.RefreshToken = refreshToken
	}

	sessionId := uuid.NewString()
	err := s.userCache.Set(keys.AuthSessionId(sessionId), middleware.UserContext{
		ID:       foundUser.ID.String(),
		Email:    foundUser.Email,
		TenantId: foundUser.TenantId,
	}, cacher.StoreOptions{
		Ttl: s.config.AccessTokenExpiresIn,
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
