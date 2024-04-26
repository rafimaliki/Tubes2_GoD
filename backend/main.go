package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"backend/func/bfs"
	"backend/func/utils"
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

        fmt.Println("\033[32mSearch IDS\033[0m")
        path := []utils.Wiki{{Title: "Hololive_Produciton",URL: "https://en.wikipedia.org/wiki/Hololive_Production"}, 
                       {Title: "Taiwan", URL: "https://en.wikipedia.org/wiki/Taiwan"},
                       {Title: "SARS", URL: "https://en.wikipedia.org/wiki/SARS"},
                       {Title: "Allergic_bronchopulmonary_aspergillosis",URL: "https://en.wikipedia.org/wiki/Allergic_bronchopulmonary_aspergillosis"},
                       {Title: "Rhizopus_oryzae",URL: "https://en.wikipedia.org/wiki/Rhizopus_oryzae"},
                      }
        duration := utils.Duration{Hours: 0, Minutes: 0, Seconds: 2, Milliseconds: 12}
        checked := 1000

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

