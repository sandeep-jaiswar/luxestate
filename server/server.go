package server

import (
	"luxestate/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // User routes
    r.GET("/users", controllers.GetUsers)
    r.GET("/users/:id", controllers.GetUserByID)
    r.POST("/users", controllers.CreateUser)
    r.POST("/login", controllers.Login)

    return r
}
