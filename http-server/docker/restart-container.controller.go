package docker

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type restartContainerController struct {}


func NewRestartContainerController() *restartContainerController {
	return &restartContainerController{}
}

func (t *restartContainerController) Handler(ctx *gin.Context) {
	restartContainerUseCase := NewRestartContainerUseCase()

	id := ctx.Param("id")

	response := restartContainerUseCase.Execute(id)

	ctx.JSON(http.StatusOK, gin.H{"error": response})
}




