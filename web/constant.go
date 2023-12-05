package web

import "github.com/gin-gonic/gin"

type ResponseWriter interface {
	WriteResponse(c *gin.Context, response WebResponse)
	AbortHttp(c *gin.Context, err error)
	CustomMsgAbortHttp(c *gin.Context, message string, code int)
	New404Error(msg string) error
	Write200Response(c *gin.Context, msg string, data any)
	New403Error(msg string) error
	New401Error(msg string) error
	Write200ResponseWithMetadata(c *gin.Context, msg string, data any, metadata MetaData)
}

type ResponseWriterImpl struct{}

type RequestReader interface {
	GetParams(c *gin.Context, p any) error
	DefaultPage(q *GetPostParams)
	DefaultLimit(q *GetPostParams)
}

type RequestReaderImpl struct{}
