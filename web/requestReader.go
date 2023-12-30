package web

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type GetPostParams struct {
	Page    int      `form:"page"`
	Limit   int      `form:"limit"`
	Tags    []string `form:"tags"`
	UserIds []string `form:"userIds"`
}

func NewRequestReader() RequestReader {
	return &RequestReaderImpl{}
}

func (r *RequestReaderImpl) GetParams(c *gin.Context, p any) error {
	return c.ShouldBind(p)
}

func (r *RequestReaderImpl) DefaultPage(q *GetPostParams) *RequestReaderImpl {
	if q.Page == 0 {
		q.Page = 1
	}
	return r
}

func (r *RequestReaderImpl) DefaultLimit(q *GetPostParams) *RequestReaderImpl {
	if q.Limit == 0 {
		q.Limit = 20
	}
	return r
}

func (r *RequestReaderImpl) ParseTags(q *GetPostParams) *RequestReaderImpl {
	if q.Tags != nil && len(q.Tags) > 0 {
		q.Tags = strings.Split(q.Tags[0], ",")
	}
	return r
}

func (r *RequestReaderImpl) ParseUserIds(q *GetPostParams) *RequestReaderImpl {
	if q.UserIds != nil && len(q.Tags) > 0 {
		q.UserIds = strings.Split(q.UserIds[0], ",")
	}
	return r
}
