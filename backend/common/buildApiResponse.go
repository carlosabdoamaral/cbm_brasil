package common

import (
	"net/http"

	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/models"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func BuildAndReturnApiResponse(c *gin.Context, req protoreflect.ProtoMessage) {
	if UseJwtAsApiResponse {
		res, err := ProtoToJwt(req)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(http.StatusOK, &models.JwtToken{
			Jwt: res,
		})
	} else {
		c.IndentedJSON(http.StatusOK, req)
	}

}
