package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/generate", Generate)

	router.Run(":8080")
}
