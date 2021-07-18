package rest

import (
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

// RateLimit middleware
func RateLimit(limit float64) gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(limit, nil)
	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			c.Data(httpError.StatusCode, lmt.GetMessageContentType(), []byte(httpError.Message))
			c.Abort()
		} else {
			c.Next()
		}
	}
}

