package routes

import (
	"github.com/gin-gonic/gin"

	_ "backend_community_rest/docs"
	swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)


func SwaggerGroup(route *gin.RouterGroup) {

	route.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}