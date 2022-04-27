package httpd

import (
	"github.com/gin-gonic/gin"
	"go.jwt.example/helper"
	"time"
)

func NewUser(c *gin.Context) {
	today := time.Now()
	response := make(map[string]interface{})
	response["date"] = today.Format("2006-02-01 15:04:05")
	helper.Success(c, response)
}

func SaveUser(c *gin.Context) {
	today := time.Now()
	response := make(map[string]interface{})
	response["date"] = today.Format("2006-02-01 15:04:05")
	helper.Success(c, response)
}

func DeleteUser(c *gin.Context) {
	today := time.Now()
	response := make(map[string]interface{})
	response["date"] = today.Format("2006-02-01 15:04:05")
	helper.Success(c, response)
}

func GetUser(c *gin.Context) {
	today := time.Now()
	response := make(map[string]interface{})
	response["date"] = today.Format("2006-02-01 15:04:05")
	helper.Success(c, response)
}
func GetUserDetail(c *gin.Context) {
	today := time.Now()
	response := make(map[string]interface{})
	response["date"] = today.Format("2006-02-01 15:04:05")
	helper.Success(c, response)
}
