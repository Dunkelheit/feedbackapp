package main

import (
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/Dunkelheit/feedbackapp/controller"
	"github.com/Dunkelheit/feedbackapp/database"
	"github.com/Dunkelheit/feedbackapp/model"
	"gopkg.in/gin-gonic/gin.v1"
)

func ping(c *gin.Context) {
	card := model.Card{Title: "Hello"}
	c.JSON(http.StatusOK, card)
}

func main() {
	router := gin.Default()

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("", controller.CreateUser)
		userRoutes.GET("/:userId", controller.UserByID)
		userRoutes.PUT("/:userId", ping)
		userRoutes.DELETE("/:userId", controller.DeleteUser)
	}

	cardRoutes := router.Group("/cards")
	{
		cardRoutes.GET("", controller.AllCards)
		cardRoutes.POST("", controller.CreateCard)
		cardRoutes.PUT("/:cardId", controller.UpdateCard)
		cardRoutes.DELETE("/:cardId", controller.DeleteCard)
	}

	reviewRoutes := router.Group("/reviews")
	{
		reviewRoutes.GET("", controller.AllReviews)
		reviewRoutes.POST("", ping)
		reviewRoutes.GET("/:reviewId", ping)
		reviewRoutes.PUT("/:reviewId", ping)
		reviewRoutes.DELETE("/:reviewId", ping)
	}

	router.Run() // listen and server on 0.0.0.0:8080
}

func init() {
	database.OpenDB()
}
