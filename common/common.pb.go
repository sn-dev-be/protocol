// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: common/common.proto

package common

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

type BusinessMQEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event string `protobuf:"bytes,1,opt,name=event,proto3" json:"event"`
}

func (x *BusinessMQEvent) Reset() {
	*x = BusinessMQEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusinessMQEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusinessMQEvent) ProtoMessage() {}

func (x *BusinessMQEvent) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusinessMQEvent.ProtoReflect.Descriptor instead.
func (*BusinessMQEvent) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{0}
}

func (x *BusinessMQEvent) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

type CommonBusinessMQEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType      string          `protobuf:"bytes,1,opt,name=event_type,json=eventType,proto3" json:"event_type"`
	ClubServer     *ClubServer     `protobuf:"bytes,2,opt,name=club_server,json=clubServer,proto3" json:"club_server"`
	ClubServerUser *ClubServerUser `protobuf:"bytes,3,opt,name=club_server_user,json=clubServerUser,proto3" json:"club_server_user"`
	UserRelation   *UserRelation   `protobuf:"bytes,4,opt,name=user_relation,json=userRelation,proto3" json:"user_relation"`
}

func (x *CommonBusinessMQEvent) Reset() {
	*x = CommonBusinessMQEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonBusinessMQEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonBusinessMQEvent) ProtoMessage() {}

func (x *CommonBusinessMQEvent) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonBusinessMQEvent.ProtoReflect.Descriptor instead.
func (*CommonBusinessMQEvent) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{1}
}

func (x *CommonBusinessMQEvent) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *CommonBusinessMQEvent) GetClubServer() *ClubServer {
	if x != nil {
		return x.ClubServer
	}
	return nil
}

func (x *CommonBusinessMQEvent) GetClubServerUser() *ClubServerUser {
	if x != nil {
		return x.ClubServerUser
	}
	return nil
}

func (x *CommonBusinessMQEvent) GetUserRelation() *UserRelation {
	if x != nil {
		return x.UserRelation
	}
	return nil
}

type ClubServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClubServerId string `protobuf:"bytes,1,opt,name=club_server_id,json=clubServerId,proto3" json:"club_server_id"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Icon         string `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon"`
	Banner       string `protobuf:"bytes,4,opt,name=banner,proto3" json:"banner"`
	IsPublic     bool   `protobuf:"varint,5,opt,name=is_public,json=isPublic,proto3" json:"is_public"`
}

func (x *ClubServer) Reset() {
	*x = ClubServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClubServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClubServer) ProtoMessage() {}

func (x *ClubServer) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClubServer.ProtoReflect.Descriptor instead.
func (*ClubServer) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{2}
}

func (x *ClubServer) GetClubServerId() string {
	if x != nil {
		return x.ClubServerId
	}
	return ""
}

func (x *ClubServer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClubServer) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *ClubServer) GetBanner() string {
	if x != nil {
		return x.Banner
	}
	return ""
}

func (x *ClubServer) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

type ClubServerUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId string `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id"`
	UserId   string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id"`
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname"`
}

func (x *ClubServerUser) Reset() {
	*x = ClubServerUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClubServerUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClubServerUser) ProtoMessage() {}

func (x *ClubServerUser) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClubServerUser.ProtoReflect.Descriptor instead.
func (*ClubServerUser) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{3}
}

func (x *ClubServerUser) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *ClubServerUser) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ClubServerUser) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

type UserRelation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id"`
	TargetUserId string `protobuf:"bytes,2,opt,name=target_user_id,json=targetUserId,proto3" json:"target_user_id"`
	Remark       string `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark"`
}

func (x *UserRelation) Reset() {
	*x = UserRelation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRelation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRelation) ProtoMessage() {}

func (x *UserRelation) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRelation.ProtoReflect.Descriptor instead.
func (*UserRelation) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{4}
}

func (x *UserRelation) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserRelation) GetTargetUserId() string {
	if x != nil {
		return x.TargetUserId
	}
	return ""
}

func (x *UserRelation) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

var File_common_common_proto protoreflect.FileDescriptor

var file_common_common_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x4d, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x27, 0x0a, 0x0f, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4d, 0x51, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x22, 0x8f, 0x02, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4d, 0x51, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x40, 0x0a, 0x0b,
	0x63, 0x6c, 0x75, 0x62, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x75, 0x62, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x0a, 0x63, 0x6c, 0x75, 0x62, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x4d,
	0x0a, 0x10, 0x63, 0x6c, 0x75, 0x62, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x49,
	0x4d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43,
	0x6c, 0x75, 0x62, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x0e, 0x63,
	0x6c, 0x75, 0x62, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x46, 0x0a,
	0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x4d, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8f, 0x01, 0x0a, 0x0a, 0x43, 0x6c, 0x75, 0x62, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6c, 0x75, 0x62, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c,
	0x75, 0x62, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0x62, 0x0a, 0x0e, 0x43, 0x6c, 0x75, 0x62, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x65, 0x0a, 0x0c, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x4d, 0x53, 0x44, 0x4b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_common_common_proto_rawDescOnce sync.Once
	file_common_common_proto_rawDescData = file_common_common_proto_rawDesc
)

func file_common_common_proto_rawDescGZIP() []byte {
	file_common_common_proto_rawDescOnce.Do(func() {
		file_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_common_proto_rawDescData)
	})
	return file_common_common_proto_rawDescData
}

var file_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_common_common_proto_goTypes = []interface{}{
	(*BusinessMQEvent)(nil),       // 0: OpenIMServer.common.BusinessMQEvent
	(*CommonBusinessMQEvent)(nil), // 1: OpenIMServer.common.CommonBusinessMQEvent
	(*ClubServer)(nil),            // 2: OpenIMServer.common.ClubServer
	(*ClubServerUser)(nil),        // 3: OpenIMServer.common.ClubServerUser
	(*UserRelation)(nil),          // 4: OpenIMServer.common.UserRelation
}
var file_common_common_proto_depIdxs = []int32{
	2, // 0: OpenIMServer.common.CommonBusinessMQEvent.club_server:type_name -> OpenIMServer.common.ClubServer
	3, // 1: OpenIMServer.common.CommonBusinessMQEvent.club_server_user:type_name -> OpenIMServer.common.ClubServerUser
	4, // 2: OpenIMServer.common.CommonBusinessMQEvent.user_relation:type_name -> OpenIMServer.common.UserRelation
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_common_common_proto_init() }
func file_common_common_proto_init() {
	if File_common_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusinessMQEvent); i {
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
		file_common_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonBusinessMQEvent); i {
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
		file_common_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClubServer); i {
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
		file_common_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClubServerUser); i {
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
		file_common_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRelation); i {
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
			RawDescriptor: file_common_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_common_proto_goTypes,
		DependencyIndexes: file_common_common_proto_depIdxs,
		MessageInfos:      file_common_common_proto_msgTypes,
	}.Build()
	File_common_common_proto = out.File
	file_common_common_proto_rawDesc = nil
	file_common_common_proto_goTypes = nil
	file_common_common_proto_depIdxs = nil
}
