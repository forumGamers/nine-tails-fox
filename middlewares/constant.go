package middlewares

import (
	"github.com/forumGamers/nine-tails-fox/web"
	"github.com/gin-gonic/gin"
)

type Middlewares interface {
	Authentication(c *gin.Context)
}

type MiddlewaresImpl struct {
	web.ResponseWriter
}

func NewMiddlewares(w web.ResponseWriter) Middlewares {
	return &MiddlewaresImpl{w}
}
