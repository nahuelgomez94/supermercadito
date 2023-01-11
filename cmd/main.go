package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bootcamp/supermercadito/internal/producto"
	"github.com/bootcamp/supermercadito/internal/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	server := gin.Default()
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	fmt.Println(os.Getenv("TOKEN"))

	server.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Armo los endpoint del server
	repo := producto.NewProductRepository()
	router := routers.NewRouter(server, repo)
	router.SetProductGroupRoutes(server)

	server.Run(":8080")
}
