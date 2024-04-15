package routes

import "github.com/gin-gonic/gin"

// AuthRoutes used for defining the api routes
func AuthRoutes(router *gin.Engine) {
	router.POST("users/signup", controllers.Signup())
}
