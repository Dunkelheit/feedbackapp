package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"

	"github.com/Dunkelheit/feedbackapp/model"
	"gopkg.in/gin-gonic/gin.v1"
)

func ping(c *gin.Context) {
	card := model.Card{ID: 1, Title: "Hello"}
	c.JSON(http.StatusOK, card)
}

func main() {

	// initialize the DbMap
	dbmap := initDb()
	defer dbmap.Db.Close()

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

func initDb() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("postgres", "user=arturo.martinez dbname=feedbackapp sslmode=disable")
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(model.Card{}, "cards").SetKeys(true, "ID")
	dbmap.AddTableWithName(model.User{}, "users").SetKeys(true, "ID")
	dbmap.AddTableWithName(model.Review{}, "reviews").SetKeys(true, "ID")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
