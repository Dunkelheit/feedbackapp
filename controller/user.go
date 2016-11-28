package controller

import (
	"net/http"

	"github.com/Dunkelheit/feedbackapp/database"
	"github.com/Dunkelheit/feedbackapp/model"
	"github.com/Dunkelheit/feedbackapp/util"

	"gopkg.in/gin-gonic/gin.v1"
)

// UserByID gets a single user by its identifier
func UserByID(c *gin.Context) {
	var user model.User
	if database.DB.First(&user, util.StringToID(c.Param("userId"))).RecordNotFound() {
		c.JSON(http.StatusNotFound, "User not found")
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
	in := &model.User{}
	err := c.Bind(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user := &model.User{
		Name:     in.Name,
		Email:    in.Email,
		JobTitle: in.JobTitle,
		Avatar:   in.Avatar,
	}
	database.DB.Create(user)

	c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a single user
func DeleteUser(c *gin.Context) {
	var user model.User
	if database.DB.First(&user, util.StringToID(c.Param("userId"))).RecordNotFound() {
		c.JSON(http.StatusNotFound, false)
		return
	}
	database.DB.Delete(&user)
	c.JSON(http.StatusOK, true)
}
