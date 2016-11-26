package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Dunkelheit/feedbackapp/model"
	"github.com/astaxie/beego/orm"
	"gopkg.in/gin-gonic/gin.v1"
)

// AllCards retrieves all the available cards
func AllCards(c *gin.Context) {
	var cards []*model.Card
	_, err := o.QueryTable("card").All(&cards)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, cards)
}

// UserById gets a single user by its identifier
func UserById(c *gin.Context) {
	userID, _ := strconv.ParseInt(c.Param("userId"), 10, 64)
	user := model.User{ID: model.ID(userID)}
	err := o.Read(&user)

	if err == orm.ErrNoRows {
		fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	}

	c.JSON(http.StatusOK, user)
}
