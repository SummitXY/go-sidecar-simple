package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

const port = ":9999"

func main() {
	fmt.Println("sidecar start...")

	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World, sidecar !")
	})

	api := r.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 转发
	api.GET("/add/:a/:b", func(c *gin.Context) {

		target := "http://192.168.64.2:30001"
		remote, err := url.Parse(target)
		if err != nil {
			c.String(http.StatusOK, "Error is :%s",err.Error())
		} else {
			proxy := httputil.NewSingleHostReverseProxy(remote)
			proxy.ServeHTTP(c.Writer, c.Request)
		}
	})

	r.Run(port)
}
