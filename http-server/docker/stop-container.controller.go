package docker

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type stopContainerController struct {}

func NewStopContainerController() *stopContainerController {
	return &stopContainerController{}
}

func (t stopContainerController) Handle(ctx *gin.Context){
	stopContainerUseCase := NewStopContainerUseCase()

	id := ctx.Param("id")

	response := stopContainerUseCase.Execute(id)

	ctx.JSON(http.StatusOK, gin.H{"error": response})
}