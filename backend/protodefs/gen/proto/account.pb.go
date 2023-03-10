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

type CreateAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullName      string                 `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Cpf           string                 `protobuf:"bytes,3,opt,name=cpf,proto3" json:"cpf,omitempty"`
	BirthDate     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	Passwd        string                 `protobuf:"bytes,5,opt,name=passwd,proto3" json:"passwd,omitempty"`
	TwoFactorCode string                 `protobuf:"bytes,6,opt,name=two_factor_code,json=twoFactorCode,proto3" json:"two_factor_code,omitempty"`
	Location      *CreateAccountLocation `protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *CreateAccount) Reset() {
	*x = CreateAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccount) ProtoMessage() {}

func (x *CreateAccount) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateAccount.ProtoReflect.Descriptor instead.
func (*CreateAccount) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccount) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateAccount) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateAccount) GetCpf() string {
	if x != nil {
		return x.Cpf
	}
	return ""
}

func (x *CreateAccount) GetBirthDate() *timestamppb.Timestamp {
	if x != nil {
		return x.BirthDate
	}
	return nil
}

func (x *CreateAccount) GetPasswd() string {
	if x != nil {
		return x.Passwd
	}
	return ""
}

func (x *CreateAccount) GetTwoFactorCode() string {
	if x != nil {
		return x.TwoFactorCode
	}
	return ""
}

func (x *CreateAccount) GetLocation() *CreateAccountLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

type CreateAccountLocation struct {
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

func (x *CreateAccountLocation) Reset() {
	*x = CreateAccountLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountLocation) ProtoMessage() {}

func (x *CreateAccountLocation) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateAccountLocation.ProtoReflect.Descriptor instead.
func (*CreateAccountLocation) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccountLocation) GetCep() string {
	if x != nil {
		return x.Cep
	}
	return ""
}

func (x *CreateAccountLocation) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *CreateAccountLocation) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *CreateAccountLocation) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *CreateAccountLocation) GetNeighborhood() string {
	if x != nil {
		return x.Neighborhood
	}
	return ""
}

func (x *CreateAccountLocation) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *CreateAccountLocation) GetPlaceNumber() int64 {
	if x != nil {
		return x.PlaceNumber
	}
	return 0
}

func (x *CreateAccountLocation) GetComplement() string {
	if x != nil {
		return x.Complement
	}
	return ""
}

type AccountDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName      string                  `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email         string                  `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Cpf           string                  `protobuf:"bytes,4,opt,name=cpf,proto3" json:"cpf,omitempty"`
	BirthDate     *timestamppb.Timestamp  `protobuf:"bytes,5,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	Passwd        string                  `protobuf:"bytes,6,opt,name=passwd,proto3" json:"passwd,omitempty"`
	TwoFactorCode string                  `protobuf:"bytes,7,opt,name=two_factor_code,json=twoFactorCode,proto3" json:"two_factor_code,omitempty"`
	Location      *AccountLocationDetails `protobuf:"bytes,8,opt,name=location,proto3" json:"location,omitempty"`
	CreatedAt     *timestamppb.Timestamp  `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp  `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	SoftDeleted   bool                    `protobuf:"varint,11,opt,name=soft_deleted,json=softDeleted,proto3" json:"soft_deleted,omitempty"`
}

func (x *AccountDetails) Reset() {
	*x = AccountDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountDetails) ProtoMessage() {}

func (x *AccountDetails) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AccountDetails.ProtoReflect.Descriptor instead.
func (*AccountDetails) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{2}
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

func (x *AccountDetails) GetLocation() *AccountLocationDetails {
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

type AccountLocationDetails struct {
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

func (x *AccountLocationDetails) Reset() {
	*x = AccountLocationDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountLocationDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountLocationDetails) ProtoMessage() {}

func (x *AccountLocationDetails) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AccountLocationDetails.ProtoReflect.Descriptor instead.
func (*AccountLocationDetails) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{3}
}

func (x *AccountLocationDetails) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AccountLocationDetails) GetIdAccount() int64 {
	if x != nil {
		return x.IdAccount
	}
	return 0
}

func (x *AccountLocationDetails) GetCep() string {
	if x != nil {
		return x.Cep
	}
	return ""
}

func (x *AccountLocationDetails) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *AccountLocationDetails) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *AccountLocationDetails) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *AccountLocationDetails) GetNeighborhood() string {
	if x != nil {
		return x.Neighborhood
	}
	return ""
}

func (x *AccountLocationDetails) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *AccountLocationDetails) GetPlaceNumber() int64 {
	if x != nil {
		return x.PlaceNumber
	}
	return 0
}

func (x *AccountLocationDetails) GetComplement() string {
	if x != nil {
		return x.Complement
	}
	return ""
}

type EditAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName  string                 `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email     string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Cpf       string                 `protobuf:"bytes,4,opt,name=cpf,proto3" json:"cpf,omitempty"`
	BirthDate *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	Passwd    string                 `protobuf:"bytes,6,opt,name=passwd,proto3" json:"passwd,omitempty"`
	Location  *EditAccountLocation   `protobuf:"bytes,8,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *EditAccount) Reset() {
	*x = EditAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditAccount) ProtoMessage() {}

func (x *EditAccount) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditAccount.ProtoReflect.Descriptor instead.
func (*EditAccount) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{4}
}

func (x *EditAccount) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EditAccount) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *EditAccount) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EditAccount) GetCpf() string {
	if x != nil {
		return x.Cpf
	}
	return ""
}

func (x *EditAccount) GetBirthDate() *timestamppb.Timestamp {
	if x != nil {
		return x.BirthDate
	}
	return nil
}

func (x *EditAccount) GetPasswd() string {
	if x != nil {
		return x.Passwd
	}
	return ""
}

func (x *EditAccount) GetLocation() *EditAccountLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

type EditAccountLocation struct {
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

func (x *EditAccountLocation) Reset() {
	*x = EditAccountLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditAccountLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditAccountLocation) ProtoMessage() {}

func (x *EditAccountLocation) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditAccountLocation.ProtoReflect.Descriptor instead.
func (*EditAccountLocation) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{5}
}

func (x *EditAccountLocation) GetCep() string {
	if x != nil {
		return x.Cep
	}
	return ""
}

func (x *EditAccountLocation) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *EditAccountLocation) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *EditAccountLocation) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *EditAccountLocation) GetNeighborhood() string {
	if x != nil {
		return x.Neighborhood
	}
	return ""
}

func (x *EditAccountLocation) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *EditAccountLocation) GetPlaceNumber() int64 {
	if x != nil {
		return x.PlaceNumber
	}
	return 0
}

func (x *EditAccountLocation) GetComplement() string {
	if x != nil {
		return x.Complement
	}
	return ""
}

var File_account_proto protoreflect.FileDescriptor

var file_account_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x02, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70,
	0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x66, 0x12, 0x39, 0x0a, 0x0a,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x73, 0x73, 0x77, 0x64, 0x12,
	0x26, 0x0a, 0x0f, 0x74, 0x77, 0x6f, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x77, 0x6f, 0x46, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xec, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x63,
	0x65, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x65, 0x70, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f,
	0x72, 0x68, 0x6f, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x22, 0xb4, 0x03, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x66, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x66, 0x12, 0x39, 0x0a, 0x0a, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x73, 0x73, 0x77, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x73, 0x73, 0x77, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x74,
	0x77, 0x6f, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x77, 0x6f, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6f, 0x66, 0x74, 0x5f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x6f, 0x66, 0x74,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x9c, 0x02, 0x0a, 0x16, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x65, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x63, 0x65, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x65, 0x69, 0x67, 0x68,
	0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e,
	0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72,
	0x65, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xed, 0x01, 0x0a, 0x0b, 0x45, 0x64, 0x69, 0x74, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x66,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x66, 0x12, 0x39, 0x0a, 0x0a, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x62, 0x69, 0x72,
	0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x73, 0x73, 0x77, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x73, 0x73, 0x77, 0x64, 0x12, 0x36,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xea, 0x01, 0x0a, 0x13, 0x45, 0x64, 0x69, 0x74, 0x41,
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
	0x65, 0x6e, 0x74, 0x32, 0x97, 0x02, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x2b, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x09, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x49, 0x64, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x45, 0x64,
	0x69, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x64, 0x69, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x32, 0x0a, 0x0e, 0x53, 0x6f, 0x66, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x09, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x64, 0x1a, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x09, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x64, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_account_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_account_proto_goTypes = []interface{}{
	(*CreateAccount)(nil),          // 0: proto.CreateAccount
	(*CreateAccountLocation)(nil),  // 1: proto.CreateAccountLocation
	(*AccountDetails)(nil),         // 2: proto.AccountDetails
	(*AccountLocationDetails)(nil), // 3: proto.AccountLocationDetails
	(*EditAccount)(nil),            // 4: proto.EditAccount
	(*EditAccountLocation)(nil),    // 5: proto.EditAccountLocation
	(*timestamppb.Timestamp)(nil),  // 6: google.protobuf.Timestamp
	(*Id)(nil),                     // 7: proto.Id
	(*StatusResponse)(nil),         // 8: proto.StatusResponse
}
var file_account_proto_depIdxs = []int32{
	6,  // 0: proto.CreateAccount.birth_date:type_name -> google.protobuf.Timestamp
	1,  // 1: proto.CreateAccount.location:type_name -> proto.CreateAccountLocation
	6,  // 2: proto.AccountDetails.birth_date:type_name -> google.protobuf.Timestamp
	3,  // 3: proto.AccountDetails.location:type_name -> proto.AccountLocationDetails
	6,  // 4: proto.AccountDetails.created_at:type_name -> google.protobuf.Timestamp
	6,  // 5: proto.AccountDetails.updated_at:type_name -> google.protobuf.Timestamp
	6,  // 6: proto.EditAccount.birth_date:type_name -> google.protobuf.Timestamp
	5,  // 7: proto.EditAccount.location:type_name -> proto.EditAccountLocation
	0,  // 8: proto.AccountService.Create:input_type -> proto.CreateAccount
	7,  // 9: proto.AccountService.GetById:input_type -> proto.Id
	4,  // 10: proto.AccountService.EditById:input_type -> proto.EditAccount
	7,  // 11: proto.AccountService.SoftDeleteById:input_type -> proto.Id
	7,  // 12: proto.AccountService.RecoverAccountById:input_type -> proto.Id
	2,  // 13: proto.AccountService.Create:output_type -> proto.AccountDetails
	2,  // 14: proto.AccountService.GetById:output_type -> proto.AccountDetails
	2,  // 15: proto.AccountService.EditById:output_type -> proto.AccountDetails
	8,  // 16: proto.AccountService.SoftDeleteById:output_type -> proto.StatusResponse
	8,  // 17: proto.AccountService.RecoverAccountById:output_type -> proto.StatusResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_account_proto_init() }
func file_account_proto_init() {
	if File_account_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccount); i {
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
			switch v := v.(*CreateAccountLocation); i {
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
		file_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountLocationDetails); i {
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
		file_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditAccount); i {
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
		file_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditAccountLocation); i {
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
			NumMessages:   6,
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
