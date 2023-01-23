package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/responses"
	"github.com/gin-gonic/gin"
)

func HandleSignInRequestByCPF(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	jsonModel := responses.SignInByCPFRequest{}
	err = json.Unmarshal(body, &jsonModel)
	if err != nil {
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	protoMessage := responses.NewSignInByCPFRequestFromJSONToProto(&jsonModel)
	res, err := common.AuthServiceClient.SignInByCPF(ctx.Request.Context(), protoMessage)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if res.GetEmail() == "" {
		ctx.IndentedJSON(http.StatusOK, "account not found with this credentials")
		return
	}

	ctx.IndentedJSON(http.StatusOK, responses.NewAccountDetailsFromProtoToJSON(res))
}

func HandleSignInRequestByEmail(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	jsonModel := responses.SignInByEmailRequest{}
	err = json.Unmarshal(body, &jsonModel)
	if err != nil {
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	protoMessage := responses.NewSignInByEmailRequestFromJSONToProto(&jsonModel)
	res, err := common.AuthServiceClient.SignInByEmail(ctx.Request.Context(), protoMessage)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if res.GetEmail() == "" {
		ctx.IndentedJSON(http.StatusOK, "account not found with this credentials")
		return
	}

	ctx.IndentedJSON(http.StatusOK, responses.NewAccountDetailsFromProtoToJSON(res))
}

func HandleSignOutRequest(ctx *gin.Context) {}

func HandleRecoverRequest(ctx *gin.Context) {}
