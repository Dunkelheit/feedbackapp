package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/Dunkelheit/feedbackapp/controller"
	"github.com/Dunkelheit/feedbackapp/model"
	"github.com/astaxie/beego/orm"
	"gopkg.in/gin-gonic/gin.v1"
)

var o orm.Ormer

func ping(c *gin.Context) {
	card := model.Card{ID: 1, Title: "Hello"}
	c.JSON(http.StatusOK, card)
}

func main() {
	router := gin.Default()

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("", controller.CreateUser)
		userRoutes.GET("/:userId", controller.UserByID)
		userRoutes.PUT("/:userId", ping)
		userRoutes.DELETE("/:userId", ping)
	}

	cardRoutes := router.Group("/cards")
	{
		cardRoutes.GET("", controller.AllCards)
		cardRoutes.POST("", controller.CreateCard)
		cardRoutes.PUT("/:cardId", ping)
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
	orm.RegisterDriver("mysql", orm.DRPostgres)
	orm.RegisterModel(new(model.Card), new(model.Review), new(model.User))
	orm.RegisterDataBase("default", "postgres", "sslmode=disable dbname=feedbackapp host=localhost user=arturo.martinez")
	err := orm.RunSyncdb("default", true, true)
	if err != nil {
		fmt.Println(err)
	}
	o = orm.NewOrm()
	o.Using("default")

	cards := []model.Card{
		{Title: "Great moves", Category: model.CardCategoryPositive},
		{Title: "Computer hacking skills", Category: model.CardCategoryPositive},
		{Title: "Lazy eye", Category: model.CardCategoryNegative},
		{Title: "Stinky feet", Category: model.CardCategoryNegative},
	}
	users := []model.User{
		{Name: "Arturo Martinez", Email: "arturo@icemobile.com", JobTitle: "Node.js Developer", Avatar: "none"},
		{Name: "Rosa van Colmjon", Email: "rosa@icemobile.com", JobTitle: "Scrum Master Extraordinaire", Avatar: "none"},
		{Name: "Caio Borges", Email: "caio@icemobile.com", JobTitle: "Test Engineer", Avatar: "none"},
		{Name: "Marco Silva", Email: "marco@icemobile.com", JobTitle: "Visual Designer", Avatar: "none"},
	}
	fmt.Println(o.InsertMulti(len(cards), cards))
	fmt.Println(o.InsertMulti(len(users), users))
}
