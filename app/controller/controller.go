package controller

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func IndexGet(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World")
}

func UserPost(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.PostForm("message")

	fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
}