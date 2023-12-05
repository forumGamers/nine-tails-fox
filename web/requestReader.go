package web

import "github.com/gin-gonic/gin"

type GetPostParams struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

func NewRequestReader() RequestReader {
	return &RequestReaderImpl{}
}

func (r *RequestReaderImpl) GetParams(c *gin.Context, p any) error {
	return c.ShouldBind(p)
}

func (r *RequestReaderImpl) DefaultPage(q *GetPostParams) {
	if q.Page == 0 {
		q.Page = 1
	}
}

func (r *RequestReaderImpl) DefaultLimit(q *GetPostParams) {
	if q.Limit == 0 {
		q.Limit = 20
	}
}
