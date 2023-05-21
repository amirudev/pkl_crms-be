package controllers

import (
	"net/http"

	"github.com/amirudev/pkl_crms-be/config"
	"github.com/amirudev/pkl_crms-be/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCustomers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var customers []models.Customer

	// rows, err := db.Table("Customer")

	db.Find(&customers)

	c.JSON(http.StatusOK, gin.H{"data": customers})
}

func GetCustomer(c *gin.Context) {
	var customer models.Customer
	id := c.Param("id")
	db := config.DBConn
	db.First(&customer, id)

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	db := config.DBConn
	c.BindJSON(&customer)
	db.Create(&customer)

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func UpdateCustomer(c *gin.Context) {
	var customer models.Customer
	id := c.Param("id")
	db := config.DBConn
	db.First(&customer, id)
	c.BindJSON(&customer)
	db.Save(&customer)

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func DeleteCustomer(c *gin.Context) {
	var customer models.Customer
	id := c.Param("id")
	db := config.DBConn
	db.Delete(&customer, id)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
