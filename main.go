package main

import (
  "github.com/gin-gonic/gin"
  "root/config"
)

func main() {
  r := gin.Default() 
  r = config.Load(r) // Load the routes.
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}