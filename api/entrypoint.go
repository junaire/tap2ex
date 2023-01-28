package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lyleshaw/tap2ex/internal/pkg/service"
	"net/http"
)

// Handler entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	router := gin.Default()
	router.POST("/submit", service.Submit)
	router.ServeHTTP(w, r)
}
