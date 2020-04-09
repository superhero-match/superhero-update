/*
  Copyright (C) 2019 - 2020 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package controller

import (
	"fmt"
	"net/http"
	"time"

	ctrl "github.com/superhero-match/superhero-update/cmd/api/model"
	"github.com/superhero-match/superhero-update/internal/producer/model"

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
	err = ctl.Service.Producer.UpdateSuperhero(
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
