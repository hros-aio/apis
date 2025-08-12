package user

import "github.com/hros-aio/apis/libs/psql/common/base"

type UserModel struct {
	base.Model
	TenantId   string `json:"tenantId,omitempty"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	Email      string `json:"email,omitempty"`
	IsVerified bool   `json:"isVerified"`
	IsBanned   bool   `json:"isBanned"`
	IsAdmin    bool   `json:"isAdmin"`
}
