package routes

import (
	"unisun/api/auth-listener/src/controllers"

	"github.com/gin-gonic/gin"
)

func Auth(g *gin.RouterGroup) {
	g.POST("/signin", controllers.Signin)
	g.POST("/register", controllers.Register)
	g.POST("/revoke", controllers.Signout)
	g.POST("/refresh-token", controllers.RefreshToken)
	g.GET("/connect/providers/callback", controllers.CallbackProviderLogin)
	g.POST("/forget-password", controllers.ForgetPassword)
	g.POST("/reset-password", controllers.ResetPassword)
}
