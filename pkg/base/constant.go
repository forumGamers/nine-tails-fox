package base

import (
	"context"

	cfg "github.com/forumGamers/nine-tails-fox/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type CollectionName string

const (
	Post    CollectionName = "post"
	Like    CollectionName = "like"
	Comment CollectionName = "comment"
	Reply   CollectionName = "replyComment"
	Share   CollectionName = "share"
	Log     CollectionName = "log"
)

type BaseRepoImpl struct {
	DB *mongo.Collection
}

type BaseRepo interface {
	FindByQuery(ctx context.Context, query any) (*mongo.Cursor, error)
	BulkUpdate(ctx context.Context, updateModel []mongo.WriteModel) (*mongo.BulkWriteResult, error)
	Aggregations(ctx context.Context, aggregation any) (*mongo.Cursor, error)
	GetSession() (mongo.Session, error)
	UpdateMany(ctx context.Context, filter any, update any) (*mongo.UpdateResult, error)
	DeleteMany(ctx context.Context, filter any) (*mongo.DeleteResult, error)
}

func GetCollection(name CollectionName) *mongo.Collection {
	return cfg.DB.Collection(string(name))
}
