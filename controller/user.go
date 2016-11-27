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

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
	o := orm.NewOrm()
	o.Using("default")

	in := &model.User{}
	err := c.Bind(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request.")
		return
	}
	user := &model.User{
		Name:     in.Name,
		Email:    in.Email,
		JobTitle: in.JobTitle,
		Avatar:   in.Avatar,
	}
	id, err := o.Insert(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	user.ID = model.ID(id)
	c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a single user
func DeleteUser(c *gin.Context) {
	o := orm.NewOrm()
	o.Using("default")

	num, err := o.Delete(&model.User{ID: util.StringToID(c.Param("userId"))})

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else if num == 0 {
		c.JSON(http.StatusInternalServerError, "No user found")
	} else {
		c.JSON(http.StatusOK, num)
	}
}
