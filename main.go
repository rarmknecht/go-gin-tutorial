// main.go

package main

import (
    "github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

    // Set the router as the default one provided by Gin
    router = gin.Default()

    // Load all templates into memory for performance
    router.LoadHTMLGlob("templates/*")

    // Initialize Routes
    initializeRoutes()

    // Start serving the application
    router.Run()

}
