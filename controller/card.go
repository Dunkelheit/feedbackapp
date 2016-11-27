package controller

import (
	"net/http"

	"github.com/Dunkelheit/feedbackapp/model"
	"github.com/Dunkelheit/feedbackapp/util"
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
		c.JSON(http.StatusBadRequest, err.Error())
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

// UpdateCard updates a single card
func UpdateCard(c *gin.Context) {
	o := orm.NewOrm()
	o.Using("default")

	in := &model.Card{ID: util.StringToID(c.Param("cardId"))}
	err := c.Bind(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	card := &model.Card{ID: util.StringToID(c.Param("cardId"))}

	if err := o.Read(card); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	card.Title = in.Title
	card.Category = model.CardCategory(in.Category)
	if _, err := o.Update(card); err == nil {
		c.JSON(http.StatusOK, card)
	} else {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}

// DeleteCard deletes a single card
func DeleteCard(c *gin.Context) {
	o := orm.NewOrm()
	o.Using("default")

	num, err := o.Delete(&model.Card{ID: util.StringToID(c.Param("cardId"))})

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else if num == 0 {
		c.JSON(http.StatusInternalServerError, "No card found")
	} else {
		c.JSON(http.StatusOK, num)
	}
}
