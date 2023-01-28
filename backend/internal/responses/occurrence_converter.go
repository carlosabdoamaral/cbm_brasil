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
		IdAuthor:      model.IdAuthor,
		IdFirefighter: model.IdFirefighter,
		BannerX64:     model.BannerX64,
		Title:         model.Title,
		Description:   model.Description,
		CreatedAt:     common.TimeToTimestamp(model.CreatedAt),
		AcceptedAt:    common.TimeToTimestamp(model.AcceptedAt),
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
		SoftDeleted: false,
	}
}
func NewOccurrenceDetailsModelFromProtoToJSON(model *pb.OccurrenceDetails) *OccurrenceDetails {
	return &OccurrenceDetails{
		IdAuthor:      model.GetIdAuthor(),
		IdFirefighter: model.GetIdFirefighter(),
		Title:         model.GetTitle(),
		Description:   model.GetDescription(),
		CreatedAt:     common.TimestampToTime(model.GetCreatedAt()),
		UpdatedAt:     common.TimestampToTime(model.GetUpdatedAt()),
		AcceptedAt:    common.TimestampToTime(model.GetAcceptedAt()),
		BannerX64:     model.GetBannerX64(),
		SoftDeleted:   model.GetSoftDeleted(),
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
