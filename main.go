package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/rs/cors"

	Config "hellowellness/config"
	Connect "hellowellness/connect"
)

func main() {
	webServerPort, err := Config.WebServerPort()
	if err != nil {
		log.Fatal(err)
		return
	}

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		// v1.GET("/test", Connect.CreateProduct)
		v1.GET("/products", Connect.GetAllProducts)
		v1.POST("/products", Connect.CreateProduct) // New POST route for creating a product.
	}

	c := cors.AllowAll()

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(webServerPort, handler))
}
