package server

import (
	"net/http"

	"github.com/fnproject/fn/api"
	"github.com/fnproject/fn/api/common"
	"github.com/fnproject/fn/api/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) handleAppDelete(c *gin.Context) {
	ctx := c.Request.Context()
	log := common.Logger(ctx)

	appIDorName := c.MustGet(api.App).(string)
	app := &models.App{Name: appIDorName, ID: appIDorName}

	err := s.FireBeforeAppDelete(ctx, app)

	err = s.datastore.RemoveApp(ctx, app)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	err = s.FireAfterAppDelete(ctx, app)
	if err != nil {
		log.WithError(err).Error("error firing after app delete")
		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "App deleted"})
}
