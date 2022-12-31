// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: occurrence.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OccurrenceDetailsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Occurrences []*OccurrenceDetails `protobuf:"bytes,1,rep,name=occurrences,proto3" json:"occurrences,omitempty"`
}

func (x *OccurrenceDetailsList) Reset() {
	*x = OccurrenceDetailsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_occurrence_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OccurrenceDetailsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OccurrenceDetailsList) ProtoMessage() {}

func (x *OccurrenceDetailsList) ProtoReflect() protoreflect.Message {
	mi := &file_occurrence_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OccurrenceDetailsList.ProtoReflect.Descriptor instead.
func (*OccurrenceDetailsList) Descriptor() ([]byte, []int) {
	return file_occurrence_proto_rawDescGZIP(), []int{0}
}

func (x *OccurrenceDetailsList) GetOccurrences() []*OccurrenceDetails {
	if x != nil {
		return x.Occurrences
	}
	return nil
}

type NewOccurrence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author      string           `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Title       string           `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string           `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	ImageBase64 string           `protobuf:"bytes,5,opt,name=image_base64,json=imageBase64,proto3" json:"image_base64,omitempty"`
	Location    *Location        `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	Tags        []*OccurrenceTag `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *NewOccurrence) Reset() {
	*x = NewOccurrence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_occurrence_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewOccurrence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewOccurrence) ProtoMessage() {}

func (x *NewOccurrence) ProtoReflect() protoreflect.Message {
	mi := &file_occurrence_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewOccurrence.ProtoReflect.Descriptor instead.
func (*NewOccurrence) Descriptor() ([]byte, []int) {
	return file_occurrence_proto_rawDescGZIP(), []int{1}
}

func (x *NewOccurrence) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *NewOccurrence) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NewOccurrence) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NewOccurrence) GetImageBase64() string {
	if x != nil {
		return x.ImageBase64
	}
	return ""
}

func (x *NewOccurrence) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *NewOccurrence) GetTags() []*OccurrenceTag {
	if x != nil {
		return x.Tags
	}
	return nil
}

type OccurrenceDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Author      string           `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Title       string           `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string           `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	ImageBase64 string           `protobuf:"bytes,5,opt,name=image_base64,json=imageBase64,proto3" json:"image_base64,omitempty"`
	Location    *Location        `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	Tags        []*OccurrenceTag `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *OccurrenceDetails) Reset() {
	*x = OccurrenceDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_occurrence_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OccurrenceDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OccurrenceDetails) ProtoMessage() {}

func (x *OccurrenceDetails) ProtoReflect() protoreflect.Message {
	mi := &file_occurrence_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OccurrenceDetails.ProtoReflect.Descriptor instead.
func (*OccurrenceDetails) Descriptor() ([]byte, []int) {
	return file_occurrence_proto_rawDescGZIP(), []int{2}
}

func (x *OccurrenceDetails) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OccurrenceDetails) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *OccurrenceDetails) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *OccurrenceDetails) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *OccurrenceDetails) GetImageBase64() string {
	if x != nil {
		return x.ImageBase64
	}
	return ""
}

func (x *OccurrenceDetails) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *OccurrenceDetails) GetTags() []*OccurrenceTag {
	if x != nil {
		return x.Tags
	}
	return nil
}

type OccurrenceTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *OccurrenceTag) Reset() {
	*x = OccurrenceTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_occurrence_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OccurrenceTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OccurrenceTag) ProtoMessage() {}

func (x *OccurrenceTag) ProtoReflect() protoreflect.Message {
	mi := &file_occurrence_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OccurrenceTag.ProtoReflect.Descriptor instead.
func (*OccurrenceTag) Descriptor() ([]byte, []int) {
	return file_occurrence_proto_rawDescGZIP(), []int{3}
}

func (x *OccurrenceTag) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OccurrenceTag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_occurrence_proto protoreflect.FileDescriptor

var file_occurrence_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x15, 0x4f, 0x63, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x3a, 0x0a, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x63,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52,
	0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x22, 0xd9, 0x01, 0x0a,
	0x0d, 0x4e, 0x65, 0x77, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21,
	0x0a, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x61, 0x73, 0x65, 0x36,
	0x34, 0x12, 0x2b, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x54,
	0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0xed, 0x01, 0x0a, 0x11, 0x4f, 0x63, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21,
	0x0a, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x61, 0x73, 0x65, 0x36,
	0x34, 0x12, 0x2b, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x54,
	0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x33, 0x0a, 0x0d, 0x4f, 0x63, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x61, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xe3, 0x02,
	0x0a, 0x11, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x09, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49,
	0x64, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x12, 0x3c, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65,
	0x77, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x30, 0x0a, 0x0a, 0x53, 0x6f, 0x66, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x09, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x64, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_occurrence_proto_rawDescOnce sync.Once
	file_occurrence_proto_rawDescData = file_occurrence_proto_rawDesc
)

func file_occurrence_proto_rawDescGZIP() []byte {
	file_occurrence_proto_rawDescOnce.Do(func() {
		file_occurrence_proto_rawDescData = protoimpl.X.CompressGZIP(file_occurrence_proto_rawDescData)
	})
	return file_occurrence_proto_rawDescData
}

var file_occurrence_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_occurrence_proto_goTypes = []interface{}{
	(*OccurrenceDetailsList)(nil), // 0: proto.OccurrenceDetailsList
	(*NewOccurrence)(nil),         // 1: proto.NewOccurrence
	(*OccurrenceDetails)(nil),     // 2: proto.OccurrenceDetails
	(*OccurrenceTag)(nil),         // 3: proto.OccurrenceTag
	(*Location)(nil),              // 4: proto.Location
	(*Empty)(nil),                 // 5: proto.Empty
	(*Id)(nil),                    // 6: proto.Id
	(*StatusResponse)(nil),        // 7: proto.StatusResponse
}
var file_occurrence_proto_depIdxs = []int32{
	2,  // 0: proto.OccurrenceDetailsList.occurrences:type_name -> proto.OccurrenceDetails
	4,  // 1: proto.NewOccurrence.location:type_name -> proto.Location
	3,  // 2: proto.NewOccurrence.tags:type_name -> proto.OccurrenceTag
	4,  // 3: proto.OccurrenceDetails.location:type_name -> proto.Location
	3,  // 4: proto.OccurrenceDetails.tags:type_name -> proto.OccurrenceTag
	5,  // 5: proto.OccurrenceService.GetAll:input_type -> proto.Empty
	6,  // 6: proto.OccurrenceService.GetById:input_type -> proto.Id
	4,  // 7: proto.OccurrenceService.GetNearby:input_type -> proto.Location
	1,  // 8: proto.OccurrenceService.Create:input_type -> proto.NewOccurrence
	2,  // 9: proto.OccurrenceService.Update:input_type -> proto.OccurrenceDetails
	6,  // 10: proto.OccurrenceService.SoftDelete:input_type -> proto.Id
	0,  // 11: proto.OccurrenceService.GetAll:output_type -> proto.OccurrenceDetailsList
	2,  // 12: proto.OccurrenceService.GetById:output_type -> proto.OccurrenceDetails
	0,  // 13: proto.OccurrenceService.GetNearby:output_type -> proto.OccurrenceDetailsList
	7,  // 14: proto.OccurrenceService.Create:output_type -> proto.StatusResponse
	7,  // 15: proto.OccurrenceService.Update:output_type -> proto.StatusResponse
	7,  // 16: proto.OccurrenceService.SoftDelete:output_type -> proto.StatusResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_occurrence_proto_init() }
func file_occurrence_proto_init() {
	if File_occurrence_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_occurrence_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OccurrenceDetailsList); i {
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
		file_occurrence_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewOccurrence); i {
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
		file_occurrence_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OccurrenceDetails); i {
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
		file_occurrence_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OccurrenceTag); i {
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
			RawDescriptor: file_occurrence_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_occurrence_proto_goTypes,
		DependencyIndexes: file_occurrence_proto_depIdxs,
		MessageInfos:      file_occurrence_proto_msgTypes,
	}.Build()
	File_occurrence_proto = out.File
	file_occurrence_proto_rawDesc = nil
	file_occurrence_proto_goTypes = nil
	file_occurrence_proto_depIdxs = nil
}
