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

	jsonModel := responses.SignInByCPF{}
	err = json.Unmarshal(body, &jsonModel)
	if err != nil {
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	protoMessage := responses.NewSignInByCPFModelFromJSONToProto(&jsonModel)
	res, err := common.AuthServiceClient.SignInByCPF(ctx.Request.Context(), protoMessage)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if res.GetEmail() == "" {
		ctx.IndentedJSON(http.StatusOK, "account not found with this credentials")
		return
	}

	ctx.IndentedJSON(http.StatusOK, responses.NewAccountDetailsModelFromProtoToJSON(res))
}

func HandleSignInRequestByEmail(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	jsonModel := responses.SignInByEmail{}
	err = json.Unmarshal(body, &jsonModel)
	if err != nil {
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	protoMessage := responses.NewSignInByEmailModelFromJSONToProto(&jsonModel)
	res, err := common.AuthServiceClient.SignInByEmail(ctx.Request.Context(), protoMessage)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if res.GetEmail() == "" {
		ctx.IndentedJSON(http.StatusOK, "account not found with this credentials")
		return
	}

	ctx.IndentedJSON(http.StatusOK, responses.NewAccountDetailsModelFromProtoToJSON(res))
}
