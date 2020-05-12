// main.go

package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

    // Set the router as the default one provided by Gin
    router := gin.Default()

    // Load all templates into memory for performance
    router.LoadHTMLGlob("templates/*")

    // Define route:  /
    // Inline Route Handler Example
    router.GET("/", func(c *gin.Context) {
        
        // Call the HTML method of the Context to render a template
        c.HTML(
            // Set HTTP Status to 200
            http.StatusOK,
            // Use the index.html template
            "index.html",
            // Pass the data that the page uses
            gin.H{
                "title": "Home Page",
            },
        )

    })

    // Start serving the application
    router.Run()

}
