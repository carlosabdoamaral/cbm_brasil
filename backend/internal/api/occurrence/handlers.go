package occurrence

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/rabbit"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/responses"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
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

func HandleFetchOccurrenceByIdRequest(ctx *gin.Context) {
	occurrenceIdParam := ctx.Param("occurrence_id")
	occurrenceId, err := strconv.ParseInt(occurrenceIdParam, 10, 64)
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	protoMessage := &pb.Id{
		Id:        occurrenceId,
		AuthToken: "",
	}

	res, err := common.OccurrenceServiceClient.GetById(ctx.Request.Context(), protoMessage)
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusNoContent, err.Error())
		return
	}

	if res.GetTitle() == "" {
		ctx.IndentedJSON(http.StatusOK, "no occurrence found with this id")
		return
	}

	ctx.IndentedJSON(http.StatusOK, responses.NewOccurrenceDetailsModelFromProtoToJSON(res))
}

func HandleFetchAllOccurrencesRequest(ctx *gin.Context)    {}
func HandleFetchNearbyOccurrencesRequest(ctx *gin.Context) {}
func HandleEditOccurrenceRequest(ctx *gin.Context)         {}
func HandleAcceptOccurrenceRequest(ctx *gin.Context)       {}
func HandleCancelOccurrenceRequest(ctx *gin.Context)       {}
func HandleFinishOccurenceRequest(ctx *gin.Context)        {}
func HandleSoftDeleteOccurrenceById(ctx *gin.Context)      {}
