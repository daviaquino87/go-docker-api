package httpServer

import (
	"api/http-server/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func OpenServerHttp() {
	app := gin.Default()

	if err := godotenv.Load(); err != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000" 
	}

	address := fmt.Sprintf(":%s", port)

	routes.AppRoutes(app)

	app.Run(address)
}

