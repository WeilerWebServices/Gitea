package middleware

import (
	"github.com/go-gitea/lgtm/store/datastore"

	"github.com/gin-gonic/gin"
	"github.com/ianschenck/envflag"
)

var (
	driver     = envflag.String("DATABASE_DRIVER", "sqlite3", "")
	datasource = envflag.String("DATABASE_DATASOURCE", "lgtm.sqlite", "")
)

// Store is a middleware to initialize the database.
func Store() gin.HandlerFunc {
	store := datastore.New(*driver, *datasource)
	return func(c *gin.Context) {
		c.Set("store", store)
		c.Next()
	}
}
