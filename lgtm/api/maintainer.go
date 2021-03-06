package api

import (
	"github.com/go-gitea/lgtm/cache"
	"github.com/go-gitea/lgtm/model"
	"github.com/go-gitea/lgtm/remote"
	"github.com/go-gitea/lgtm/router/middleware/session"
	"github.com/go-gitea/lgtm/store"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// GetMaintainer gets the MAINTAINER configuration file.
func GetMaintainer(c *gin.Context) {
	var (
		owner = c.Param("owner")
		name  = c.Param("repo")
		user  = session.User(c)
	)
	repo, err := store.GetRepoOwnerName(c, owner, name)
	if err != nil {
		log.Errorf("Error getting repository %s. %s", name, err)
		c.AbortWithStatus(404)
		return
	}
	file, err := remote.GetContents(c, user, repo, "MAINTAINERS")
	if err != nil {
		log.Debugf("no MAINTAINERS file for %s. Checking for team members.", repo.Slug)
		members, merr := cache.GetMembers(c, user, repo.Owner)
		if merr != nil {
			log.Errorf("Error getting repository %s. %s", repo.Slug, err)
			log.Errorf("Error getting org members %s. %s", repo.Owner, merr)
			c.String(404, "MAINTAINERS file not found. %s", err)
			return
		}

		log.Printf("found %v members", len(members))
		for _, member := range members {
			file = append(file, member.Login...)
			file = append(file, '\n')
		}
	}

	maintainer, err := model.ParseMaintainer(file)
	if err != nil {
		log.Errorf("Error parsing MAINTAINERS file for %s. %s", repo.Slug, err)
		c.String(500, "Error parsing MAINTAINERS file. %s.", err)
		return
	}
	c.JSON(200, maintainer)
}

// GetMaintainerOrg gets the MAINTAINER configuration file and returns
// a subset of the file with members belonging to the specified organization.
func GetMaintainerOrg(c *gin.Context) {
	var (
		owner = c.Param("owner")
		name  = c.Param("repo")
		team  = c.Param("org")
		user  = session.User(c)
	)
	repo, err := store.GetRepoOwnerName(c, owner, name)
	if err != nil {
		log.Errorf("Error getting repository %s. %s", name, err)
		c.AbortWithStatus(404)
		return
	}
	file, err := remote.GetContents(c, user, repo, "MAINTAINERS")
	if err != nil {
		log.Errorf("Error getting repository %s. %s", repo.Slug, err)
		c.String(404, "MAINTAINERS file not found. %s", err)
		return
	}
	maintainer, err := model.ParseMaintainer(file)
	if err != nil {
		log.Errorf("Error parsing MAINTAINERS file for %s. %s", repo.Slug, err)
		c.String(500, "Error parsing MAINTAINERS file. %s.", err)
		return
	}
	subset, err := model.FromOrg(maintainer, team)
	if err != nil {
		log.Errorf("Error getting subset of MAINTAINERS file for %s/%s. %s", repo.Slug, team, err)
		c.String(500, "Error getting subset of MAINTAINERS file. %s.", err)
		return
	}
	c.JSON(200, subset)
}
