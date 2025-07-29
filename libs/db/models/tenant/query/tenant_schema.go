package query

import "github.com/tinh-tinh/mongoose/v2"

type ContactPersonSchema struct {
	ContactName  string `bson:"name" json:"name"`
	ContactEmail string `bson:"email" json:"email"`
	ContactPhone string `bson:"phone" json:"phone"`
}

type TenantSchema struct {
	mongoose.BaseSchema `bson:"inline"`
	SyncId              string              `bson:"syncId" json:"syncId"`
	TenantId            string              `bson:"tenantId" json:"tenantId"`
	Name                string              `bson:"name" json:"name"`
	Description         string              `bson:"description" json:"description"`
	Contact             ContactPersonSchema `bson:"contact" json:"contact"`
}

func (TenantSchema) CollectionName() string {
	return "tenants"
}
