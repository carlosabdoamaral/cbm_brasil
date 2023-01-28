package occurrence

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
)

func insertOccurrence(ctx *context.Context, req *pb.CreateOccurrence) (idOccurrence int64, err error) {
	var (
		db    *sql.DB = common.Database
		query string  = `
		INSERT INTO occurrence_tb(id_author, id_firefighter, banner_x64, title, description)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;`
	)

	rows, err := db.Query(
		query,
		req.GetIdAuthor(),
		req.GetIdFirefighter(),
		req.GetBannerX64(),
		req.GetTitle(),
		req.GetDescription(),
	)
	if err != nil {
		return 0, err
	}

	for rows.Next() {
		err = rows.Scan(&idOccurrence)
		if err != nil {
			return 0, err
		}
	}

	return
}

func insertOccurrenceLocationById(ctx *context.Context, idOccurrence int64, location *pb.CreateOccurrenceLocation) (err error) {
	var (
		db    *sql.DB = common.Database
		query string  = `
		INSERT INTO occurrence_location_tb(id_occurrence, cep, country, state, city, neighborhood, street, place_number, complement, latitude, longitude)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`
	)

	_, err = db.Exec(
		query,
		idOccurrence,
		location.GetCep(),
		location.GetCountry(),
		location.GetState(),
		location.GetCity(),
		location.GetNeighborhood(),
		location.GetStreet(),
		location.GetPlaceNumber(),
		location.GetComplement(),
		location.GetLatitude(),
		location.GetLongitude(),
	)
	if err != nil {
		return err
	}

	return nil
}

func insertOccurrenceLogById(ctx *context.Context, id int64, log string) (err error) {
	var (
		db    *sql.DB = common.Database
		query string  = `INSERT INTO occurrence_logs_tb(id_occurrence, msg) VALUES ($1, $2);`
	)

	_, err = db.Exec(query, id, log)
	if err != nil {
		return err
	}

	return nil
}

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
