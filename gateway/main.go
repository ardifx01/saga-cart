package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func proxyHandler(target string, stripPrefix string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Remove the prefix from the path before forwarding
		path := c.Request.URL.Path
		if stripPrefix != "" && strings.HasPrefix(path, stripPrefix) {
			path = strings.TrimPrefix(path, stripPrefix)
			// Ensure the path starts with a slash
			if !strings.HasPrefix(path, "/") {
				path = "/" + path
			}
		}

		proxyReq, err := http.NewRequest(c.Request.Method, target+path, c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			return
		}

		// Copy headers
		proxyReq.Header = c.Request.Header.Clone()

		// Copy query parameters
		proxyReq.URL.RawQuery = c.Request.URL.RawQuery

		client := &http.Client{}
		resp, err := client.Do(proxyReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to service"})
			return
		}
		defer resp.Body.Close()

		// Copy response headers
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

	router.Any("/api/products/*path", proxyHandler("http://localhost:8081", "/api/products"))
	router.Any("/api/orders/*path", proxyHandler("http://localhost:8082", "/api/orders"))
	router.Any("/api/auth/*path", proxyHandler("http://localhost:8084", "/api/auth"))

	log.Println("API Gateway running on port 3333...")
	log.Fatal(router.Run(":3333"))
}
