package main

import (
  "app/controller"
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  router.GET("/", controller.IndexGet)
  router.POST("/", controller.UserPost)
  router.Run(":5000")
}