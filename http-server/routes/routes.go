package routes

import (
	docker "api/http-server/docker"

	"github.com/gin-gonic/gin"
)


func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	listContainersInfoController := docker.NewListContainersInfoController()
	restartContainerController := docker.NewRestartContainerController()
	stopContainerController := docker.NewStopContainerController()
	startContainerController := docker.NewStartContainerController()

	v1:= router.Group("/v1")

	v1.Use(AuthorizeMiddleware())
		
	v1.GET("/list-containers-info", listContainersInfoController.Handler);
	v1.POST("/restart-container/:id", restartContainerController.Handler);
	v1.POST("/stop-container/:id", stopContainerController.Handle);
	v1.POST("/start-container/:id", startContainerController.Handle);


	return v1
}