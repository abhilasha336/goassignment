// package main

// import (
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/abhilasha336/goassignment/cache"
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	cache := cache.NewLRUCache(100) // Initialize the cache with capacity

// 	r := gin.Default()
// 	// Enable CORS
// 	r.Use(func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(http.StatusOK)
// 			return
// 		}
// 		c.Next()
// 	})

// 	//endpoint to get inserted key value in cache
// 	r.GET("/cache/:key", func(c *gin.Context) {
// 		key := c.Param("key")
// 		value, found := cache.Get(key)
// 		if !found {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"value": value})
// 	})

// 	//endpoint to post insert key value in cache
// 	r.POST("/cache/:key", func(c *gin.Context) {
// 		var json struct {
// 			Value    string        `json:"value"`
// 			Duration time.Duration `json:"duration"`
// 		}
// 		if err := c.BindJSON(&json); err != nil {
// 			log.Printf("error in post endpoint cache(bind json) is : %s", err.Error())
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		cache.Set(c.Param("key"), json.Value, json.Duration)
// 		c.JSON(http.StatusCreated, gin.H{"status": "success"})
// 	})

// 	//endpoint to delete insert key value in cache
// 	r.DELETE("/cache/:key", func(c *gin.Context) {
// 		cache.Delete(c.Param("key"))
// 		c.JSON(http.StatusOK, gin.H{"status": "success"})
// 	})

// 	// Specify the port you want to use
// 	port := "8080"

// 	// Start the server and log a message if it fails
// 	err := r.Run(":" + port)
// 	if err != nil {
// 		log.Fatalf("Failed to start server on port %s: %v", port, err)
// 	}
// }

// routes.go
package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/abhilasha336/goassignment/cache" // Update with your actual import path
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

var cacheInstance = cache.NewLRUCache(2) // Initialize the cache with capacity

func main() {
	r := gin.Default()

	// Enable CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	// Endpoint to get inserted key value in cache
	r.GET("/cache/:key", func(c *gin.Context) {
		key := c.Param("key")
		value, found := cacheInstance.Get(key)
		if !found {
			c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"value": value})
	})

	// Endpoint to post insert key value in cache
	r.POST("/cache/:key", func(c *gin.Context) {
		var json struct {
			Value    string `json:"value"`
			Duration string `json:"duration"`
		}
		if err := c.BindJSON(&json); err != nil {
			log.Printf("error in post endpoint cache(bind json) is : %s", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Parse the duration string to int64
		durationSec, err := strconv.ParseInt(json.Duration, 10, 64)
		if err != nil {
			log.Printf("error parsing duration: %s", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid duration format"})
			return
		}

		duration := time.Duration(durationSec) * time.Second // Change from Nanosecond to Second
		cacheInstance.Set(c.Param("key"), json.Value, duration)
		c.JSON(http.StatusCreated, gin.H{"status": "success"})
	})
	// Endpoint to delete insert key value in cache
	r.DELETE("/cache/:key", func(c *gin.Context) {
		cacheInstance.Delete(c.Param("key"))
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	// Socket.io server
	server := socketio.NewServer(nil)

	r.GET("/socket.io/*any", gin.WrapH(server))
	server.OnConnect("/", func(so socketio.Conn) error {
		log.Println("Socket.io client connected")
		return nil
	})

	// Start the socket.io server
	go server.Serve()
	defer server.Close()

	// Specify the port you want to use
	port := "8080"

	// Start the server and log a message if it fails
	err := r.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server on port %s: %v", port, err)
	}
}
