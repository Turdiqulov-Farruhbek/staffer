// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: jobsOffer.proto

package genprotos

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

// JobOffer modelini ta'riflash
type JobOffer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Company     string `protobuf:"bytes,4,opt,name=company,proto3" json:"company,omitempty"`
	Location    string `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	Status      string `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"` // e.g., "active", "closed"
}

func (x *JobOffer) Reset() {
	*x = JobOffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jobsOffer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobOffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobOffer) ProtoMessage() {}

func (x *JobOffer) ProtoReflect() protoreflect.Message {
	mi := &file_jobsOffer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobOffer.ProtoReflect.Descriptor instead.
func (*JobOffer) Descriptor() ([]byte, []int) {
	return file_jobsOffer_proto_rawDescGZIP(), []int{0}
}

func (x *JobOffer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *JobOffer) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *JobOffer) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *JobOffer) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *JobOffer) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *JobOffer) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// Yaratish uchun so'rov
type CreateJobOfferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobOffer *JobOffer `protobuf:"bytes,1,opt,name=job_offer,json=jobOffer,proto3" json:"job_offer,omitempty"`
}

func (x *CreateJobOfferRequest) Reset() {
	*x = CreateJobOfferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jobsOffer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobOfferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobOfferRequest) ProtoMessage() {}

func (x *CreateJobOfferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jobsOffer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobOfferRequest.ProtoReflect.Descriptor instead.
func (*CreateJobOfferRequest) Descriptor() ([]byte, []int) {
	return file_jobsOffer_proto_rawDescGZIP(), []int{1}
}

func (x *CreateJobOfferRequest) GetJobOffer() *JobOffer {
	if x != nil {
		return x.JobOffer
	}
	return nil
}

// ID bo'yicha so'rov
type GetByIdJobOfferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetByIdJobOfferRequest) Reset() {
	*x = GetByIdJobOfferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jobsOffer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIdJobOfferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdJobOfferRequest) ProtoMessage() {}

func (x *GetByIdJobOfferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jobsOffer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIdJobOfferRequest.ProtoReflect.Descriptor instead.
func (*GetByIdJobOfferRequest) Descriptor() ([]byte, []int) {
	return file_jobsOffer_proto_rawDescGZIP(), []int{2}
}

func (x *GetByIdJobOfferRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Barcha ish takliflarini olish uchun so'rov
type GetAllJobOffersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllJobOffersRequest) Reset() {
	*x = GetAllJobOffersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jobsOffer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllJobOffersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllJobOffersRequest) ProtoMessage() {}

func (x *GetAllJobOffersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jobsOffer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllJobOffersRequest.ProtoReflect.Descriptor instead.
func (*GetAllJobOffersRequest) Descriptor() ([]byte, []int) {
	return file_jobsOffer_proto_rawDescGZIP(), []int{3}
}

// Taklifni yangilash uchun so'rov
type UpdateJobOfferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	JobOffer *JobOffer `protobuf:"bytes,2,opt,name=job_offer,json=jobOffer,proto3" json:"job_offer,omitempty"`
}

func (x *UpdateJobOfferRequest) Reset() {
	*x = UpdateJobOfferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jobsOffer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJobOfferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobOfferRequest) ProtoMessage() {}

func (x *UpdateJobOfferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jobsOffer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobOfferRequest.ProtoReflect.Descriptor instead.
func (*UpdateJobOfferRequest) Descriptor() ([]byte, []int) {
	return file_jobsOffer_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateJobOfferRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateJobOfferRequest) GetJobOffer() *JobOffer {
	if x != nil {
		return x.JobOffer
	}
	return nil
}

// Taklifni o'chirish uchun so'rov
type DeleteJobOfferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteJobOfferRequest) Reset() {
	*x = DeleteJobOfferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jobsOffer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJobOfferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJobOfferRequest) ProtoMessage() {}

func (x *DeleteJobOfferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jobsOffer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJobOfferRequest.ProtoReflect.Descriptor instead.
func (*DeleteJobOfferRequest) Descriptor() ([]byte, []int) {
	return file_jobsOffer_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteJobOfferRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Javoblar
type JobOfferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobOffer *JobOffer `protobuf:"bytes,1,opt,name=job_offer,json=jobOffer,proto3" json:"job_offer,omitempty"`
}

func (x *JobOfferResponse) Reset() {
	*x = JobOfferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jobsOffer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobOfferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobOfferResponse) ProtoMessage() {}

func (x *JobOfferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_jobsOffer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobOfferResponse.ProtoReflect.Descriptor instead.
func (*JobOfferResponse) Descriptor() ([]byte, []int) {
	return file_jobsOffer_proto_rawDescGZIP(), []int{6}
}

func (x *JobOfferResponse) GetJobOffer() *JobOffer {
	if x != nil {
		return x.JobOffer
	}
	return nil
}

type JobOffersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobOffers []*JobOffer `protobuf:"bytes,1,rep,name=job_offers,json=jobOffers,proto3" json:"job_offers,omitempty"`
}

func (x *JobOffersResponse) Reset() {
	*x = JobOffersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jobsOffer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobOffersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobOffersResponse) ProtoMessage() {}

func (x *JobOffersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_jobsOffer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobOffersResponse.ProtoReflect.Descriptor instead.
func (*JobOffersResponse) Descriptor() ([]byte, []int) {
	return file_jobsOffer_proto_rawDescGZIP(), []int{7}
}

func (x *JobOffersResponse) GetJobOffers() []*JobOffer {
	if x != nil {
		return x.JobOffers
	}
	return nil
}

var File_jobsOffer_proto protoreflect.FileDescriptor

var file_jobsOffer_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6a, 0x6f, 0x62, 0x73, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x73, 0x74, 0x61, 0x66, 0x66, 0x22, 0xa0, 0x01, 0x0a, 0x08, 0x4a, 0x6f, 0x62,
	0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x45, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e,
	0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x08, 0x6a, 0x6f, 0x62, 0x4f, 0x66, 0x66,
	0x65, 0x72, 0x22, 0x28, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x4a, 0x6f, 0x62,
	0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x18, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x55, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x2c, 0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x4a, 0x6f, 0x62, 0x4f, 0x66,
	0x66, 0x65, 0x72, 0x52, 0x08, 0x6a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x22, 0x27, 0x0a,
	0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x10, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x6a, 0x6f,
	0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x08,
	0x6a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x22, 0x43, 0x0a, 0x11, 0x4a, 0x6f, 0x62, 0x4f,
	0x66, 0x66, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a,
	0x0a, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66,
	0x65, 0x72, 0x52, 0x09, 0x6a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x32, 0x83, 0x03,
	0x0a, 0x0f, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x47, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x66,
	0x66, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x1d, 0x2e,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x4a, 0x6f, 0x62,
	0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x2e, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4a,
	0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e,
	0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x47, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x66,
	0x66, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x66,
	0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x74, 0x61,
	0x66, 0x66, 0x2e, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jobsOffer_proto_rawDescOnce sync.Once
	file_jobsOffer_proto_rawDescData = file_jobsOffer_proto_rawDesc
)

func file_jobsOffer_proto_rawDescGZIP() []byte {
	file_jobsOffer_proto_rawDescOnce.Do(func() {
		file_jobsOffer_proto_rawDescData = protoimpl.X.CompressGZIP(file_jobsOffer_proto_rawDescData)
	})
	return file_jobsOffer_proto_rawDescData
}

var file_jobsOffer_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_jobsOffer_proto_goTypes = []any{
	(*JobOffer)(nil),               // 0: staff.JobOffer
	(*CreateJobOfferRequest)(nil),  // 1: staff.CreateJobOfferRequest
	(*GetByIdJobOfferRequest)(nil), // 2: staff.GetByIdJobOfferRequest
	(*GetAllJobOffersRequest)(nil), // 3: staff.GetAllJobOffersRequest
	(*UpdateJobOfferRequest)(nil),  // 4: staff.UpdateJobOfferRequest
	(*DeleteJobOfferRequest)(nil),  // 5: staff.DeleteJobOfferRequest
	(*JobOfferResponse)(nil),       // 6: staff.JobOfferResponse
	(*JobOffersResponse)(nil),      // 7: staff.JobOffersResponse
}
var file_jobsOffer_proto_depIdxs = []int32{
	0, // 0: staff.CreateJobOfferRequest.job_offer:type_name -> staff.JobOffer
	0, // 1: staff.UpdateJobOfferRequest.job_offer:type_name -> staff.JobOffer
	0, // 2: staff.JobOfferResponse.job_offer:type_name -> staff.JobOffer
	0, // 3: staff.JobOffersResponse.job_offers:type_name -> staff.JobOffer
	1, // 4: staff.JobOfferService.CreateJobOffer:input_type -> staff.CreateJobOfferRequest
	2, // 5: staff.JobOfferService.GetByIdJobOffer:input_type -> staff.GetByIdJobOfferRequest
	3, // 6: staff.JobOfferService.GetAllJobOffers:input_type -> staff.GetAllJobOffersRequest
	4, // 7: staff.JobOfferService.UpdateJobOffer:input_type -> staff.UpdateJobOfferRequest
	5, // 8: staff.JobOfferService.DeleteJobOffer:input_type -> staff.DeleteJobOfferRequest
	6, // 9: staff.JobOfferService.CreateJobOffer:output_type -> staff.JobOfferResponse
	6, // 10: staff.JobOfferService.GetByIdJobOffer:output_type -> staff.JobOfferResponse
	7, // 11: staff.JobOfferService.GetAllJobOffers:output_type -> staff.JobOffersResponse
	6, // 12: staff.JobOfferService.UpdateJobOffer:output_type -> staff.JobOfferResponse
	6, // 13: staff.JobOfferService.DeleteJobOffer:output_type -> staff.JobOfferResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_jobsOffer_proto_init() }
func file_jobsOffer_proto_init() {
	if File_jobsOffer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_jobsOffer_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*JobOffer); i {
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
		file_jobsOffer_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateJobOfferRequest); i {
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
		file_jobsOffer_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetByIdJobOfferRequest); i {
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
		file_jobsOffer_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllJobOffersRequest); i {
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
		file_jobsOffer_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateJobOfferRequest); i {
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
		file_jobsOffer_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteJobOfferRequest); i {
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
		file_jobsOffer_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*JobOfferResponse); i {
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
		file_jobsOffer_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*JobOffersResponse); i {
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
			RawDescriptor: file_jobsOffer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_jobsOffer_proto_goTypes,
		DependencyIndexes: file_jobsOffer_proto_depIdxs,
		MessageInfos:      file_jobsOffer_proto_msgTypes,
	}.Build()
	File_jobsOffer_proto = out.File
	file_jobsOffer_proto_rawDesc = nil
	file_jobsOffer_proto_goTypes = nil
	file_jobsOffer_proto_depIdxs = nil
}
