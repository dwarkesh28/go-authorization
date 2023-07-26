package main

import (
	"go-jwt/controllers"
	"go-jwt/initializers"
	"go-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Sugnup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	// r.PUT("/posts/:id", controllers.PostUpdate)
	// r.GET("/posts", controllers.PostsIndex)
	// r.GET("/posts/:id", controllers.PostById)
	// r.DELETE("/posts/:id", controllers.PostDelete)

	// fmt.Println("time", time.Now().String())

	r.Run() // listen and serve on 0.0.0.0:8080
}
