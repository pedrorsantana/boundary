// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: controller/api/services/v1/host_service.proto

package services

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	hosts "github.com/hashicorp/boundary/internal/gen/controller/api/resources/hosts"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetHostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetHostRequest) Reset() {
	*x = GetHostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_services_v1_host_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHostRequest) ProtoMessage() {}

func (x *GetHostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_host_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHostRequest.ProtoReflect.Descriptor instead.
func (*GetHostRequest) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_host_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetHostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetHostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *hosts.Host `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *GetHostResponse) Reset() {
	*x = GetHostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_services_v1_host_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHostResponse) ProtoMessage() {}

func (x *GetHostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_host_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHostResponse.ProtoReflect.Descriptor instead.
func (*GetHostResponse) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_host_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetHostResponse) GetItem() *hosts.Host {
	if x != nil {
		return x.Item
	}
	return nil
}

type ListHostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostCatalogId string `protobuf:"bytes,1,opt,name=host_catalog_id,proto3" json:"host_catalog_id,omitempty"`
	Filter        string `protobuf:"bytes,30,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListHostsRequest) Reset() {
	*x = ListHostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_services_v1_host_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHostsRequest) ProtoMessage() {}

func (x *ListHostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_host_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHostsRequest.ProtoReflect.Descriptor instead.
func (*ListHostsRequest) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_host_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListHostsRequest) GetHostCatalogId() string {
	if x != nil {
		return x.HostCatalogId
	}
	return ""
}

func (x *ListHostsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type ListHostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*hosts.Host `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListHostsResponse) Reset() {
	*x = ListHostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_services_v1_host_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHostsResponse) ProtoMessage() {}

func (x *ListHostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_host_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHostsResponse.ProtoReflect.Descriptor instead.
func (*ListHostsResponse) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_host_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListHostsResponse) GetItems() []*hosts.Host {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreateHostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *hosts.Host `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *CreateHostRequest) Reset() {
	*x = CreateHostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_services_v1_host_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHostRequest) ProtoMessage() {}

func (x *CreateHostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_host_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHostRequest.ProtoReflect.Descriptor instead.
func (*CreateHostRequest) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_host_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateHostRequest) GetItem() *hosts.Host {
	if x != nil {
		return x.Item
	}
	return nil
}

type CreateHostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri  string      `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	Item *hosts.Host `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *CreateHostResponse) Reset() {
	*x = CreateHostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_services_v1_host_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHostResponse) ProtoMessage() {}

func (x *CreateHostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_host_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHostResponse.ProtoReflect.Descriptor instead.
func (*CreateHostResponse) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_host_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateHostResponse) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *CreateHostResponse) GetItem() *hosts.Host {
	if x != nil {
		return x.Item
	}
	return nil
}

type UpdateHostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Item       *hosts.Host           `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,3,opt,name=update_mask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateHostRequest) Reset() {
	*x = UpdateHostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_services_v1_host_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHostRequest) ProtoMessage() {}

func (x *UpdateHostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_host_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHostRequest.ProtoReflect.Descriptor instead.
func (*UpdateHostRequest) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_host_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateHostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateHostRequest) GetItem() *hosts.Host {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *UpdateHostRequest) GetUpdateMask() *field_mask.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type UpdateHostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *hosts.Host `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *UpdateHostResponse) Reset() {
	*x = UpdateHostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_services_v1_host_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHostResponse) ProtoMessage() {}

func (x *UpdateHostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_host_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHostResponse.ProtoReflect.Descriptor instead.
func (*UpdateHostResponse) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_host_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateHostResponse) GetItem() *hosts.Host {
	if x != nil {
		return x.Item
	}
	return nil
}

type DeleteHostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteHostRequest) Reset() {
	*x = DeleteHostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_services_v1_host_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHostRequest) ProtoMessage() {}

func (x *DeleteHostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_host_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHostRequest.ProtoReflect.Descriptor instead.
func (*DeleteHostRequest) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_host_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteHostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteHostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteHostResponse) Reset() {
	*x = DeleteHostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_services_v1_host_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHostResponse) ProtoMessage() {}

func (x *DeleteHostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_host_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHostResponse.ProtoReflect.Descriptor instead.
func (*DeleteHostResponse) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_host_service_proto_rawDescGZIP(), []int{9}
}

var File_controller_api_services_v1_host_service_proto protoreflect.FileDescriptor

var file_controller_api_services_v1_host_service_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x73,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x68,
	0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4e, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x54, 0x0a, 0x10, 0x4c,
	0x69, 0x73, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x28, 0x0a, 0x0f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x22, 0x52, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x50, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x04, 0x69, 0x74,
	0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73,
	0x74, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x63, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12,
	0x3b, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x9e, 0x01, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x3b, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x68, 0x6f, 0x73, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12,
	0x3c, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x22, 0x51, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x68, 0x6f, 0x73,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d,
	0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb9, 0x06, 0x0a, 0x0b,
	0x48, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x98, 0x01, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f,
	0x73, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x62, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x92, 0x41,
	0x15, 0x12, 0x13, 0x47, 0x65, 0x74, 0x73, 0x20, 0x61, 0x20, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x20, 0x48, 0x6f, 0x73, 0x74, 0x2e, 0x12, 0xa9, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x48,
	0x6f, 0x73, 0x74, 0x73, 0x12, 0x2c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x3f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x68,
	0x6f, 0x73, 0x74, 0x73, 0x92, 0x41, 0x2b, 0x12, 0x29, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x61, 0x6c,
	0x6c, 0x20, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x20, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2e, 0x12, 0xa4, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73,
	0x74, 0x12, 0x2d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x37, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f,
	0x73, 0x74, 0x73, 0x3a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x62, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x92,
	0x41, 0x17, 0x12, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x73, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x20, 0x48, 0x6f, 0x73, 0x74, 0x2e, 0x12, 0xa2, 0x01, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x2d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x32,
	0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a,
	0x04, 0x69, 0x74, 0x65, 0x6d, 0x62, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x92, 0x41, 0x10, 0x12, 0x0e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x48, 0x6f, 0x73, 0x74, 0x2e, 0x12, 0x96,
	0x01, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x2d, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x10, 0x2a, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x92, 0x41, 0x10, 0x12, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20,
	0x61, 0x20, 0x48, 0x6f, 0x73, 0x74, 0x2e, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_api_services_v1_host_service_proto_rawDescOnce sync.Once
	file_controller_api_services_v1_host_service_proto_rawDescData = file_controller_api_services_v1_host_service_proto_rawDesc
)

func file_controller_api_services_v1_host_service_proto_rawDescGZIP() []byte {
	file_controller_api_services_v1_host_service_proto_rawDescOnce.Do(func() {
		file_controller_api_services_v1_host_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_api_services_v1_host_service_proto_rawDescData)
	})
	return file_controller_api_services_v1_host_service_proto_rawDescData
}

var file_controller_api_services_v1_host_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_controller_api_services_v1_host_service_proto_goTypes = []interface{}{
	(*GetHostRequest)(nil),       // 0: controller.api.services.v1.GetHostRequest
	(*GetHostResponse)(nil),      // 1: controller.api.services.v1.GetHostResponse
	(*ListHostsRequest)(nil),     // 2: controller.api.services.v1.ListHostsRequest
	(*ListHostsResponse)(nil),    // 3: controller.api.services.v1.ListHostsResponse
	(*CreateHostRequest)(nil),    // 4: controller.api.services.v1.CreateHostRequest
	(*CreateHostResponse)(nil),   // 5: controller.api.services.v1.CreateHostResponse
	(*UpdateHostRequest)(nil),    // 6: controller.api.services.v1.UpdateHostRequest
	(*UpdateHostResponse)(nil),   // 7: controller.api.services.v1.UpdateHostResponse
	(*DeleteHostRequest)(nil),    // 8: controller.api.services.v1.DeleteHostRequest
	(*DeleteHostResponse)(nil),   // 9: controller.api.services.v1.DeleteHostResponse
	(*hosts.Host)(nil),           // 10: controller.api.resources.hosts.v1.Host
	(*field_mask.FieldMask)(nil), // 11: google.protobuf.FieldMask
}
var file_controller_api_services_v1_host_service_proto_depIdxs = []int32{
	10, // 0: controller.api.services.v1.GetHostResponse.item:type_name -> controller.api.resources.hosts.v1.Host
	10, // 1: controller.api.services.v1.ListHostsResponse.items:type_name -> controller.api.resources.hosts.v1.Host
	10, // 2: controller.api.services.v1.CreateHostRequest.item:type_name -> controller.api.resources.hosts.v1.Host
	10, // 3: controller.api.services.v1.CreateHostResponse.item:type_name -> controller.api.resources.hosts.v1.Host
	10, // 4: controller.api.services.v1.UpdateHostRequest.item:type_name -> controller.api.resources.hosts.v1.Host
	11, // 5: controller.api.services.v1.UpdateHostRequest.update_mask:type_name -> google.protobuf.FieldMask
	10, // 6: controller.api.services.v1.UpdateHostResponse.item:type_name -> controller.api.resources.hosts.v1.Host
	0,  // 7: controller.api.services.v1.HostService.GetHost:input_type -> controller.api.services.v1.GetHostRequest
	2,  // 8: controller.api.services.v1.HostService.ListHosts:input_type -> controller.api.services.v1.ListHostsRequest
	4,  // 9: controller.api.services.v1.HostService.CreateHost:input_type -> controller.api.services.v1.CreateHostRequest
	6,  // 10: controller.api.services.v1.HostService.UpdateHost:input_type -> controller.api.services.v1.UpdateHostRequest
	8,  // 11: controller.api.services.v1.HostService.DeleteHost:input_type -> controller.api.services.v1.DeleteHostRequest
	1,  // 12: controller.api.services.v1.HostService.GetHost:output_type -> controller.api.services.v1.GetHostResponse
	3,  // 13: controller.api.services.v1.HostService.ListHosts:output_type -> controller.api.services.v1.ListHostsResponse
	5,  // 14: controller.api.services.v1.HostService.CreateHost:output_type -> controller.api.services.v1.CreateHostResponse
	7,  // 15: controller.api.services.v1.HostService.UpdateHost:output_type -> controller.api.services.v1.UpdateHostResponse
	9,  // 16: controller.api.services.v1.HostService.DeleteHost:output_type -> controller.api.services.v1.DeleteHostResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_controller_api_services_v1_host_service_proto_init() }
func file_controller_api_services_v1_host_service_proto_init() {
	if File_controller_api_services_v1_host_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_api_services_v1_host_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHostRequest); i {
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
		file_controller_api_services_v1_host_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHostResponse); i {
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
		file_controller_api_services_v1_host_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHostsRequest); i {
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
		file_controller_api_services_v1_host_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHostsResponse); i {
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
		file_controller_api_services_v1_host_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHostRequest); i {
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
		file_controller_api_services_v1_host_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHostResponse); i {
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
		file_controller_api_services_v1_host_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHostRequest); i {
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
		file_controller_api_services_v1_host_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHostResponse); i {
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
		file_controller_api_services_v1_host_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHostRequest); i {
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
		file_controller_api_services_v1_host_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHostResponse); i {
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
			RawDescriptor: file_controller_api_services_v1_host_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_controller_api_services_v1_host_service_proto_goTypes,
		DependencyIndexes: file_controller_api_services_v1_host_service_proto_depIdxs,
		MessageInfos:      file_controller_api_services_v1_host_service_proto_msgTypes,
	}.Build()
	File_controller_api_services_v1_host_service_proto = out.File
	file_controller_api_services_v1_host_service_proto_rawDesc = nil
	file_controller_api_services_v1_host_service_proto_goTypes = nil
	file_controller_api_services_v1_host_service_proto_depIdxs = nil
}
