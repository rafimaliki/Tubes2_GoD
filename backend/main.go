package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"backend/bfs"
)

func main() {
    r := gin.Default()

   
    r.Use(cors.Default())

    /* ROUTES */

    r.GET("/api/status", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "1",
        })
    })

    r.GET("/api/IDS", func(c *gin.Context) {

        source_wiki := c.Query("source")
        target_wiki := c.Query("target")

        path, duration, err := bfs.Entrypoint(source_wiki, target_wiki)

		c.JSON(http.StatusOK, gin.H{
            "path" : path,
            "duration" : duration,
            "error" : err,
		})

        fmt.Println("\033[32mSearch IDS\033[0m")
        // fmt.Println("source: ", source_wiki)
        // fmt.Println("target: ", target_wiki)
	})

    r.GET("/api/BFS", func(c *gin.Context) {

        source_wiki := c.Query("source")
        target_wiki := c.Query("target")

        fmt.Println("\033[32mSearch BFS\033[0m")
		path, duration, err := bfs.Entrypoint(source_wiki, target_wiki)

		c.JSON(http.StatusOK, gin.H{
            "path" : path,
            "duration" : duration,
            "error" : err,
		})
        // fmt.Println("source: ", source_wiki)
        // fmt.Println("target: ", target_wiki)
	})

    // Run server
    r.Run(":8080")
}
