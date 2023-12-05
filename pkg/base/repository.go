package base

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewBaseRepo(db *mongo.Collection) BaseRepo {
	return &BaseRepoImpl{
		DB: db,
	}
}

func (b *BaseRepoImpl) FindByQuery(ctx context.Context, query any) (*mongo.Cursor, error) {
	return b.DB.Find(ctx, query)
}

func (b *BaseRepoImpl) BulkUpdate(ctx context.Context, updateModel []mongo.WriteModel) (*mongo.BulkWriteResult, error) {
	return b.DB.BulkWrite(ctx, updateModel)
}

func (b *BaseRepoImpl) Aggregations(ctx context.Context, aggregation any) (*mongo.Cursor, error) {
	return b.DB.Aggregate(ctx, aggregation)
}

func (b *BaseRepoImpl) GetSession() (mongo.Session, error) {
	return b.DB.Database().Client().StartSession()
}

func (b *BaseRepoImpl) UpdateMany(ctx context.Context, filter any, update any) (*mongo.UpdateResult, error) {
	return b.DB.UpdateMany(ctx, filter, update)
}

func (b *BaseRepoImpl) DeleteMany(ctx context.Context, filter any) (*mongo.DeleteResult, error) {
	return b.DB.DeleteMany(ctx, filter)
}
