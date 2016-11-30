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

	apiRoutes := router.Group("/api")

	apiRoutes.POST("/login", controller.Login)

	userRoutes := apiRoutes.Group("/users")
	{
		userRoutes.GET("/:userId", controller.UserByID)
		userRoutes.PUT("/:userId", ping)
		userRoutes.DELETE("/:userId", controller.DeleteUser)
	}

	cardRoutes := apiRoutes.Group("/cards")
	{
		cardRoutes.GET("", controller.AllCards)
		cardRoutes.POST("", controller.CreateCard)
		cardRoutes.PUT("/:cardId", controller.UpdateCard)
		cardRoutes.DELETE("/:cardId", controller.DeleteCard)
	}

	reviewRoutes := apiRoutes.Group("/reviews")
	{
		reviewRoutes.GET("", controller.AllReviews)
		reviewRoutes.POST("", controller.CreateReview)
		reviewRoutes.GET("/:reviewId", ping)
		reviewRoutes.PUT("/:reviewId", ping)
		reviewRoutes.DELETE("/:reviewId", ping)
	}

	router.StaticFile("/", "./web/dist/index.html")
	router.StaticFile("/admin", "./web/dist/index.html")
	router.Static("/static", "./web/dist/static/")

	router.Run() // listen and server on 0.0.0.0:8080
}

func init() {
	database.OpenDB()
}
