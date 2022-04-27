package httpd

import (
	"github.com/gin-gonic/gin"
	"go.jwt.example/helper"
	"time"
)

func HealthyCheck(c *gin.Context) {
	today := time.Now()
	response := make(map[string]interface{})
	response["date"] = today.Format("2006-02-01 15:04:05")
	helper.Success(c, response)
}

func Index(c *gin.Context) {
	response := make(map[string]interface{})
	response["msg"] = "go jwt example"
	helper.Success(c, response)
}
