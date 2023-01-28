package occurrence

import (
	"context"
	"fmt"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
)

func CreateOccurrenceHandler(ctx *context.Context, req *pb.CreateOccurrence) error {
	idOccurrence, err := insertOccurrence(ctx, req)
	if err != nil {
		return err
	}

	err = insertOccurrenceLocationById(ctx, idOccurrence, req.Location)
	if err != nil {
		return err
	}

	logMsg := fmt.Sprintf("occurrence %d created successfully", idOccurrence)
	err = insertOccurrenceLogById(ctx, idOccurrence, logMsg)
	if err != nil {
		return err
	}

	return nil
}

func GetOccurreceByIdHandler(ctx *context.Context, req *pb.Id) (*pb.OccurrenceDetails, error) {
	return selectOccurrenceById(req.GetId())
}

func GetAllOccurrences(ctx *context.Context) (*pb.OccurrenceDetailsList, error) {
	return selectAllOccurrences()
}

func AcceptOccurenceByIdHandler(ctx *context.Context, req *pb.UpdateOccurrenceStatus) (*pb.StatusResponse, error) {
	query := `
	UPDATE occurrence_tb
	SET 
		id_firefighter = $1,
		is_accepted = TRUE,
		accepted_at = NOW(),
		updated_at = NOW()
	WHERE id = $2;`

	db := common.Database
	_, err := db.Exec(query, req.GetIdFirefighter(), req.GetIdOccurrence())
	if err != nil {
		common.LogError(err.Error())
		return nil, err
	}

	logMsg := fmt.Sprintf("occurrence %d was accepted by firefighter with id %d", req.GetIdOccurrence(), req.GetIdFirefighter())
	err = insertOccurrenceLogById(ctx, req.GetIdOccurrence(), logMsg)
	if err != nil {
		common.LogError(err.Error())
		return nil, err
	}

	return &pb.StatusResponse{
		Message: "occurrence accepted successfully!",
	}, nil
}
