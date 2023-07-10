//
//Copyright 2022 Google LLC
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//https://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: sapapp/sapapp.proto

package sapapp

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

type InstanceType int32

const (
	InstanceType_INSTANCE_TYPE_UNDEFINED InstanceType = 0
	InstanceType_HANA                    InstanceType = 1
	InstanceType_NETWEAVER               InstanceType = 2
)

// Enum value maps for InstanceType.
var (
	InstanceType_name = map[int32]string{
		0: "INSTANCE_TYPE_UNDEFINED",
		1: "HANA",
		2: "NETWEAVER",
	}
	InstanceType_value = map[string]int32{
		"INSTANCE_TYPE_UNDEFINED": 0,
		"HANA":                    1,
		"NETWEAVER":               2,
	}
)

func (x InstanceType) Enum() *InstanceType {
	p := new(InstanceType)
	*p = x
	return p
}

func (x InstanceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InstanceType) Descriptor() protoreflect.EnumDescriptor {
	return file_sapapp_sapapp_proto_enumTypes[0].Descriptor()
}

func (InstanceType) Type() protoreflect.EnumType {
	return &file_sapapp_sapapp_proto_enumTypes[0]
}

func (x InstanceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InstanceType.Descriptor instead.
func (InstanceType) EnumDescriptor() ([]byte, []int) {
	return file_sapapp_sapapp_proto_rawDescGZIP(), []int{0}
}

type InstanceSite int32

const (
	InstanceSite_INSTANCE_SITE_UNDEFINED InstanceSite = 0
	InstanceSite_HANA_PRIMARY            InstanceSite = 1
	InstanceSite_HANA_SECONDARY          InstanceSite = 2
	InstanceSite_HANA_STANDALONE         InstanceSite = 3
)

// Enum value maps for InstanceSite.
var (
	InstanceSite_name = map[int32]string{
		0: "INSTANCE_SITE_UNDEFINED",
		1: "HANA_PRIMARY",
		2: "HANA_SECONDARY",
		3: "HANA_STANDALONE",
	}
	InstanceSite_value = map[string]int32{
		"INSTANCE_SITE_UNDEFINED": 0,
		"HANA_PRIMARY":            1,
		"HANA_SECONDARY":          2,
		"HANA_STANDALONE":         3,
	}
)

func (x InstanceSite) Enum() *InstanceSite {
	p := new(InstanceSite)
	*p = x
	return p
}

func (x InstanceSite) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InstanceSite) Descriptor() protoreflect.EnumDescriptor {
	return file_sapapp_sapapp_proto_enumTypes[1].Descriptor()
}

func (InstanceSite) Type() protoreflect.EnumType {
	return &file_sapapp_sapapp_proto_enumTypes[1]
}

func (x InstanceSite) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InstanceSite.Descriptor instead.
func (InstanceSite) EnumDescriptor() ([]byte, []int) {
	return file_sapapp_sapapp_proto_rawDescGZIP(), []int{1}
}

type InstanceKind int32

const (
	InstanceKind_INSTANCE_KIND_UNDEFINED InstanceKind = 0
	InstanceKind_APP                     InstanceKind = 1
	InstanceKind_CS                      InstanceKind = 2
	InstanceKind_ERS                     InstanceKind = 3
)

// Enum value maps for InstanceKind.
var (
	InstanceKind_name = map[int32]string{
		0: "INSTANCE_KIND_UNDEFINED",
		1: "APP",
		2: "CS",
		3: "ERS",
	}
	InstanceKind_value = map[string]int32{
		"INSTANCE_KIND_UNDEFINED": 0,
		"APP":                     1,
		"CS":                      2,
		"ERS":                     3,
	}
)

func (x InstanceKind) Enum() *InstanceKind {
	p := new(InstanceKind)
	*p = x
	return p
}

func (x InstanceKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InstanceKind) Descriptor() protoreflect.EnumDescriptor {
	return file_sapapp_sapapp_proto_enumTypes[2].Descriptor()
}

func (InstanceKind) Type() protoreflect.EnumType {
	return &file_sapapp_sapapp_proto_enumTypes[2]
}

func (x InstanceKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InstanceKind.Descriptor instead.
func (InstanceKind) EnumDescriptor() ([]byte, []int) {
	return file_sapapp_sapapp_proto_rawDescGZIP(), []int{2}
}

type SAPInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sapsid                  string       `protobuf:"bytes,1,opt,name=sapsid,proto3" json:"sapsid,omitempty"`                                       // HDB
	InstanceNumber          string       `protobuf:"bytes,2,opt,name=instance_number,json=instanceNumber,proto3" json:"instance_number,omitempty"` // 00
	ServiceName             string       `protobuf:"bytes,3,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Type                    InstanceType `protobuf:"varint,4,opt,name=type,proto3,enum=cloud.partners.sap.gcagent.protos.sapapp.InstanceType" json:"type,omitempty"` // HANA, NetWeaver
	Site                    InstanceSite `protobuf:"varint,5,opt,name=site,proto3,enum=cloud.partners.sap.gcagent.protos.sapapp.InstanceSite" json:"site,omitempty"` // PRIMARY, SECONDARY
	HanaHaMembers           []string     `protobuf:"bytes,6,rep,name=hana_ha_members,json=hanaHaMembers,proto3" json:"hana_ha_members,omitempty"`                    // List of HANA instance names that form a HANA HA configuration.
	SapcontrolPath          string       `protobuf:"bytes,7,opt,name=sapcontrol_path,json=sapcontrolPath,proto3" json:"sapcontrol_path,omitempty"`                   // /usr/sap/HDB/HDB00/exe/sapcontrol
	User                    string       `protobuf:"bytes,8,opt,name=user,proto3" json:"user,omitempty"`                                                             // hdbadm
	InstanceId              string       `protobuf:"bytes,9,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`                               // HDB00 - unique identifier for SAP Instance.
	NetweaverHttpPort       string       `protobuf:"bytes,10,opt,name=netweaver_http_port,json=netweaverHttpPort,proto3" json:"netweaver_http_port,omitempty"`
	HanaDbUser              string       `protobuf:"bytes,11,opt,name=hana_db_user,json=hanaDbUser,proto3" json:"hana_db_user,omitempty"`
	HanaDbPassword          string       `protobuf:"bytes,12,opt,name=hana_db_password,json=hanaDbPassword,proto3" json:"hana_db_password,omitempty"`
	LdLibraryPath           string       `protobuf:"bytes,13,opt,name=ld_library_path,json=ldLibraryPath,proto3" json:"ld_library_path,omitempty"` // The Instance's LD_LIBRARY_PATH.
	ProfilePath             string       `protobuf:"bytes,14,opt,name=profile_path,json=profilePath,proto3" json:"profile_path,omitempty"`         // The instance's profile path.
	NetweaverHealthCheckUrl string       `protobuf:"bytes,15,opt,name=netweaver_health_check_url,json=netweaverHealthCheckUrl,proto3" json:"netweaver_health_check_url,omitempty"`
	Kind                    InstanceKind `protobuf:"varint,16,opt,name=kind,proto3,enum=cloud.partners.sap.gcagent.protos.sapapp.InstanceKind" json:"kind,omitempty"` // APP, CS, ERS
}

func (x *SAPInstance) Reset() {
	*x = SAPInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sapapp_sapapp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SAPInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SAPInstance) ProtoMessage() {}

func (x *SAPInstance) ProtoReflect() protoreflect.Message {
	mi := &file_sapapp_sapapp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SAPInstance.ProtoReflect.Descriptor instead.
func (*SAPInstance) Descriptor() ([]byte, []int) {
	return file_sapapp_sapapp_proto_rawDescGZIP(), []int{0}
}

func (x *SAPInstance) GetSapsid() string {
	if x != nil {
		return x.Sapsid
	}
	return ""
}

func (x *SAPInstance) GetInstanceNumber() string {
	if x != nil {
		return x.InstanceNumber
	}
	return ""
}

func (x *SAPInstance) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *SAPInstance) GetType() InstanceType {
	if x != nil {
		return x.Type
	}
	return InstanceType_INSTANCE_TYPE_UNDEFINED
}

func (x *SAPInstance) GetSite() InstanceSite {
	if x != nil {
		return x.Site
	}
	return InstanceSite_INSTANCE_SITE_UNDEFINED
}

func (x *SAPInstance) GetHanaHaMembers() []string {
	if x != nil {
		return x.HanaHaMembers
	}
	return nil
}

func (x *SAPInstance) GetSapcontrolPath() string {
	if x != nil {
		return x.SapcontrolPath
	}
	return ""
}

func (x *SAPInstance) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *SAPInstance) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *SAPInstance) GetNetweaverHttpPort() string {
	if x != nil {
		return x.NetweaverHttpPort
	}
	return ""
}

func (x *SAPInstance) GetHanaDbUser() string {
	if x != nil {
		return x.HanaDbUser
	}
	return ""
}

func (x *SAPInstance) GetHanaDbPassword() string {
	if x != nil {
		return x.HanaDbPassword
	}
	return ""
}

func (x *SAPInstance) GetLdLibraryPath() string {
	if x != nil {
		return x.LdLibraryPath
	}
	return ""
}

func (x *SAPInstance) GetProfilePath() string {
	if x != nil {
		return x.ProfilePath
	}
	return ""
}

func (x *SAPInstance) GetNetweaverHealthCheckUrl() string {
	if x != nil {
		return x.NetweaverHealthCheckUrl
	}
	return ""
}

func (x *SAPInstance) GetKind() InstanceKind {
	if x != nil {
		return x.Kind
	}
	return InstanceKind_INSTANCE_KIND_UNDEFINED
}

type SAPInstances struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instances          []*SAPInstance `protobuf:"bytes,1,rep,name=instances,proto3" json:"instances,omitempty"`
	LinuxClusterMember bool           `protobuf:"varint,2,opt,name=linux_cluster_member,json=linuxClusterMember,proto3" json:"linux_cluster_member,omitempty"`
}

func (x *SAPInstances) Reset() {
	*x = SAPInstances{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sapapp_sapapp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SAPInstances) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SAPInstances) ProtoMessage() {}

func (x *SAPInstances) ProtoReflect() protoreflect.Message {
	mi := &file_sapapp_sapapp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SAPInstances.ProtoReflect.Descriptor instead.
func (*SAPInstances) Descriptor() ([]byte, []int) {
	return file_sapapp_sapapp_proto_rawDescGZIP(), []int{1}
}

func (x *SAPInstances) GetInstances() []*SAPInstance {
	if x != nil {
		return x.Instances
	}
	return nil
}

func (x *SAPInstances) GetLinuxClusterMember() bool {
	if x != nil {
		return x.LinuxClusterMember
	}
	return false
}

var File_sapapp_sapapp_proto protoreflect.FileDescriptor

var file_sapapp_sapapp_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x61, 0x70, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x61, 0x70, 0x61, 0x70, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x61, 0x70, 0x2e, 0x67, 0x63, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x61, 0x70, 0x61, 0x70, 0x70, 0x22,
	0xdf, 0x05, 0x0a, 0x0b, 0x53, 0x41, 0x50, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x61, 0x70, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x61, 0x70, 0x73, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x36, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x73, 0x2e, 0x73, 0x61, 0x70, 0x2e, 0x67, 0x63, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x61, 0x70, 0x61, 0x70, 0x70, 0x2e, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x4a, 0x0a, 0x04, 0x73, 0x69, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x73,
	0x61, 0x70, 0x2e, 0x67, 0x63, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x73, 0x61, 0x70, 0x61, 0x70, 0x70, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x53, 0x69, 0x74, 0x65, 0x52, 0x04, 0x73, 0x69, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x68,
	0x61, 0x6e, 0x61, 0x5f, 0x68, 0x61, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x68, 0x61, 0x6e, 0x61, 0x48, 0x61, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x61, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x61,
	0x70, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x2e, 0x0a, 0x13, 0x6e, 0x65, 0x74, 0x77, 0x65, 0x61, 0x76, 0x65, 0x72, 0x5f, 0x68,
	0x74, 0x74, 0x70, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x6e, 0x65, 0x74, 0x77, 0x65, 0x61, 0x76, 0x65, 0x72, 0x48, 0x74, 0x74, 0x70, 0x50, 0x6f, 0x72,
	0x74, 0x12, 0x20, 0x0a, 0x0c, 0x68, 0x61, 0x6e, 0x61, 0x5f, 0x64, 0x62, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x61, 0x6e, 0x61, 0x44, 0x62, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x68, 0x61, 0x6e, 0x61, 0x5f, 0x64, 0x62, 0x5f, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x68,
	0x61, 0x6e, 0x61, 0x44, 0x62, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x26, 0x0a,
	0x0f, 0x6c, 0x64, 0x5f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x64, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x50, 0x61, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x3b, 0x0a, 0x1a, 0x6e, 0x65, 0x74, 0x77,
	0x65, 0x61, 0x76, 0x65, 0x72, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x6e, 0x65,
	0x74, 0x77, 0x65, 0x61, 0x76, 0x65, 0x72, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x4a, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x61, 0x70, 0x2e, 0x67, 0x63, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x61, 0x70, 0x61, 0x70, 0x70, 0x2e, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x22, 0x95, 0x01, 0x0a, 0x0c, 0x53, 0x41, 0x50, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x12, 0x53, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x61, 0x70, 0x2e, 0x67, 0x63, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x61, 0x70, 0x61, 0x70, 0x70,
	0x2e, 0x53, 0x41, 0x50, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x09, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x6c, 0x69, 0x6e, 0x75, 0x78,
	0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2a, 0x44, 0x0a, 0x0c, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x49, 0x4e, 0x53,
	0x54, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46,
	0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x41, 0x4e, 0x41, 0x10, 0x01,
	0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x45, 0x54, 0x57, 0x45, 0x41, 0x56, 0x45, 0x52, 0x10, 0x02, 0x2a,
	0x66, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x69, 0x74, 0x65, 0x12,
	0x1b, 0x0a, 0x17, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x53, 0x49, 0x54, 0x45,
	0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x48, 0x41, 0x4e, 0x41, 0x5f, 0x50, 0x52, 0x49, 0x4d, 0x41, 0x52, 0x59, 0x10, 0x01, 0x12, 0x12,
	0x0a, 0x0e, 0x48, 0x41, 0x4e, 0x41, 0x5f, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x41, 0x52, 0x59,
	0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x48, 0x41, 0x4e, 0x41, 0x5f, 0x53, 0x54, 0x41, 0x4e, 0x44,
	0x41, 0x4c, 0x4f, 0x4e, 0x45, 0x10, 0x03, 0x2a, 0x45, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x17, 0x49, 0x4e, 0x53, 0x54, 0x41,
	0x4e, 0x43, 0x45, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x50, 0x50, 0x10, 0x01, 0x12, 0x06, 0x0a,
	0x02, 0x43, 0x53, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x52, 0x53, 0x10, 0x03, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sapapp_sapapp_proto_rawDescOnce sync.Once
	file_sapapp_sapapp_proto_rawDescData = file_sapapp_sapapp_proto_rawDesc
)

func file_sapapp_sapapp_proto_rawDescGZIP() []byte {
	file_sapapp_sapapp_proto_rawDescOnce.Do(func() {
		file_sapapp_sapapp_proto_rawDescData = protoimpl.X.CompressGZIP(file_sapapp_sapapp_proto_rawDescData)
	})
	return file_sapapp_sapapp_proto_rawDescData
}

var file_sapapp_sapapp_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_sapapp_sapapp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sapapp_sapapp_proto_goTypes = []interface{}{
	(InstanceType)(0),    // 0: cloud.partners.sap.gcagent.protos.sapapp.InstanceType
	(InstanceSite)(0),    // 1: cloud.partners.sap.gcagent.protos.sapapp.InstanceSite
	(InstanceKind)(0),    // 2: cloud.partners.sap.gcagent.protos.sapapp.InstanceKind
	(*SAPInstance)(nil),  // 3: cloud.partners.sap.gcagent.protos.sapapp.SAPInstance
	(*SAPInstances)(nil), // 4: cloud.partners.sap.gcagent.protos.sapapp.SAPInstances
}
var file_sapapp_sapapp_proto_depIdxs = []int32{
	0, // 0: cloud.partners.sap.gcagent.protos.sapapp.SAPInstance.type:type_name -> cloud.partners.sap.gcagent.protos.sapapp.InstanceType
	1, // 1: cloud.partners.sap.gcagent.protos.sapapp.SAPInstance.site:type_name -> cloud.partners.sap.gcagent.protos.sapapp.InstanceSite
	2, // 2: cloud.partners.sap.gcagent.protos.sapapp.SAPInstance.kind:type_name -> cloud.partners.sap.gcagent.protos.sapapp.InstanceKind
	3, // 3: cloud.partners.sap.gcagent.protos.sapapp.SAPInstances.instances:type_name -> cloud.partners.sap.gcagent.protos.sapapp.SAPInstance
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_sapapp_sapapp_proto_init() }
func file_sapapp_sapapp_proto_init() {
	if File_sapapp_sapapp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sapapp_sapapp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SAPInstance); i {
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
		file_sapapp_sapapp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SAPInstances); i {
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
			RawDescriptor: file_sapapp_sapapp_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sapapp_sapapp_proto_goTypes,
		DependencyIndexes: file_sapapp_sapapp_proto_depIdxs,
		EnumInfos:         file_sapapp_sapapp_proto_enumTypes,
		MessageInfos:      file_sapapp_sapapp_proto_msgTypes,
	}.Build()
	File_sapapp_sapapp_proto = out.File
	file_sapapp_sapapp_proto_rawDesc = nil
	file_sapapp_sapapp_proto_goTypes = nil
	file_sapapp_sapapp_proto_depIdxs = nil
}
