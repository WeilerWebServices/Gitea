package middleware

import (
	"time"

	"github.com/go-gitea/lgtm/cache"

	"github.com/gin-gonic/gin"
	"github.com/ianschenck/envflag"
)

var (
	ttl = envflag.Duration("CACHE_TTL", time.Minute*15, "")
)

// Cache is a simple caching middleware.
func Cache() gin.HandlerFunc {
	cacheInstance := cache.NewTTL(*ttl)
	return func(c *gin.Context) {
		c.Set("cache", cacheInstance)
		c.Next()
	}
}
