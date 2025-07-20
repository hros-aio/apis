package tenant

import "github.com/tinh-tinh/mongoose/v2"

type QueryDb struct {
	mongoose.BaseSchema `bson:"inline"`
	SyncId              string `bson:"syncId" json:"syncId"`
	TenantId            string `bson:"tenantId" json:"tenantId"`
	Name                string `bson:"name" json:"name"`
	Description         string `bson:"description" json:"description"`
}
