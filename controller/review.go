package controller

import (
	"net/http"

	"github.com/Dunkelheit/feedbackapp/database"
	"github.com/Dunkelheit/feedbackapp/model"
	"gopkg.in/gin-gonic/gin.v1"
)

// AllReviews retrieves all the available reviews
func AllReviews(c *gin.Context) {
	var reviews []model.Review
	database.DB.Find(&reviews)
	c.JSON(http.StatusOK, reviews)
}
