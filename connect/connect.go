package connect

import (
	"log"
	"net/http"

	"hellowellness/models"

	"github.com/gin-gonic/gin"

	DB "hellowellness/config"
)

func CreateProduct(c *gin.Context) {
	var product models.Product

	// Bind JSON to the product structure
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Establish a new connection to the database.
	db, errdb := DB.ConnectDB()
	if errdb != nil {
		err_response := "Missing db connection"
		c.JSON(http.StatusInternalServerError, gin.H{"result": err_response})
		log.Println(err_response)
		return
	}
	defer db.Statement.Context.Done()

	// Create a new product
	result := db.Create(&product)
	if result.Error != nil {
		err_response := "Error creating product"
		c.JSON(http.StatusInternalServerError, gin.H{"result": err_response})
		log.Println(err_response, result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Product created successfully", "data": product})
}

func GetAllProducts(c *gin.Context) {
	db, errdb := DB.ConnectDB()
	if errdb != nil {
		err_response := "Missing db connection"
		c.JSON(http.StatusNotFound, gin.H{"result": err_response})
		log.Println(err_response)
		return
	}

	// Obtain all products
	var products []models.Product
	result := db.Find(&products)
	if result.Error != nil {
		err_response := "Error retrieving products"
		c.JSON(http.StatusInternalServerError, gin.H{"result": err_response})
		log.Println(err_response, result.Error)
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		err_response := "Error getting DB connection"
		c.JSON(http.StatusInternalServerError, gin.H{"result": err_response})
		log.Println(err_response, err)
		return
	}

	sqlDB.Close() // Close the connection
	c.JSON(http.StatusOK, gin.H{"Result": products})
}
