/*
  Copyright (C) 2019 - 2022 MWSOFT
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
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	ctrl "github.com/superhero-match/superhero-update/cmd/api/model"
	"github.com/superhero-match/superhero-update/internal/producer/model"
)

// UpdateSuperhero updates Superhero.
func (ctl *Controller) UpdateSuperhero(c *gin.Context) {
	var s ctrl.Superhero

	err := c.BindJSON(&s)
	if checkError(err, c) {
		ctl.Logger.Error(
			"failed to bind request model",
			zap.String("err", err.Error()),
			zap.String("time", time.Now().UTC().Format(ctl.TimeFormat)),
		)

		return
	}

	t := time.Now().UTC()

	// Publish superhero on Kafka topic to be stored in DB and Elasticsearch.
	err = ctl.Service.UpdateSuperhero(
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
			UpdatedAt:             t.Format(ctl.TimeFormat),
		},
	)
	if checkError(err, c) {
		ctl.Logger.Error(
			"failed to update superhero",
			zap.String("err", err.Error()),
			zap.String("time", time.Now().UTC().Format(ctl.TimeFormat)),
		)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"updated": true,
	})
}

func checkError(err error, c *gin.Context) bool {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"updated": false,
		})

		return true
	}

	return false
}
