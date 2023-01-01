package handlers

import (
	"encoding/json"
	"io/ioutil"
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
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

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
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	id := &pb.Id{}
	err = json.Unmarshal(body, &id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	grpcRes, err := common.AccountServiceClient.PrivateDetails(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	res, err := common.ProtoToJwt(grpcRes)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	// c.IndentedJSON(http.StatusOK, &models.JwtToken{
	// 	Jwt: res,
	// })
	c.IndentedJSON(http.StatusOK, string(res))
}

func GetAccountPublicDetails(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	id := &pb.Id{}
	err = json.Unmarshal(body, &id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	grpcRes, err := common.AccountServiceClient.PublicDetails(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	res, err := common.ProtoToJwt(grpcRes)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, &models.JwtToken{
		Jwt: res,
	})
}

func CreateAccount(c *gin.Context) {

}
func UpdateAccountById(c *gin.Context) {}

func GenerateAccountToken(c *gin.Context) {}
func RestoreAccountById(c *gin.Context)   {}

func DeleteAccountById(c *gin.Context)     {}
func SoftDeleteAccountById(c *gin.Context) {}
