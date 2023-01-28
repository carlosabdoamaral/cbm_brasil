package rabbit

import (
	"context"
	"encoding/json"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/responses"
)

func ConsumeCreateOccurrenceRequest(ctx context.Context, body []byte) {
	common.LogInfo("ConsumeCreateOccurrenceRequest")

	jsonModel := responses.CreateOccurrence{}
	err := json.Unmarshal(body, &jsonModel)
	if err != nil {
		return
	}

	protoMessage := responses.NewCreateOccurrenceModelFromJSONToProto(&jsonModel)
	_, err = common.OccurrenceServiceClient.Create(ctx, protoMessage)
	if err != nil {
		return
	}
}