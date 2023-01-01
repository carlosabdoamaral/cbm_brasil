package common

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func UnmarshalJson(c *gin.Context, v any) (err error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	CheckError(c, err, true)

	err = json.Unmarshal(body, &v)
	CheckError(c, err, true)

	return
}
