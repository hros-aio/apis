package snapshot

import (
	"github.com/tinh-tinh/mongoose/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type SnapshotSchema struct {
	mongoose.BaseSchema
	SessionId string `bson:"sessionId" json:"sessionId"`
	Event     string `bson:"event" json:"event"`
	Completed bool   `bson:"completed" json:"completed"`
	Step      int    `bson:"step" json:"step"`
	Payload   bson.M `bson:"payload" json:"payload"`
}
