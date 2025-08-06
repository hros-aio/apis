package user

import "github.com/hros-aio/apis/libs/psql/common/base"

type UserModel struct {
	base.Model
	TenantId   string `json:"tenantId"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	IsVerified bool   `json:"isVerified"`
	IsBanned   bool   `json:"isBanned"`
}
