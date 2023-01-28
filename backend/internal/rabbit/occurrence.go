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

func ConsumeAcceptOccurrenceRequest(ctx context.Context, body []byte) {
	common.LogInfo("ConsumeAcceptOccurrenceRequest")

	jsonModel := responses.UpdateOccurrenceStatus{}
	err := json.Unmarshal(body, &jsonModel)
	if err != nil {
		return
	}

	protoMessage := responses.NewUpdateOccurrenceStatusByIdModelFromJSONToProto(&jsonModel)
	_, err = common.OccurrenceServiceClient.AcceptById(ctx, protoMessage)
	if err != nil {
		return
	}
}

func ConsumeCancelOccurrenceRequest(ctx context.Context, body []byte) {
	common.LogInfo("ConsumeCancelOccurrenceRequest")

	jsonModel := responses.UpdateOccurrenceStatus{}
	err := json.Unmarshal(body, &jsonModel)
	if err != nil {
		return
	}

	protoMessage := responses.NewUpdateOccurrenceStatusByIdModelFromJSONToProto(&jsonModel)
	_, err = common.OccurrenceServiceClient.CancelById(ctx, protoMessage)
	if err != nil {
		return
	}
}

func ConsumeFinishOccurrenceRequest(ctx context.Context, body []byte) {
	common.LogInfo("ConsumeFinishOccurrenceRequest")

	jsonModel := responses.UpdateOccurrenceStatus{}
	err := json.Unmarshal(body, &jsonModel)
	if err != nil {
		return
	}

	protoMessage := responses.NewUpdateOccurrenceStatusByIdModelFromJSONToProto(&jsonModel)
	_, err = common.OccurrenceServiceClient.FinishById(ctx, protoMessage)
	if err != nil {
		return
	}
}
