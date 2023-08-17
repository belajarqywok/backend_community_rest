package routes

import (
	
	"github.com/gin-gonic/gin"

	controller "backend_community_rest/controllers"

)


func RoutesGroup(route *gin.RouterGroup) {

	route.POST("/join", controller.AddJoinRequest)

}