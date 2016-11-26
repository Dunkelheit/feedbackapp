package main

import (
	"github.com/Dunkelheit/feedbackapp/model"
	"gopkg.in/gin-gonic/gin.v1"
)

func ping(c *gin.Context) {
	card := model.Card{ID: 1, Title: "Hello"}
	c.JSON(200, card)
}

func main() {
	router := gin.Default()

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("", ping)
		userRoutes.GET("/:userId", ping)
		userRoutes.PUT("/:userId", ping)
		userRoutes.DELETE("/:userId", ping)
	}

	cardRoutes := router.Group("/cards")
	{
		cardRoutes.GET("", ping)
		cardRoutes.POST("", ping)
		cardRoutes.PUT("/:cardId", ping)
		cardRoutes.DELETE("/:cardId", ping)
	}

	reviewRoutes := router.Group("/reviews")
	{
		reviewRoutes.GET("", ping)
		reviewRoutes.POST("", ping)
		reviewRoutes.GET("/:reviewId", ping)
		reviewRoutes.PUT("/:reviewId", ping)
		reviewRoutes.DELETE("/:reviewId", ping)
	}

	router.Run() // listen and server on 0.0.0.0:8080
}
