package main

import (
	"log"
	"os"

	"github.com/bootcamp/supermercadito/docs"
	"github.com/bootcamp/supermercadito/internal/middlewares"
	"github.com/bootcamp/supermercadito/internal/producto"
	"github.com/bootcamp/supermercadito/internal/routers"
	"github.com/bootcamp/supermercadito/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Supermercadito.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	//server := gin.Default()
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	server := gin.New()
	server.Use(gin.Recovery())

	server.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Armo los endpoint del server
	storage := store.NewProductStorage()
	repo := producto.NewProductRepository(storage)
	router := routers.NewRouter(server, repo)
	server.Use(middlewares.NewLog())

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.SetProductGroupRoutes(server)

	server.Run(":8080")
}
