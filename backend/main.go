package main

import (
	"backend/bfs2"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "net/http/pprof"
)

type Wiki struct {
	Title string
	URL   string
}

type Duration struct {
    Hours       int
    Minutes     int
    Seconds     int
    Milliseconds int
}


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

        fmt.Println("\033[32mSearch IDS\033[0m")
        path := []Wiki{{"Hololive_Produciton","https://en.wikipedia.org/wiki/Hololive_Production"}, 
                       {"Taiwan", "https://en.wikipedia.org/wiki/Taiwan"},
                       {"SARS", "https://en.wikipedia.org/wiki/SARS"},
                       {"Allergic_bronchopulmonary_aspergillosis","https://en.wikipedia.org/wiki/Allergic_bronchopulmonary_aspergillosis"},
                       {"Rhizopus_oryzae","https://en.wikipedia.org/wiki/Rhizopus_oryzae"},
                      }
        duration := Duration{0, 0, 2, 12}
        err := "Error message"

		c.JSON(http.StatusOK, gin.H{
            "path" : path,
            "duration" : duration,
            "error" : err,
		})

	})

    r.GET("/api/BFS", func(c *gin.Context) {

        source_wiki := c.Query("source")
        target_wiki := c.Query("target")

        fmt.Println("\033[32mSearch BFS\033[0m")
		path, duration, err := bfs2.Entrypoint(source_wiki, target_wiki)

		c.JSON(http.StatusOK, gin.H{
            "path" : path,
            "duration" : duration,
            "error" : err,
		})
	})

    // Run server
    r.Run(":8080")
}
