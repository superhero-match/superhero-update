package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/superhero-update/internal/config"
	"github.com/superhero-update/internal/producer"
)

const (
	timeFormat = "2006-01-02T15:04:05"
)

// Controller holds the controller data.
type Controller struct {
	Producer *producer.Producer
}

// NewController returns new controller.
func NewController(cfg *config.Config) (ctrl *Controller, err error) {
	return &Controller{
		Producer: producer.NewProducer(cfg),
	}, nil
}

// RegisterRoutes registers all the superhero update API routes.
func (ctl *Controller) RegisterRoutes() *gin.Engine {
	router := gin.Default()

	sr := router.Group("/api/v1/superhero_update")

	// sr.Use(c.Authorize)

	sr.POST("/update_superhero", ctl.UpdateSuperhero)

	return router
}
