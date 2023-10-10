package docker

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type listContainersInfoController struct {}

func NewListContainersInfoController() *listContainersInfoController {
	return &listContainersInfoController{}
}

func (t *listContainersInfoController) Handler(ctx *gin.Context) {
	listContainersInfoUseCase := NewListContainersInfoUseCase()

	response := listContainersInfoUseCase.Execute()

	ctx.JSON(http.StatusOK, gin.H{"data": response})
}



