package controller

import (
	"net/http"

	ctrl "github.com/superheroville-municipality/cmd/api/model"

	"github.com/gin-gonic/gin"
)

// UpdateSuperhero updates Superhero.
func (ctl *Controller) UpdateSuperhero(c *gin.Context) {
	var superhero ctrl.Superhero

	err := c.BindJSON(&superhero)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal server error!",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Superhero updated successfully!",
	})
}
