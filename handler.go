package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/naman-dave/technotest/db"
)

const (
	APIInvalidIDError = "Invalid id"
	APICommonError    = "something went wrong!"
)

func HandleUsers(c *gin.Context) {
	users := db.GetUsers()

	c.JSON(http.StatusOK, gin.H{
		"error":   "",
		"message": "fetched users successfully",
		"data":    users,
	})
}

func HandleUserByID(c *gin.Context) {

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   APIInvalidIDError,
			"message": APIInvalidIDError,
			"data":    nil,
		})
	}

	user, err := db.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": APICommonError,
			"data":    nil,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   "",
		"message": "fetched user successfully",
		"data":    user,
	})

}

func HandleAddUser(c *gin.Context) {

	user := db.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "invalid request payload",
			"data":    nil,
		})
	}

	err := db.AddUser(user.ID, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": APICommonError,
			"data":    nil,
		})
	}

	c.JSON(http.StatusAccepted, gin.H{
		"error":   "",
		"message": "added user successfully",
		"data":    user,
	})
}

func HandleEditUsers(c *gin.Context) {

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   APIInvalidIDError,
			"message": APIInvalidIDError,
			"data":    nil,
		})
	}

	user := db.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "invalid request payload",
			"data":    nil,
		})
	}

	if err := db.EditUser(id, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": APICommonError,
			"data":    nil,
		})
	}

	c.JSON(http.StatusAccepted, gin.H{
		"error":   "",
		"message": "edited user successfully",
		"data":    user,
	})

}

func HandleDeleteUsers(c *gin.Context) {

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   APIInvalidIDError,
			"message": APIInvalidIDError,
			"data":    nil,
		})
	}

	if err := db.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": APICommonError,
			"data":    nil,
		})
	}

	c.JSON(http.StatusAccepted, gin.H{
		"error":   "",
		"message": "edited user successfully",
		"data":    nil,
	})
}
