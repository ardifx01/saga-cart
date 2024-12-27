package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func proxyHandler(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		proxyReq, err := http.NewRequest(c.Request.Method, target+c.Request.URL.Path, c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			return
		}


		proxyReq.Header = c.Request.Header

	
		client := &http.Client{}
		resp, err := client.Do(proxyReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to service"})
			return
		}
		defer resp.Body.Close()

	
		for key, values := range resp.Header {
			for _, value := range values {
				c.Header(key, value)
			}
		}

	
		c.Status(resp.StatusCode)


		io.Copy(c.Writer, resp.Body)
	}
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, 
	}))
	
	router.Any("/api/products/*path", proxyHandler("http://localhost:8081"))
	router.Any("/api/orders/*path", proxyHandler("http://localhost:8082"))

	log.Println("API Gateway running on port 8080...")
	log.Fatal(router.Run(":3333"))
}
