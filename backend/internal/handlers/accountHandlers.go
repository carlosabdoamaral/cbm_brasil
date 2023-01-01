package handlers

import (
	"net/http"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/models"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
	"github.com/gin-gonic/gin"
)

func GetAllAccounts(c *gin.Context) {
	model := &pb.Empty{
		Dummy: "ABCDEF",
	}
	grpcRes, err := common.AccountServiceClient.GetAll(c, model)
	common.CheckError(c, err, true)

	var res []*models.AccountPublicInfos
	for _, account := range grpcRes.GetAccounts() {
		res = append(res, &models.AccountPublicInfos{
			Id:        account.GetId(),
			Name:      account.GetName(),
			Email:     account.GetEmail(),
			CreatedAt: account.GetCreatedAt().AsTime(),
		})
	}

	c.IndentedJSON(http.StatusOK, res)
}

func GetAccountPrivateDetails(c *gin.Context) {
	id := &pb.Id{}
	common.UnmarshalJson(c, &id)

	grpcRes, err := common.AccountServiceClient.PrivateDetails(c, id)
	common.CheckError(c, err, true)

	common.BuildAndReturnApiResponse(c, grpcRes)
}

func GetAccountPublicDetails(c *gin.Context) {
	model := &pb.Id{}
	common.UnmarshalJson(c, &model)

	grpcRes, err := common.AccountServiceClient.PublicDetails(c, model)
	common.CheckError(c, err, true)

	common.BuildAndReturnApiResponse(c, grpcRes)
}

func CreateAccount(c *gin.Context) {
	model := &pb.NewAccountRequest{}
	common.UnmarshalJson(c, &model)

	c.IndentedJSON(http.StatusOK, model)
}
func UpdateAccountById(c *gin.Context) {}

func GenerateAccountToken(c *gin.Context) {}
func RestoreAccountById(c *gin.Context)   {}

func DeleteAccountById(c *gin.Context)     {}
func SoftDeleteAccountById(c *gin.Context) {}
