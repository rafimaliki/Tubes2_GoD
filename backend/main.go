package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Cross-Origin Resource Sharing (CORS) middleware
    // Biar bisa ngakses API yang beda port
    r.Use(cors.Default())

    /* ROUTES */

    // Check API Response
    r.GET("/api/status", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "1",
        })
    })

    r.GET("/api/IDS", func(c *gin.Context) {

        source_wiki := c.Query("source")
        target_wiki := c.Query("target")

        // Panggil fungsi disini

		c.JSON(http.StatusOK, gin.H{
            // "source" : source_wiki,
            // "target" : target_wiki,
            // Hasil fungsinya taro di sini
			"result" : "None",
		})

        fmt.Println("\033[32mSearch IDS\033[0m")
        fmt.Println("source: ", source_wiki)
        fmt.Println("target: ", target_wiki)
	})

    r.GET("/api/BFS", func(c *gin.Context) {

        source_wiki := c.Query("source")
        target_wiki := c.Query("target")

        // Panggil fungsi disini

		c.JSON(http.StatusOK, gin.H{
            // "source" : source_wiki,
            // "target" : target_wiki,
            // Hasil fungsinya taro di sini 
			"result" : "None",
		})

        fmt.Println("\033[32mSearch BFS\033[0m")
        fmt.Println("source: ", source_wiki)
        fmt.Println("target: ", target_wiki)
	})

    // Run server
    r.Run(":8080")
}
