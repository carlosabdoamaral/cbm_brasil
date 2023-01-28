package occurrence

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/rabbit"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/responses"
	"github.com/gin-gonic/gin"
)

func createOccurrenceObjIsEmpty(model *responses.CreateOccurrence) bool {
	if model.IdFirefighter == 0 ||
		model.IdAuthor == 0 ||
		model.BannerX64 == "" ||
		model.Description == "" ||
		model.Title == "" ||
		model.Location.CEP == "" ||
		model.Location.Country == "" ||
		model.Location.State == "" ||
		model.Location.City == "" ||
		model.Location.Neighborhood == "" ||
		model.Location.Street == "" ||
		model.Location.PlaceNumber == 0 ||
		model.Location.Complement == "" ||
		model.Location.Latitude == 0 ||
		model.Location.Longitude == 0 {
		return true
	}

	return false
}

func HandleNewOccurrenceRequest(ctx *gin.Context) {
	var (
		body []byte
		err  error
	)

	body, err = ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	jsonModel := responses.CreateOccurrence{}
	err = json.Unmarshal(body, &jsonModel)
	if err != nil {
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	if createOccurrenceObjIsEmpty(&jsonModel) {
		ctx.IndentedJSON(http.StatusConflict, "some field is empty")
		return
	}

	err = rabbit.PublishMessage(body, common.CREATE_OCCURRENCE_RABBIT_MSG_TYPE)
	if err != nil {
		ctx.IndentedJSON(http.StatusConflict, err.Error())
	}

	ctx.IndentedJSON(http.StatusOK, "create occurrence was published to queue!")
}
func HandleFetchOccurrenceByIdRequest(ctx *gin.Context)    {}
func HandleFetchAllOccurrencesRequest(ctx *gin.Context)    {}
func HandleFetchNearbyOccurrencesRequest(ctx *gin.Context) {}
func HandleEditOccurrenceRequest(ctx *gin.Context)         {}
func HandleAcceptOccurrenceRequest(ctx *gin.Context)       {}
func HandleCancelOccurrenceRequest(ctx *gin.Context)       {}
func HandleFinishOccurenceRequest(ctx *gin.Context)        {}
func HandleSoftDeleteOccurrenceById(ctx *gin.Context)      {}
