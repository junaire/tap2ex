package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lyleshaw/tap2ex/internal/pkg/service"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("public/*")
	router.GET("/", service.Index)
	router.POST("/submit", service.Submit)
	err := router.Run(":3000")
	if err != nil {
		return
	}
	return
}
