package auth

type LoginInput struct {
	Email      string `json:"email" validate:"isEmail,required" example:"paimon@gmail.com"`
	Password   string `json:"password" validate:"isStrongPassword,required" example:"12345678@Tc"`
	RememberMe bool   `json:"rememberMe" example:"false"`
}

type RefreshAccessTokenInput struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type TokenResponse struct {
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
}
