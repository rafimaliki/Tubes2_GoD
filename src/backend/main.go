package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"backend/func/bfs"
	"backend/func/ids"
	_ "net/http/pprof"
)

func main() {

    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()

    r := gin.Default()

    r.Use(cors.Default())

    /* ROUTES */

    r.GET("/api/status", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "1",
        })
    })

    r.GET("/api/IDS", func(c *gin.Context) {

        // source_wiki := c.Query("source")
        // target_wiki := c.Query("target")

        source_wiki := c.Query("source")
        target_wiki := c.Query("target")

        fmt.Println("\033[32mSearch BFS\033[0m")
		path, duration, checked := ids.EntryPoint(source_wiki, target_wiki)

		c.JSON(http.StatusOK, gin.H{
            "path" : path,
            "duration" : duration,
            "checked" : checked,
		})

	})

    r.GET("/api/BFS", func(c *gin.Context) {

        source_wiki := c.Query("source")
        target_wiki := c.Query("target")

        fmt.Println("\033[32mSearch BFS\033[0m")
		path, duration, checked := bfs.EntryPoint(source_wiki, target_wiki)

		c.JSON(http.StatusOK, gin.H{
            "path" : path,
            "duration" : duration,
            "checked" : checked,
		})
	})

    // Run server
    r.Run(":8080")
}

