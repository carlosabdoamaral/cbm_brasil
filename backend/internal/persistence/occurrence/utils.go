package occurrence

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/responses"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
)

const selectOccurrenceDetailsQuery = `
	SELECT
		occurrence_tb.id as occurrence_id,
		occurrence_tb.title as occurrence_title,
		occurrence_tb.description as occurrence_description,
		occurrence_tb.banner_x64 as banner_x64,

		occurrence_tb.created_at as occurrence_created_at,
		occurrence_tb.updated_at as occurrence_updated_at,
		occurrence_tb.accepted_at as occurrence_accepted_at,
		occurrence_tb.canceled_at as occurrence_canceled_at,
		occurrence_tb.finished_at as occurrence_finished_at,

		occurrence_tb.soft_deleted as occurrence_is_soft_deleted,
		occurrence_tb.is_accepted as occurrence_is_accepted,
		occurrence_tb.is_canceled as occurrence_is_canceled,
		occurrence_tb.is_finished as occurrence_is_finished,
		
		firefighter_account_tb.id as firefighter_id,
		firefighter_account_tb.full_name as firefighter_name,
		firefighter_account_tb.email as firefighter_email,
		
		account_tb.id as author_id,
		account_tb.full_name as author_name,
		account_tb.email as author_email,
		
		occurrence_location_tb.id as location_id,
		occurrence_location_tb.cep as location_cep,
		occurrence_location_tb.country as location_country,
		occurrence_location_tb.state as location_state,
		occurrence_location_tb.city as location_city,
		occurrence_location_tb.neighborhood as location_neighborhood,
		occurrence_location_tb.street as location_street,
		occurrence_location_tb.place_number as location_place_number,
		occurrence_location_tb.complement as location_complement,
		occurrence_location_tb.latitude as location_latitude,
		occurrence_location_tb.longitude as location_longitude
	FROM occurrence_tb
	INNER JOIN occurrence_location_tb ON occurrence_tb.id = occurrence_location_tb.id_occurrence
	INNER JOIN firefighter_account_tb ON firefighter_account_tb.id = occurrence_tb.id_firefighter
	INNER JOIN account_tb ON account_tb.id = occurrence_tb.id_author
`

func scanOccurrenceDetailsRows(rows *sql.Rows) (*pb.OccurrenceDetails, error) {
	jsonModel := &responses.OccurrenceDetails{}
	err := rows.Scan(
		&jsonModel.IdOccurrence,
		&jsonModel.Title,
		&jsonModel.Description,
		&jsonModel.BannerX64,

		&jsonModel.CreatedAt,
		&jsonModel.UpdatedAt,
		&jsonModel.AcceptedAt,
		&jsonModel.CanceledAt,
		&jsonModel.FinishedAt,

		&jsonModel.SoftDeleted,
		&jsonModel.IsAccepted,
		&jsonModel.IsCanceled,
		&jsonModel.IsFinished,

		&jsonModel.IdFirefighter,
		&jsonModel.FirefighterFullname,
		&jsonModel.FirefighterEmail,

		&jsonModel.IdAuthor,
		&jsonModel.AuthorFullName,
		&jsonModel.AuthorEmail,

		&jsonModel.Location.Id,
		&jsonModel.Location.CEP,
		&jsonModel.Location.Country,
		&jsonModel.Location.State,
		&jsonModel.Location.City,
		&jsonModel.Location.Neighborhood,
		&jsonModel.Location.Street,
		&jsonModel.Location.PlaceNumber,
		&jsonModel.Location.Complement,
		&jsonModel.Location.Latitude,
		&jsonModel.Location.Longitude,
	)
	if err != nil {
		common.LogError(err.Error())
		return nil, err
	}

	return responses.NewOccurrenceDetailsModelFromJSONToProto(jsonModel), nil
}

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
		common.LogError(err.Error())
		return 0, err
	}

	for rows.Next() {
		err = rows.Scan(&idOccurrence)
		if err != nil {
			common.LogError(err.Error())
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

func selectOccurrenceById(id int64) (*pb.OccurrenceDetails, error) {
	query := fmt.Sprintf("%s %s", selectOccurrenceDetailsQuery, "WHERE occurrence_tb.id = $1")

	db := common.Database

	rows, err := db.Query(query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no rows founded")
		}

		common.LogError(err.Error())
		return nil, err
	}

	res := &pb.OccurrenceDetails{}
	for rows.Next() {
		res, err = scanOccurrenceDetailsRows(rows)
		if err != nil {
			common.LogError(err.Error())
			return nil, err
		}
	}

	return res, nil
}

func selectAllOccurrences() (*pb.OccurrenceDetailsList, error) {
	query := selectOccurrenceDetailsQuery

	db := common.Database

	rows, err := db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no rows founded")
		}

		common.LogError(err.Error())
		return nil, err
	}

	res := &pb.OccurrenceDetailsList{}
	for rows.Next() {
		scanRes, err := scanOccurrenceDetailsRows(rows)
		if err != nil {
			common.LogError(err.Error())
			return nil, err
		}

		res.List = append(res.List, scanRes)
	}

	return res, nil
}
