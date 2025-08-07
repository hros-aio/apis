package user

import "github.com/tinh-tinh/sqlorm/v2"

type UserDB struct {
	sqlorm.Model `gorm:"embedded"`
	TenantId     string `gorm:"column:tenant_id;type:varchar(64);index:idx_users_tenant_id" json:"tenantId"`
	Username     string `gorm:"column:username;type:varchar(64);not null;" json:"username"`
	Password     string `gorm:"column:password;not null" json:"password"`
	Email        string `gorm:"column:email; not null" json:"email"`
	IsVerified   bool   `gorm:"column:is_verified;default:false" json:"isVerified"`
	IsBanned     bool   `gorm:"column:is_banned;default:false" json:"isBanned"`
	IsAdmin      bool   `gorm:"column:is_admin;default:false" json:"isAdmin"`
}

func (UserDB) TableName() string {
	return "users"
}
