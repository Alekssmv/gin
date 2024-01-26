package config

import (
	  "github.com/gin-gonic/gin"
	  "root/handlers"
)

// Load loads the middlewares, routes, handlers.
func Load(*gin.Engine) *gin.Engine {
	  r := gin.Default()

	  // Middlewares.
	  r.Use(gin.Logger())
	  r.Use(gin.Recovery())

	//   // 404 Handler.
	//   r.NoRoute(handlers.NotFound)

	  // Ping test.
	  r.GET("/ping", handlers.Ping)

	//   // The user handlers, requiring authentication
	//   authorized := r.Group("/")
	//   //authorized.Use(handlers.AuthMiddleware())
	//   {
	// 	    authorized.GET("users", handlers.Users)
	// 	    authorized.GET("user/:id", handlers.User)
	//   }

	  return r
}