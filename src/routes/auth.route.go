package routes

import "github.com/gin-gonic/gin"

func Auth(g *gin.RouterGroup) {
	g.POST("/signin")
	g.POST("/register")
	g.POST("/revoke")
	g.POST("/refresh-token")
	g.GET("/connect/providers/callback")
}
