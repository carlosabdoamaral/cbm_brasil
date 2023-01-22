// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: account.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NewAccountLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cep          string `protobuf:"bytes,1,opt,name=cep,proto3" json:"cep,omitempty"`
	Country      string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	State        string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	City         string `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Neighborhood string `protobuf:"bytes,5,opt,name=neighborhood,proto3" json:"neighborhood,omitempty"`
	Street       string `protobuf:"bytes,6,opt,name=street,proto3" json:"street,omitempty"`
	PlaceNumber  int64  `protobuf:"varint,7,opt,name=place_number,json=placeNumber,proto3" json:"place_number,omitempty"`
	Complement   string `protobuf:"bytes,8,opt,name=complement,proto3" json:"complement,omitempty"`
}

func (x *NewAccountLocation) Reset() {
	*x = NewAccountLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewAccountLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewAccountLocation) ProtoMessage() {}

func (x *NewAccountLocation) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewAccountLocation.ProtoReflect.Descriptor instead.
func (*NewAccountLocation) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{0}
}

func (x *NewAccountLocation) GetCep() string {
	if x != nil {
		return x.Cep
	}
	return ""
}

func (x *NewAccountLocation) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *NewAccountLocation) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *NewAccountLocation) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *NewAccountLocation) GetNeighborhood() string {
	if x != nil {
		return x.Neighborhood
	}
	return ""
}

func (x *NewAccountLocation) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *NewAccountLocation) GetPlaceNumber() int64 {
	if x != nil {
		return x.PlaceNumber
	}
	return 0
}

func (x *NewAccountLocation) GetComplement() string {
	if x != nil {
		return x.Complement
	}
	return ""
}

type NewAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullName      string                 `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Cpf           string                 `protobuf:"bytes,3,opt,name=cpf,proto3" json:"cpf,omitempty"`
	BirthDate     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	Passwd        string                 `protobuf:"bytes,5,opt,name=passwd,proto3" json:"passwd,omitempty"`
	TwoFactorCode string                 `protobuf:"bytes,6,opt,name=two_factor_code,json=twoFactorCode,proto3" json:"two_factor_code,omitempty"`
	Location      *NewAccountLocation    `protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *NewAccountRequest) Reset() {
	*x = NewAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewAccountRequest) ProtoMessage() {}

func (x *NewAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewAccountRequest.ProtoReflect.Descriptor instead.
func (*NewAccountRequest) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{1}
}

func (x *NewAccountRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *NewAccountRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *NewAccountRequest) GetCpf() string {
	if x != nil {
		return x.Cpf
	}
	return ""
}

func (x *NewAccountRequest) GetBirthDate() *timestamppb.Timestamp {
	if x != nil {
		return x.BirthDate
	}
	return nil
}

func (x *NewAccountRequest) GetPasswd() string {
	if x != nil {
		return x.Passwd
	}
	return ""
}

func (x *NewAccountRequest) GetTwoFactorCode() string {
	if x != nil {
		return x.TwoFactorCode
	}
	return ""
}

func (x *NewAccountRequest) GetLocation() *NewAccountLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

type LocationDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IdAccount    int64  `protobuf:"varint,2,opt,name=id_account,json=idAccount,proto3" json:"id_account,omitempty"`
	Cep          string `protobuf:"bytes,3,opt,name=cep,proto3" json:"cep,omitempty"`
	Country      string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	State        string `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	City         string `protobuf:"bytes,6,opt,name=city,proto3" json:"city,omitempty"`
	Neighborhood string `protobuf:"bytes,7,opt,name=neighborhood,proto3" json:"neighborhood,omitempty"`
	Street       string `protobuf:"bytes,8,opt,name=street,proto3" json:"street,omitempty"`
	PlaceNumber  int64  `protobuf:"varint,9,opt,name=place_number,json=placeNumber,proto3" json:"place_number,omitempty"`
	Complement   string `protobuf:"bytes,10,opt,name=complement,proto3" json:"complement,omitempty"`
}

func (x *LocationDetails) Reset() {
	*x = LocationDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationDetails) ProtoMessage() {}

func (x *LocationDetails) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationDetails.ProtoReflect.Descriptor instead.
func (*LocationDetails) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{2}
}

func (x *LocationDetails) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LocationDetails) GetIdAccount() int64 {
	if x != nil {
		return x.IdAccount
	}
	return 0
}

func (x *LocationDetails) GetCep() string {
	if x != nil {
		return x.Cep
	}
	return ""
}

func (x *LocationDetails) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *LocationDetails) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *LocationDetails) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *LocationDetails) GetNeighborhood() string {
	if x != nil {
		return x.Neighborhood
	}
	return ""
}

func (x *LocationDetails) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *LocationDetails) GetPlaceNumber() int64 {
	if x != nil {
		return x.PlaceNumber
	}
	return 0
}

func (x *LocationDetails) GetComplement() string {
	if x != nil {
		return x.Complement
	}
	return ""
}

type AccountDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName      string                 `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Cpf           string                 `protobuf:"bytes,4,opt,name=cpf,proto3" json:"cpf,omitempty"`
	BirthDate     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	Passwd        string                 `protobuf:"bytes,6,opt,name=passwd,proto3" json:"passwd,omitempty"`
	TwoFactorCode string                 `protobuf:"bytes,7,opt,name=two_factor_code,json=twoFactorCode,proto3" json:"two_factor_code,omitempty"`
	Location      *LocationDetails       `protobuf:"bytes,8,opt,name=location,proto3" json:"location,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	SoftDeleted   bool                   `protobuf:"varint,11,opt,name=soft_deleted,json=softDeleted,proto3" json:"soft_deleted,omitempty"`
}

func (x *AccountDetails) Reset() {
	*x = AccountDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountDetails) ProtoMessage() {}

func (x *AccountDetails) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountDetails.ProtoReflect.Descriptor instead.
func (*AccountDetails) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{3}
}

func (x *AccountDetails) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AccountDetails) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *AccountDetails) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AccountDetails) GetCpf() string {
	if x != nil {
		return x.Cpf
	}
	return ""
}

func (x *AccountDetails) GetBirthDate() *timestamppb.Timestamp {
	if x != nil {
		return x.BirthDate
	}
	return nil
}

func (x *AccountDetails) GetPasswd() string {
	if x != nil {
		return x.Passwd
	}
	return ""
}

func (x *AccountDetails) GetTwoFactorCode() string {
	if x != nil {
		return x.TwoFactorCode
	}
	return ""
}

func (x *AccountDetails) GetLocation() *LocationDetails {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *AccountDetails) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *AccountDetails) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *AccountDetails) GetSoftDeleted() bool {
	if x != nil {
		return x.SoftDeleted
	}
	return false
}

var File_account_proto protoreflect.FileDescriptor

var file_account_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x01, 0x0a, 0x12, 0x4e, 0x65, 0x77, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x63, 0x65, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x65, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72,
	0x68, 0x6f, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x69, 0x67,
	0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0x8a, 0x02, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c,
	0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75,
	0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03,
	0x63, 0x70, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x66, 0x12, 0x39,
	0x0a, 0x0a, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x64, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x77, 0x6f, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x77, 0x6f, 0x46,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x95, 0x02, 0x0a, 0x0f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x64, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x65, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x63, 0x65, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x65, 0x69,
	0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xad, 0x03, 0x0a, 0x0e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10,
	0x0a, 0x03, 0x63, 0x70, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x66,
	0x12, 0x39, 0x0a, 0x0a, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x77, 0x6f, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x77,
	0x6f, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6f, 0x66, 0x74, 0x5f, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x6f, 0x66,
	0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x32, 0x4b, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_proto_rawDescOnce sync.Once
	file_account_proto_rawDescData = file_account_proto_rawDesc
)

func file_account_proto_rawDescGZIP() []byte {
	file_account_proto_rawDescOnce.Do(func() {
		file_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_proto_rawDescData)
	})
	return file_account_proto_rawDescData
}

var file_account_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_account_proto_goTypes = []interface{}{
	(*NewAccountLocation)(nil),    // 0: proto.NewAccountLocation
	(*NewAccountRequest)(nil),     // 1: proto.NewAccountRequest
	(*LocationDetails)(nil),       // 2: proto.LocationDetails
	(*AccountDetails)(nil),        // 3: proto.AccountDetails
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_account_proto_depIdxs = []int32{
	4, // 0: proto.NewAccountRequest.birth_date:type_name -> google.protobuf.Timestamp
	0, // 1: proto.NewAccountRequest.location:type_name -> proto.NewAccountLocation
	4, // 2: proto.AccountDetails.birth_date:type_name -> google.protobuf.Timestamp
	2, // 3: proto.AccountDetails.location:type_name -> proto.LocationDetails
	4, // 4: proto.AccountDetails.created_at:type_name -> google.protobuf.Timestamp
	4, // 5: proto.AccountDetails.updated_at:type_name -> google.protobuf.Timestamp
	1, // 6: proto.AccountService.Create:input_type -> proto.NewAccountRequest
	3, // 7: proto.AccountService.Create:output_type -> proto.AccountDetails
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_account_proto_init() }
func file_account_proto_init() {
	if File_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewAccountLocation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewAccountRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationDetails); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountDetails); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_proto_goTypes,
		DependencyIndexes: file_account_proto_depIdxs,
		MessageInfos:      file_account_proto_msgTypes,
	}.Build()
	File_account_proto = out.File
	file_account_proto_rawDesc = nil
	file_account_proto_goTypes = nil
	file_account_proto_depIdxs = nil
}
