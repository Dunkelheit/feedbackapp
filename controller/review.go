package controller

import (
	"net/http"

	"github.com/Dunkelheit/feedbackapp/model"
	"github.com/astaxie/beego/orm"
	"gopkg.in/gin-gonic/gin.v1"
)

// AllReviews retrieves all the available reviews
func AllReviews(c *gin.Context) {
	o := orm.NewOrm()
	o.Using("default")

	var reviews []*model.Review
	_, err := o.QueryTable("review").All(&reviews)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, reviews)
}
