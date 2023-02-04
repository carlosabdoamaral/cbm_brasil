package occurrence

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/rabbit"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/responses"
	"github.com/gin-gonic/gin"
)

func HandleUpdateStatusRequest(ctx *gin.Context, publishToQueueWithType string) {
	var (
		body []byte
		err  error
	)

	body, err = ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	jsonModel := responses.UpdateOccurrenceStatus{}
	err = json.Unmarshal(body, &jsonModel)
	if err != nil {
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	if jsonModel.IdFirefighter == 0 || jsonModel.IdOccurrence == 0 {
		ctx.IndentedJSON(http.StatusConflict, "eighter occurrence or firefighter id and must be greater than 0")
		return
	}

	err = rabbit.PublishMessage(body, publishToQueueWithType)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, "accept request successfully sent to queue")
}
