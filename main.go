package main

import (
	"net/http"
	"github.com/gin-gonic/gin"

	route "backend_community_rest/routes"
	middleware "backend_community_rest/middlewares"
)


var(Application *gin.Engine)


// @title           Join Community
// @version         1.0
// @description     Join Community REST API Service
// @termsOfService  https://belajarqywok.github.io/

// @contact.name    Qywok
// @contact.url     https://belajarqywok.github.io/
// @contact.email   belajarqywok@gmail.com

// @license.name    MIT
// @license.url     https://mit-license.org/

// @host            localhost:3000
// @BasePath        /api
func main() {
	// initialize gin
    Application = gin.New()

	// use middleware
	Application.Use(middleware.CorsMiddleware())

	// Service Routes Grouping
	service_routes := Application.Group("/api")
	route.RoutesGroup(service_routes)

	// Swagger Routes Grouping
	swagger_route := Application.Group("/swagger")
	route.SwaggerGroup(swagger_route)

	// prevent user if access not found route
	Application.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
				"code": "PAGE_NOT_FOUND",
				"message": "Page not found",
		})
	})

	Application.Run("0.0.0.0:3000");
}