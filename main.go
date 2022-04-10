package main

import (
	"atest/config"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatal("cannot init config ", err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf(`i'm alive and my message is "%s"`, config.EnvConfig.Message)})
	})
	r.GET("/ping2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": `ver 1`})
	})
	// env:
	// - name: MESSAGE
	//   valueFrom:
	// 	configMapKeyRef:
	// 	  name: atest-config
	// 	  key: MESSAGE

	r.NoRoute(func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.Writer.Header().Set("Connection", "keep-alive")
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			c.Writer.Header().Set("Access-Control-Max-Age", "86400")
			c.Writer.WriteHeader(http.StatusNoContent)
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "404 not found"})
	})

	// if err := r.Run("0.0.0.0:80"); err != nil {
	if err := r.Run(fmt.Sprintf(":%d", config.EnvConfig.Port)); err != nil {
		log.Fatal("cannot run server", err)
	}
}
