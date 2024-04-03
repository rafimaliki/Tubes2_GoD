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
    r.GET("/api/data", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "1",
        })
    })

    r.GET("/api/search", func(c *gin.Context) {

        source_wiki := c.Query("source")
        target_wiki := c.Query("target")

        // Panggil fungsi disini

		c.JSON(http.StatusOK, gin.H{
            // Hasil fungsinya taro di sini
			"result" : "None",
		})

        fmt.Println("source: ", source_wiki)
        fmt.Println("target: ", target_wiki)
	})

    // Run server
    r.Run(":8080")
}
