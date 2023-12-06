package like

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Like struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserId    string             `json:"userId" bson:"userId,omitempty"`
	PostId    primitive.ObjectID `json:"postId" bson:"postId,omitempty"`
	CreatedAt time.Time          `json:"CreatedAt" bson:"CreatedAt"`
	UpdatedAt time.Time          `json:"UpdatedAt" bson:"UpdatedAt"`
}

type PostLikes struct {
	Id        primitive.ObjectID `json:"postId" bson:"_id"`
	TotalLike int                `json:"totalLike" bson:"totalLike"`
}
