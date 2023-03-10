package responses

import (
	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
)

func NewCreateOccurrenceModelFromJSONToProto(model *CreateOccurrence) *pb.CreateOccurrence {
	return &pb.CreateOccurrence{
		IdFirefighter: model.IdFirefighter,
		IdAuthor:      model.IdAuthor,
		BannerX64:     model.BannerX64,
		Title:         model.Title,
		Description:   model.Description,
		CreatedAt:     common.TimeToTimestamp(model.CreatedAt),
		Location: &pb.CreateOccurrenceLocation{
			Cep:          model.Location.CEP,
			Country:      model.Location.Country,
			State:        model.Location.State,
			City:         model.Location.City,
			Neighborhood: model.Location.Neighborhood,
			Street:       model.Location.Street,
			PlaceNumber:  model.Location.PlaceNumber,
			Complement:   model.Location.Complement,
			Latitude:     model.Location.Latitude,
			Longitude:    model.Location.Longitude,
		},
	}
}
func NewCreateOccurrenceModelFromProtoToJSON(model *pb.CreateOccurrence) *CreateOccurrence {
	return &CreateOccurrence{
		IdAuthor:    model.GetIdAuthor(),
		Title:       model.GetTitle(),
		Description: model.GetDescription(),
		CreatedAt:   common.TimestampToTime(model.GetCreatedAt()),
		BannerX64:   model.GetBannerX64(),
		Location: CreateOccurrenceLocation{
			CEP:          model.Location.GetCep(),
			Country:      model.Location.GetCountry(),
			State:        model.Location.GetState(),
			City:         model.Location.GetCity(),
			Neighborhood: model.Location.GetNeighborhood(),
			Street:       model.Location.GetStreet(),
			PlaceNumber:  model.Location.GetPlaceNumber(),
			Complement:   model.Location.GetComplement(),
			Latitude:     model.Location.GetLatitude(),
			Longitude:    model.Location.GetLongitude(),
		},
	}
}

func NewOccurrenceDetailsModelFromJSONToProto(model *OccurrenceDetails) *pb.OccurrenceDetails {
	return &pb.OccurrenceDetails{
		IdOccurrence:        model.IdOccurrence,
		Title:               model.Title,
		Description:         model.Description,
		BannerX64:           model.BannerX64,
		IdFirefighter:       model.IdFirefighter,
		FirefighterFullName: model.FirefighterFullname,
		FirefighterEmail:    model.FirefighterEmail,
		IdAuthor:            model.IdAuthor,
		AuthorFullName:      model.AuthorFullName,
		AuthorEmail:         model.AuthorEmail,
		CreatedAt:           common.TimeToTimestamp(model.CreatedAt),
		UpdatedAt:           common.TimeToTimestamp(model.UpdatedAt),
		AcceptedAt:          common.TimeToTimestamp(model.AcceptedAt),
		CanceledAt:          common.TimeToTimestamp(model.CanceledAt),
		FinishedAt:          common.TimeToTimestamp(model.FinishedAt),
		IsAccepted:          model.IsAccepted,
		IsCanceled:          model.IsCanceled,
		IsFinished:          model.IsFinished,
		Location: &pb.OccurrenceDetailsLocation{
			Id:           model.Location.Id,
			Cep:          model.Location.CEP,
			Country:      model.Location.Country,
			State:        model.Location.State,
			City:         model.Location.City,
			Neighborhood: model.Location.Neighborhood,
			Street:       model.Location.Street,
			PlaceNumber:  model.Location.PlaceNumber,
			Complement:   model.Location.Complement,
			Latitude:     model.Location.Latitude,
			Longitude:    model.Location.Longitude,
		},
	}
}
func NewOccurrenceDetailsModelFromProtoToJSON(model *pb.OccurrenceDetails) *OccurrenceDetails {
	return &OccurrenceDetails{
		IdOccurrence:        model.GetIdOccurrence(),
		Title:               model.GetTitle(),
		Description:         model.GetDescription(),
		BannerX64:           model.GetBannerX64(),
		IdFirefighter:       model.GetIdFirefighter(),
		FirefighterFullname: model.GetFirefighterFullName(),
		FirefighterEmail:    model.GetFirefighterEmail(),
		IdAuthor:            model.GetIdAuthor(),
		AuthorFullName:      model.GetAuthorFullName(),
		AuthorEmail:         model.GetAuthorEmail(),
		CreatedAt:           common.TimestampToTime(model.GetCreatedAt()),
		UpdatedAt:           common.TimestampToTime(model.GetUpdatedAt()),
		AcceptedAt:          common.TimestampToTime(model.GetAcceptedAt()),
		CanceledAt:          common.TimestampToTime(model.GetCanceledAt()),
		FinishedAt:          common.TimestampToTime(model.GetFinishedAt()),
		IsAccepted:          model.GetIsAccepted(),
		IsCanceled:          model.GetIsCanceled(),
		IsFinished:          model.GetIsFinished(),
		Location: OccurrenceDetailsLocation{
			Id:           model.Location.GetId(),
			CEP:          model.Location.GetCep(),
			Country:      model.Location.GetCountry(),
			State:        model.Location.GetState(),
			City:         model.Location.GetCity(),
			Neighborhood: model.Location.GetNeighborhood(),
			Street:       model.Location.GetStreet(),
			PlaceNumber:  model.Location.GetPlaceNumber(),
			Complement:   model.Location.GetComplement(),
			Latitude:     model.Location.GetLatitude(),
			Longitude:    model.Location.GetLongitude(),
		},
	}
}

func NewEditOccurrenceModelFromJSONToProto(model *EditOccurrence) *pb.EditOccurrence {
	return &pb.EditOccurrence{
		Id:          model.Id,
		IdAuthor:    model.IdAuthor,
		BannerX64:   model.BannerX64,
		Title:       model.Title,
		Description: model.Description,
		Location: &pb.EditOccurrenceLocation{
			Id:           model.Location.Id,
			Cep:          model.Location.CEP,
			Country:      model.Location.Country,
			State:        model.Location.State,
			City:         model.Location.City,
			Neighborhood: model.Location.Neighborhood,
			Street:       model.Location.Street,
			PlaceNumber:  model.Location.PlaceNumber,
			Complement:   model.Location.Complement,
			Latitude:     model.Location.Latitude,
			Longitude:    model.Location.Longitude,
		},
	}
}
func NewEditOccurrenceModelFromProtoToJSON(model *pb.EditOccurrence) *EditOccurrence {
	return &EditOccurrence{
		Id:          model.GetId(),
		IdAuthor:    model.GetIdAuthor(),
		BannerX64:   model.GetBannerX64(),
		Title:       model.GetTitle(),
		Description: model.GetDescription(),
		Location: EditOccurrenceLocation{
			Id:           model.Location.GetId(),
			CEP:          model.Location.GetCep(),
			Country:      model.Location.GetCountry(),
			State:        model.Location.GetState(),
			City:         model.Location.GetCity(),
			Neighborhood: model.Location.GetNeighborhood(),
			Street:       model.Location.GetStreet(),
			PlaceNumber:  model.Location.GetPlaceNumber(),
			Complement:   model.Location.GetComplement(),
			Latitude:     model.Location.GetLatitude(),
			Longitude:    model.Location.GetLongitude(),
		},
	}
}

func NewUpdateOccurrenceStatusByIdModelFromJSONToProto(model *UpdateOccurrenceStatus) *pb.UpdateOccurrenceStatus {
	return &pb.UpdateOccurrenceStatus{
		IdOccurrence:  model.IdOccurrence,
		IdFirefighter: model.IdFirefighter,
	}
}
func NewUpdateOccurrenceStatusByIdModelFromProtoToJSON(model *pb.UpdateOccurrenceStatus) *UpdateOccurrenceStatus {
	return &UpdateOccurrenceStatus{
		IdOccurrence:  model.GetIdOccurrence(),
		IdFirefighter: model.GetIdFirefighter(),
	}
}
