package helper

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func Success(c *gin.Context, data interface{}) {
	var response = gin.H{
		"data": data,
	}
	c.JSON(http.StatusOK, response)
}

func SuccessRows(c *gin.Context, data interface{}, total int) {
	var response = gin.H{
		"data":         data,
		"totalRecords": total,
	}
	c.JSON(http.StatusOK, response)
}

func Error(c *gin.Context, error string) {
	var response = gin.H{
		"error": error,
	}
	c.JSON(http.StatusBadRequest, response)
}

func ErrorAuth(c *gin.Context, err error) {
	var response = gin.H{
		"error": err.Error(),
	}
	c.JSON(http.StatusInternalServerError, response)
	c.AbortWithError(500, err)
}

func ApiRequest(c *gin.Context) (map[string]interface{}, error) {
	var data map[string]interface{}
	paramMap := make(map[string]interface{}, 0)
	for k, v := range c.Request.URL.Query() {
		if len(v) == 1 && len(v[0]) != 0 {
			paramMap[k] = v[0]
		}
	}
	if len(paramMap) > 0 {
		return paramMap, nil
	} else {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			return data, err
		}
		json.Unmarshal(body, &data)
		return data, err
	}

}
func ApiRequestArray(c *gin.Context) ([]interface{}, error) {
	var data []interface{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(400, err)
		return data, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		c.AbortWithError(400, err)
		return data, err
	}
	return data, nil
}
