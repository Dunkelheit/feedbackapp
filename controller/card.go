package controller

import (
	"net/http"

	"github.com/Dunkelheit/feedbackapp/model"
	"github.com/astaxie/beego/orm"
	"gopkg.in/gin-gonic/gin.v1"
)

// AllCards retrieves all the available cards
func AllCards(c *gin.Context) {
	o := orm.NewOrm()
	o.Using("default")

	var cards []*model.Card
	_, err := o.QueryTable("card").All(&cards)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, cards)
}

// CreateCard creates a new card
func CreateCard(c *gin.Context) {
	o := orm.NewOrm()
	o.Using("default")

	in := &model.Card{}
	err := c.Bind(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request.")
		return
	}
	card := &model.Card{
		Title:    in.Title,
		Category: model.CardCategory(in.Category),
	}
	id, err := o.Insert(card)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	card.ID = model.ID(id)
	c.JSON(http.StatusOK, card)
}
