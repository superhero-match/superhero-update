package controller

import (
	"fmt"
	"net/http"
	"time"

	ctrl "github.com/superhero-update/cmd/api/model"
	"github.com/superhero-update/internal/producer/model"

	"github.com/gin-gonic/gin"
)

// UpdateSuperhero updates Superhero.
func (ctl *Controller) UpdateSuperhero(c *gin.Context) {
	var s ctrl.Superhero

	err := c.BindJSON(&s)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"updated": false,
		})

		return
	}

	t := time.Now().UTC()

	// Publish superhero on Kafka topic to be stored in DB and Elasticsearch.
	err = ctl.Producer.UpdateSuperhero(
		model.Superhero{
			ID:                    s.ID,
			LookingForGender:      s.LookingForGender,
			Age:                   s.Age,
			LookingForAgeMin:      s.LookingForAgeMin,
			LookingForAgeMax:      s.LookingForAgeMax,
			LookingForDistanceMax: s.LookingForDistanceMax,
			DistanceUnit:          s.DistanceUnit,
			Lat:                   s.Lat,
			Lon:                   s.Lon,
			Country:               s.Country,
			City:                  s.City,
			SuperPower:            s.SuperPower,
			AccountType:           s.AccountType,
			UpdatedAt:             t.Format(timeFormat),
		},
	)
	if err != nil {
		fmt.Println("UpdateSuperhero")
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":     http.StatusInternalServerError,
			"updated": false,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"updated": true,
	})
}
