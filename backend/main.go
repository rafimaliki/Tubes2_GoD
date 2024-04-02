package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Cross-Origin Resource Sharing (CORS) middleware
    // Biar bisa ngakses API yang beda port
    r.Use(cors.Default())

    /* ROUTES */

    // Data API
    r.GET("/api/data", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Backend working!",
        })
    })

    // Power of Two API
    r.GET("/api/power-of-two", func(c *gin.Context) {
		// Get the power parameter from the query string
		powerStr := c.Query("power")
		power, err := strconv.Atoi(powerStr) // Convert powerStr to an integer
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid input: power must be a valid integer",
			})
			return
		}

        /* Logic nya golang kaga ngerti ini gmn ngitungnya */
		// Calculate the result (2^power)
		result := 1 << power // Equivalent to: result := int(math.Pow(2, float64(power)))

		// Return the result as JSON
		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	})

    // Run server
    r.Run(":8080")
}
