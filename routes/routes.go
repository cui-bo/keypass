package routes

import (
	"github.com/cui-bo/keypass/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRouter(r *gin.Engine) *gin.Engine {
	grp1 := r.Group("/v1/")
	{
		grp1.GET("users", controllers.GetUsers)
		grp1.POST("user", controllers.CreateUser)
		grp1.GET("user/:id", controllers.GetUserById)
		grp1.PUT("user/:id", controllers.UpdateUser)
		grp1.DELETE("user/:id", controllers.DeleteUser)
	}
	return r
}