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
// source: instanceinfo/instanceinfo.proto

package instanceinfo

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

type CloudProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId        string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	InstanceId       string `protobuf:"bytes,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Zone             string `protobuf:"bytes,3,opt,name=zone,proto3" json:"zone,omitempty"`
	InstanceName     string `protobuf:"bytes,4,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	Image            string `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	NumericProjectId string `protobuf:"bytes,6,opt,name=numeric_project_id,json=numericProjectId,proto3" json:"numeric_project_id,omitempty"`
	Region           string `protobuf:"bytes,7,opt,name=region,proto3" json:"region,omitempty"` // This is needed only for baremtal systems and is not
}

func (x *CloudProperties) Reset() {
	*x = CloudProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_instanceinfo_instanceinfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudProperties) ProtoMessage() {}

func (x *CloudProperties) ProtoReflect() protoreflect.Message {
	mi := &file_instanceinfo_instanceinfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudProperties.ProtoReflect.Descriptor instead.
func (*CloudProperties) Descriptor() ([]byte, []int) {
	return file_instanceinfo_instanceinfo_proto_rawDescGZIP(), []int{0}
}

func (x *CloudProperties) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CloudProperties) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *CloudProperties) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *CloudProperties) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *CloudProperties) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CloudProperties) GetNumericProjectId() string {
	if x != nil {
		return x.NumericProjectId
	}
	return ""
}

func (x *CloudProperties) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type Disk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// e.g. SCRATCH, PERSISTENT, etc.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// last element of the disk source attribute, unique per project
	// only exists if this is a persistent disk
	DiskName string `protobuf:"bytes,2,opt,name=disk_name,json=diskName,proto3" json:"disk_name,omitempty"`
	// human readable device name, does not necessarily match the disk_name
	DeviceName string `protobuf:"bytes,3,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	// UNKNOWN, LOCAL_SSD, or PD_XXX
	DeviceType string `protobuf:"bytes,4,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	IsLocalSsd bool   `protobuf:"varint,5,opt,name=is_local_ssd,json=isLocalSsd,proto3" json:"is_local_ssd,omitempty"`
	// local disk mapping for device_name
	// found by following the link to  /dev/disk/by-id/google-*
	Mapping string `protobuf:"bytes,6,opt,name=mapping,proto3" json:"mapping,omitempty"`
}

func (x *Disk) Reset() {
	*x = Disk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_instanceinfo_instanceinfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Disk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Disk) ProtoMessage() {}

func (x *Disk) ProtoReflect() protoreflect.Message {
	mi := &file_instanceinfo_instanceinfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Disk.ProtoReflect.Descriptor instead.
func (*Disk) Descriptor() ([]byte, []int) {
	return file_instanceinfo_instanceinfo_proto_rawDescGZIP(), []int{1}
}

func (x *Disk) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Disk) GetDiskName() string {
	if x != nil {
		return x.DiskName
	}
	return ""
}

func (x *Disk) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *Disk) GetDeviceType() string {
	if x != nil {
		return x.DeviceType
	}
	return ""
}

func (x *Disk) GetIsLocalSsd() bool {
	if x != nil {
		return x.IsLocalSsd
	}
	return false
}

func (x *Disk) GetMapping() string {
	if x != nil {
		return x.Mapping
	}
	return ""
}

type NetworkAdapter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NetworkIp string `protobuf:"bytes,2,opt,name=network_ip,json=networkIp,proto3" json:"network_ip,omitempty"`
	Network   string `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
	// local nic name mapping
	Mapping string `protobuf:"bytes,4,opt,name=mapping,proto3" json:"mapping,omitempty"`
}

func (x *NetworkAdapter) Reset() {
	*x = NetworkAdapter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_instanceinfo_instanceinfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkAdapter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkAdapter) ProtoMessage() {}

func (x *NetworkAdapter) ProtoReflect() protoreflect.Message {
	mi := &file_instanceinfo_instanceinfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkAdapter.ProtoReflect.Descriptor instead.
func (*NetworkAdapter) Descriptor() ([]byte, []int) {
	return file_instanceinfo_instanceinfo_proto_rawDescGZIP(), []int{2}
}

func (x *NetworkAdapter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NetworkAdapter) GetNetworkIp() string {
	if x != nil {
		return x.NetworkIp
	}
	return ""
}

func (x *NetworkAdapter) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *NetworkAdapter) GetMapping() string {
	if x != nil {
		return x.Mapping
	}
	return ""
}

type InstanceProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachineType               string            `protobuf:"bytes,1,opt,name=machine_type,json=machineType,proto3" json:"machine_type,omitempty"`
	CpuPlatform               string            `protobuf:"bytes,2,opt,name=cpu_platform,json=cpuPlatform,proto3" json:"cpu_platform,omitempty"`
	Disks                     []*Disk           `protobuf:"bytes,3,rep,name=disks,proto3" json:"disks,omitempty"`
	NetworkAdapters           []*NetworkAdapter `protobuf:"bytes,4,rep,name=network_adapters,json=networkAdapters,proto3" json:"network_adapters,omitempty"`
	CreationTimestamp         string            `protobuf:"bytes,5,opt,name=creation_timestamp,json=creationTimestamp,proto3" json:"creation_timestamp,omitempty"`
	LastMigrationEndTimestamp string            `protobuf:"bytes,6,opt,name=last_migration_end_timestamp,json=lastMigrationEndTimestamp,proto3" json:"last_migration_end_timestamp,omitempty"`
}

func (x *InstanceProperties) Reset() {
	*x = InstanceProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_instanceinfo_instanceinfo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceProperties) ProtoMessage() {}

func (x *InstanceProperties) ProtoReflect() protoreflect.Message {
	mi := &file_instanceinfo_instanceinfo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceProperties.ProtoReflect.Descriptor instead.
func (*InstanceProperties) Descriptor() ([]byte, []int) {
	return file_instanceinfo_instanceinfo_proto_rawDescGZIP(), []int{3}
}

func (x *InstanceProperties) GetMachineType() string {
	if x != nil {
		return x.MachineType
	}
	return ""
}

func (x *InstanceProperties) GetCpuPlatform() string {
	if x != nil {
		return x.CpuPlatform
	}
	return ""
}

func (x *InstanceProperties) GetDisks() []*Disk {
	if x != nil {
		return x.Disks
	}
	return nil
}

func (x *InstanceProperties) GetNetworkAdapters() []*NetworkAdapter {
	if x != nil {
		return x.NetworkAdapters
	}
	return nil
}

func (x *InstanceProperties) GetCreationTimestamp() string {
	if x != nil {
		return x.CreationTimestamp
	}
	return ""
}

func (x *InstanceProperties) GetLastMigrationEndTimestamp() string {
	if x != nil {
		return x.LastMigrationEndTimestamp
	}
	return ""
}

var File_instanceinfo_instanceinfo_proto protoreflect.FileDescriptor

var file_instanceinfo_instanceinfo_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x73, 0x2e, 0x73, 0x61, 0x70, 0x2e, 0x67, 0x63, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x69, 0x6e, 0x66,
	0x6f, 0x22, 0xe6, 0x01, 0x0a, 0x0f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x5f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0xb5, 0x01, 0x0a, 0x04, 0x44,
	0x69, 0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x6b, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x5f, 0x73, 0x73, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x73, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x22, 0x77, 0x0a, 0x0e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x64, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x81, 0x03, 0x0a, 0x12,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x70, 0x75, 0x5f, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x70, 0x75,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x4a, 0x0a, 0x05, 0x64, 0x69, 0x73, 0x6b,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x61, 0x70, 0x2e, 0x67, 0x63, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x05, 0x64,
	0x69, 0x73, 0x6b, 0x73, 0x12, 0x69, 0x0a, 0x10, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f,
	0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x73, 0x2e,
	0x73, 0x61, 0x70, 0x2e, 0x67, 0x63, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x0f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x2d, 0x0a, 0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3f,
	0x0a, 0x1c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_instanceinfo_instanceinfo_proto_rawDescOnce sync.Once
	file_instanceinfo_instanceinfo_proto_rawDescData = file_instanceinfo_instanceinfo_proto_rawDesc
)

func file_instanceinfo_instanceinfo_proto_rawDescGZIP() []byte {
	file_instanceinfo_instanceinfo_proto_rawDescOnce.Do(func() {
		file_instanceinfo_instanceinfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_instanceinfo_instanceinfo_proto_rawDescData)
	})
	return file_instanceinfo_instanceinfo_proto_rawDescData
}

var file_instanceinfo_instanceinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_instanceinfo_instanceinfo_proto_goTypes = []interface{}{
	(*CloudProperties)(nil),    // 0: cloud.partners.sap.gcagent.protos.instanceinfo.CloudProperties
	(*Disk)(nil),               // 1: cloud.partners.sap.gcagent.protos.instanceinfo.Disk
	(*NetworkAdapter)(nil),     // 2: cloud.partners.sap.gcagent.protos.instanceinfo.NetworkAdapter
	(*InstanceProperties)(nil), // 3: cloud.partners.sap.gcagent.protos.instanceinfo.InstanceProperties
}
var file_instanceinfo_instanceinfo_proto_depIdxs = []int32{
	1, // 0: cloud.partners.sap.gcagent.protos.instanceinfo.InstanceProperties.disks:type_name -> cloud.partners.sap.gcagent.protos.instanceinfo.Disk
	2, // 1: cloud.partners.sap.gcagent.protos.instanceinfo.InstanceProperties.network_adapters:type_name -> cloud.partners.sap.gcagent.protos.instanceinfo.NetworkAdapter
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_instanceinfo_instanceinfo_proto_init() }
func file_instanceinfo_instanceinfo_proto_init() {
	if File_instanceinfo_instanceinfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_instanceinfo_instanceinfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudProperties); i {
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
		file_instanceinfo_instanceinfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Disk); i {
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
		file_instanceinfo_instanceinfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkAdapter); i {
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
		file_instanceinfo_instanceinfo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceProperties); i {
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
			RawDescriptor: file_instanceinfo_instanceinfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_instanceinfo_instanceinfo_proto_goTypes,
		DependencyIndexes: file_instanceinfo_instanceinfo_proto_depIdxs,
		MessageInfos:      file_instanceinfo_instanceinfo_proto_msgTypes,
	}.Build()
	File_instanceinfo_instanceinfo_proto = out.File
	file_instanceinfo_instanceinfo_proto_rawDesc = nil
	file_instanceinfo_instanceinfo_proto_goTypes = nil
	file_instanceinfo_instanceinfo_proto_depIdxs = nil
}
