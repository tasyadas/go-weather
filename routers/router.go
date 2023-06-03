package routers

import (
	"assignment-3/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/record", controllers.RecordStatus)

	return router
}
