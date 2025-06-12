package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/add", Add)
	router.GET("/subtract", Subtract)
	router.GET("/multiply", Multiply)
	router.GET("/divide", Divide)

	fmt.Println("[*] Starting server on port 8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run the start the server: %v\n", err)
	}
}
