package comment

import (
	"context"

	b "github.com/forumGamers/nine-tails-fox/pkg/base"
	"github.com/forumGamers/nine-tails-fox/utils"
	"github.com/forumGamers/nine-tails-fox/web"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentRepo interface {
	FindPostComment(ctx context.Context, postId primitive.ObjectID, query web.GetPostParams) ([]CommentResponse, error)
}

type CommentRepoImpl struct {
	b.BaseRepo
	utils.QueryUtils
}
