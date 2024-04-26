package main

import (
	"log"
	"net/http"
	"time"

	"github.com/abhilasha336/goassignment/cache"
	"github.com/gin-gonic/gin"
)

func main() {
	cache := cache.NewLRUCache(100) // Initialize the cache with capacity

	r := gin.Default()

	//endpoint to get inserted key value in cache
	r.GET("/cache/:key", func(c *gin.Context) {
		key := c.Param("key")
		value, found := cache.Get(key)
		if !found {
			c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"value": value})
	})

	//endpoint to post insert key value in cache
	r.POST("/cache/:key", func(c *gin.Context) {
		var json struct {
			Value    string        `json:"value"`
			Duration time.Duration `json:"duration"`
		}
		if err := c.BindJSON(&json); err != nil {
			log.Printf("error in post endpoint cache(bind json) is : %s", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		cache.Set(c.Param("key"), json.Value, json.Duration)
		c.JSON(http.StatusCreated, gin.H{"status": "success"})
	})

	//endpoint to delete insert key value in cache
	r.DELETE("/cache/:key", func(c *gin.Context) {
		cache.Delete(c.Param("key"))
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	// Specify the port you want to use
	port := "8080"

	// Start the server and log a message if it fails
	err := r.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server on port %s: %v", port, err)
	}
}
