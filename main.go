package main

import (
  "github.com/gin-gonic/gin"
  "root/handlers"
)

func main() {
  r := gin.Default()
  r.GET("/ping", handlers.Ping)
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}