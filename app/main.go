package main

import (
  "app/controller"
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  router.GET("/", controller.IndexGet)
  router.Run(":5000")
}