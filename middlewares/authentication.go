package middlewares

import (
	"os"

	"github.com/forumGamers/nine-tails-fox/errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (m *MiddlewaresImpl) Authentication(c *gin.Context) {
	access_token := c.Request.Header.Get("access_token")
	if access_token == "" {
		m.AbortHttp(c, m.New403Error("Forbidden"))
		return
	}

	claim := jwt.MapClaims{}

	if token, err := jwt.ParseWithClaims(access_token, &claim, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	}); err != nil || !token.Valid {
		m.AbortHttp(c, m.New401Error(errors.InvalidToken))
		return
	}

	c.Set("user", claim)
	c.Next()
}
