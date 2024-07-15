// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: company.proto

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

type CreateCompany struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address  string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Industry string `protobuf:"bytes,3,opt,name=industry,proto3" json:"industry,omitempty"`
	Website  string `protobuf:"bytes,4,opt,name=website,proto3" json:"website,omitempty"`
}

func (x *CreateCompany) Reset() {
	*x = CreateCompany{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCompany) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCompany) ProtoMessage() {}

func (x *CreateCompany) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCompany.ProtoReflect.Descriptor instead.
func (*CreateCompany) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCompany) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCompany) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateCompany) GetIndustry() string {
	if x != nil {
		return x.Industry
	}
	return ""
}

func (x *CreateCompany) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

type CompanyUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address  string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Industry string `protobuf:"bytes,4,opt,name=industry,proto3" json:"industry,omitempty"`
	Website  string `protobuf:"bytes,5,opt,name=website,proto3" json:"website,omitempty"`
}

func (x *CompanyUpdate) Reset() {
	*x = CompanyUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyUpdate) ProtoMessage() {}

func (x *CompanyUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyUpdate.ProtoReflect.Descriptor instead.
func (*CompanyUpdate) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{1}
}

func (x *CompanyUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CompanyUpdate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CompanyUpdate) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CompanyUpdate) GetIndustry() string {
	if x != nil {
		return x.Industry
	}
	return ""
}

func (x *CompanyUpdate) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

type GetAllCompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address   string      `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Industry  string      `protobuf:"bytes,3,opt,name=industry,proto3" json:"industry,omitempty"`
	Website   string      `protobuf:"bytes,4,opt,name=website,proto3" json:"website,omitempty"`
	CompanyId string      `protobuf:"bytes,5,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Filter    *Pagination `protobuf:"bytes,6,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *GetAllCompanyRequest) Reset() {
	*x = GetAllCompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCompanyRequest) ProtoMessage() {}

func (x *GetAllCompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCompanyRequest.ProtoReflect.Descriptor instead.
func (*GetAllCompanyRequest) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllCompanyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetAllCompanyRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetAllCompanyRequest) GetIndustry() string {
	if x != nil {
		return x.Industry
	}
	return ""
}

func (x *GetAllCompanyRequest) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *GetAllCompanyRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *GetAllCompanyRequest) GetFilter() *Pagination {
	if x != nil {
		return x.Filter
	}
	return nil
}

type GetAllCompanyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Companies []*GetAllCompanyRequest `protobuf:"bytes,1,rep,name=Companies,proto3" json:"Companies,omitempty"`
	Success   bool                    `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *GetAllCompanyResponse) Reset() {
	*x = GetAllCompanyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCompanyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCompanyResponse) ProtoMessage() {}

func (x *GetAllCompanyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCompanyResponse.ProtoReflect.Descriptor instead.
func (*GetAllCompanyResponse) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllCompanyResponse) GetCompanies() []*GetAllCompanyRequest {
	if x != nil {
		return x.Companies
	}
	return nil
}

func (x *GetAllCompanyResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetByID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address   string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Industry  string `protobuf:"bytes,3,opt,name=industry,proto3" json:"industry,omitempty"`
	Website   string `protobuf:"bytes,4,opt,name=website,proto3" json:"website,omitempty"`
	CompanyId string `protobuf:"bytes,5,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Success   bool   `protobuf:"varint,6,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *GetByID) Reset() {
	*x = GetByID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByID) ProtoMessage() {}

func (x *GetByID) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByID.ProtoReflect.Descriptor instead.
func (*GetByID) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{4}
}

func (x *GetByID) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetByID) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetByID) GetIndustry() string {
	if x != nil {
		return x.Industry
	}
	return ""
}

func (x *GetByID) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *GetByID) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *GetByID) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_company_proto protoreflect.FileDescriptor

var file_company_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x73, 0x74, 0x61, 0x66, 0x66, 0x1a, 0x0a, 0x76, 0x6f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x73, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73,
	0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73,
	0x74, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x22, 0xc4, 0x01,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x22, 0x6c, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a,
	0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x09, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0xa6, 0x01, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73,
	0x69, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xfb, 0x01, 0x0a, 0x0e,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x1a, 0x0b,
	0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x43, 0x0a, 0x06, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1b, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x74, 0x61,
	0x66, 0x66, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x1a, 0x0b, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x22, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e,
	0x42, 0x79, 0x69, 0x64, 0x1a, 0x0b, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x56, 0x6f, 0x69,
	0x64, 0x12, 0x26, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0b, 0x2e, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x2e, 0x42, 0x79, 0x69, 0x64, 0x1a, 0x0e, 0x2e, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_company_proto_rawDescOnce sync.Once
	file_company_proto_rawDescData = file_company_proto_rawDesc
)

func file_company_proto_rawDescGZIP() []byte {
	file_company_proto_rawDescOnce.Do(func() {
		file_company_proto_rawDescData = protoimpl.X.CompressGZIP(file_company_proto_rawDescData)
	})
	return file_company_proto_rawDescData
}

var file_company_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_company_proto_goTypes = []any{
	(*CreateCompany)(nil),         // 0: staff.CreateCompany
	(*CompanyUpdate)(nil),         // 1: staff.CompanyUpdate
	(*GetAllCompanyRequest)(nil),  // 2: staff.GetAllCompanyRequest
	(*GetAllCompanyResponse)(nil), // 3: staff.GetAllCompanyResponse
	(*GetByID)(nil),               // 4: staff.GetByID
	(*Pagination)(nil),            // 5: staff.Pagination
	(*Byid)(nil),                  // 6: staff.Byid
	(*Void)(nil),                  // 7: staff.Void
}
var file_company_proto_depIdxs = []int32{
	5, // 0: staff.GetAllCompanyRequest.filter:type_name -> staff.Pagination
	2, // 1: staff.GetAllCompanyResponse.Companies:type_name -> staff.GetAllCompanyRequest
	0, // 2: staff.CompanyService.Create:input_type -> staff.CreateCompany
	2, // 3: staff.CompanyService.GetAll:input_type -> staff.GetAllCompanyRequest
	1, // 4: staff.CompanyService.Update:input_type -> staff.CompanyUpdate
	6, // 5: staff.CompanyService.Delete:input_type -> staff.Byid
	6, // 6: staff.CompanyService.GetById:input_type -> staff.Byid
	7, // 7: staff.CompanyService.Create:output_type -> staff.Void
	3, // 8: staff.CompanyService.GetAll:output_type -> staff.GetAllCompanyResponse
	7, // 9: staff.CompanyService.Update:output_type -> staff.Void
	7, // 10: staff.CompanyService.Delete:output_type -> staff.Void
	4, // 11: staff.CompanyService.GetById:output_type -> staff.GetByID
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_company_proto_init() }
func file_company_proto_init() {
	if File_company_proto != nil {
		return
	}
	file_void_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_company_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateCompany); i {
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
		file_company_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CompanyUpdate); i {
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
		file_company_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllCompanyRequest); i {
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
		file_company_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllCompanyResponse); i {
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
		file_company_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetByID); i {
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
			RawDescriptor: file_company_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_company_proto_goTypes,
		DependencyIndexes: file_company_proto_depIdxs,
		MessageInfos:      file_company_proto_msgTypes,
	}.Build()
	File_company_proto = out.File
	file_company_proto_rawDesc = nil
	file_company_proto_goTypes = nil
	file_company_proto_depIdxs = nil
}
