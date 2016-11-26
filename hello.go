package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/Dunkelheit/feedbackapp/model"
	"github.com/astaxie/beego/orm"
	"gopkg.in/gin-gonic/gin.v1"
)

func ping(c *gin.Context) {
	card := model.Card{ID: 1, Title: "Hello"}
	c.JSON(http.StatusOK, card)
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

func init() {
	orm.RegisterDriver("mysql", orm.DRPostgres)
	orm.RegisterModel(new(model.Card), new(model.Review), new(model.User))
	orm.RegisterDataBase("default", "postgres", "sslmode=disable dbname=feedbackapp host=localhost user=arturo.martinez")
	err := orm.RunSyncdb("default", true, true)
	if err != nil {
		fmt.Println(err)
	}
}
