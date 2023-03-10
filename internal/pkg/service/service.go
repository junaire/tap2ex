package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lyleshaw/tap2ex/internal/pkg/shellConnector"
	"github.com/lyleshaw/tap2ex/pkg/utils/log"
	"strconv"
)

type Request struct {
	IP       string `json:"ip"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Index .
// @router / [GET]
func Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

// Submit .
// @router /submit [POST]
func Submit(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	log.Infof("req: %+v", req)
	port, err := strconv.ParseInt(req.Port, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	command := "curl -s -o /tmp/install.sh https://raw.githubusercontent.com/junaire/tap2ex/main/script/run.sh && bash /tmp/install.sh 2> /tmp/install.log &"
	str, err := shellConnector.ExecuteCommand(req.IP, int(port), req.Username, req.Password, command)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.String(200, str)
}
