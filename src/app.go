package src

import (
	"os"
	"strings"
	"unisun/api/unisun-authen-listener/docs"
	"unisun/api/unisun-authen-listener/src/config/environment"
	"unisun/api/unisun-authen-listener/src/constants"
	"unisun/api/unisun-authen-listener/src/controllers"
	"unisun/api/unisun-authen-listener/src/routes"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name    API Support
// @contact.url     http://www.swagger.io/support
// @contact.email   support@swagger.io

// @license.name  MIT License Copyright (c) 2022 Uni-SUN-Team
// @license.url   https://api.unisun.dynu.com/auth/license

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func App() *gin.Engine {
	appEnv := environment.ENV.App
	ginEnv := environment.ENV.Gin
	swagEnv := environment.ENV.Swag
	docs.SwaggerInfo.Title = swagEnv.Title
	docs.SwaggerInfo.Description = swagEnv.Description
	docs.SwaggerInfo.Version = swagEnv.Version
	docs.SwaggerInfo.Host = swagEnv.Host
	docs.SwaggerInfo.BasePath = strings.Join([]string{appEnv.ContextPath, ginEnv.RootPath, ginEnv.Version}, "/")
	docs.SwaggerInfo.Schemes = strings.Split(swagEnv.Schemes, ",")

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	g := r.Group(os.Getenv(constants.CONTEXT_PATH))
	{
		g.GET("/healcheck", controllers.HealthCheckHandler)
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		g.StaticFile("/license", "./LICENSE")
		routes.Auth(g)
	}

	return r
}
