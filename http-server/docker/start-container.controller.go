package docker

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type startContainerController struct {}

func NewStartContainerController () *startContainerController {
	return &startContainerController{}
}

func (t startContainerController) Handle(ctx *gin.Context){
	startContainerUseCase := NewStartContainerUseCase()

	id := ctx.Param("id")

	response := startContainerUseCase.Execute(id)

	ctx.JSON(http.StatusOK, gin.H{"error": response})
}