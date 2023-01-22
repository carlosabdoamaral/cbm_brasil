package account

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/responses"
	"github.com/gin-gonic/gin"
)

func HandleNewAccountRequest(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
	}

	jsonModel := responses.NewAccountRequestJSON{}
	err = json.Unmarshal(body, &jsonModel)
	if err != nil {
		common.LogError(err.Error())
		return
	}

	protoMessage := responses.NewAccountRequestFromJSONToProto(&jsonModel)
	res, err := common.AccountServiceClient.Create(ctx.Request.Context(), protoMessage)
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
	}

	ctx.IndentedJSON(http.StatusOK, responses.NewAccountDetailsFromProtoToJSON(res))
}

func HandleFetchAccountByIdRequest(ctx *gin.Context)  {}
func HandleEditAccountRequest(ctx *gin.Context)       {}
func HandleSoftDeleteAccountRequest(ctx *gin.Context) {}