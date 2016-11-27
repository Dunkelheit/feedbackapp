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
