package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckError(c *gin.Context, err error, mustReturnJson bool) (containErrors bool) {
	if err != nil {
		if mustReturnJson {
			c.IndentedJSON(http.StatusInternalServerError, err.Error())
		} else {
			return true
		}
	}

	return
}
