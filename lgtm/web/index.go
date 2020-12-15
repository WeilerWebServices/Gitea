package web

import (
	"github.com/go-gitea/lgtm/cache"
	"github.com/go-gitea/lgtm/router/middleware/session"
	"github.com/go-gitea/lgtm/shared/token"

	"github.com/gin-gonic/gin"
)

// Index is the handler for index pages.
func Index(c *gin.Context) {
	user := session.User(c)

	switch {
	case user == nil:
		c.HTML(200, "brand.html", gin.H{})
	default:
		teams, _ := cache.GetTeams(c, user)
		csrf, _ := token.New(token.CsrfToken, user.Login).Sign(user.Secret)
		c.HTML(200, "index.html", gin.H{"user": user, "csrf": csrf, "teams": teams})
	}
}
