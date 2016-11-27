package controller

import (
	"net/http"

	"github.com/Dunkelheit/feedbackapp/model"
	"github.com/Dunkelheit/feedbackapp/util"
	"github.com/astaxie/beego/orm"
	"gopkg.in/gin-gonic/gin.v1"
)

// UserByID gets a single user by its identifier
func UserByID(c *gin.Context) {
	o := orm.NewOrm()
	o.Using("default")

	user := model.User{ID: util.StringToID(c.Param("userId"))}
	err := o.Read(&user)

	if err == orm.ErrNoRows || err == orm.ErrMissPK {
		c.JSON(http.StatusNotFound, "No user found.")
		return
	}
	c.JSON(http.StatusOK, user)
}
