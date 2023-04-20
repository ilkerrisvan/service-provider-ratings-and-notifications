package api

import (
	"armut-rating-api/internal/util/logger"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

func RateLimiter(loggr logger.ILogger) gin.HandlerFunc {
	rate := ratelimit.NewBucket(time.Second, 10)

	return func(c *gin.Context) {
		// Limit the size of incoming requests to 1 MiB
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 1<<20)
		if rate.TakeAvailable(1) < 1 {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests",
			})
			loggr.Error("Too many requests")
			return
		}
		c.Next()
	}
}
