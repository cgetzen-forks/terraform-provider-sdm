// Copyright 2020 StrongDM Inc
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
//
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.3
// source: accounts_access.proto

package v1

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

// AccountAccessListRequest specifies criteria for retrieving a list of AccountAccess records.
type AccountAccessListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Paging parameters for the query.
	Meta *ListRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A human-readable filter query string.
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *AccountAccessListRequest) Reset() {
	*x = AccountAccessListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_access_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountAccessListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountAccessListRequest) ProtoMessage() {}

func (x *AccountAccessListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_access_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountAccessListRequest.ProtoReflect.Descriptor instead.
func (*AccountAccessListRequest) Descriptor() ([]byte, []int) {
	return file_accounts_access_proto_rawDescGZIP(), []int{0}
}

func (x *AccountAccessListRequest) GetMeta() *ListRequestMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *AccountAccessListRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// AccountAccessListResponse returns a list of AccountAccess records that meet
// the criteria of an AccountAccessListRequest.
type AccountAccessListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Paging information for the query.
	Meta *ListResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A single page of results matching the list request criteria.
	Accesses []*AccountAccess `protobuf:"bytes,2,rep,name=accesses,proto3" json:"accesses,omitempty"`
	// Rate limit information.
	RateLimit *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
}

func (x *AccountAccessListResponse) Reset() {
	*x = AccountAccessListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_access_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountAccessListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountAccessListResponse) ProtoMessage() {}

func (x *AccountAccessListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_access_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountAccessListResponse.ProtoReflect.Descriptor instead.
func (*AccountAccessListResponse) Descriptor() ([]byte, []int) {
	return file_accounts_access_proto_rawDescGZIP(), []int{1}
}

func (x *AccountAccessListResponse) GetMeta() *ListResponseMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *AccountAccessListResponse) GetAccesses() []*AccountAccess {
	if x != nil {
		return x.Accesses
	}
	return nil
}

func (x *AccountAccessListResponse) GetRateLimit() *RateLimitMetadata {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

// AccountAccess represents an individual access grant of a Account to a Resource.
type AccountAccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of the Account to which access is granted.
	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// The unique identifier of the Resource to which access is granted.
	ResourceId string `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// The most recent time at which access was granted. If access was granted,
	// revoked, and granted again, this will reflect the later time.
	GrantedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=granted_at,json=grantedAt,proto3" json:"granted_at,omitempty"`
	// The time at which access will expire. If empty, this access has no expiration.
	ExpiresAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	// The unique identifier of the Role through which the Account was granted access to the Resource.
	// If empty, access was not granted through an AccountAttachment to a Role.
	RoleId string `protobuf:"bytes,5,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	// The unique identifier of the AccountGrant through which the Account was granted access to the Resource.
	// If empty, access was not granted through an AccountGrant.
	AccountGrantId string `protobuf:"bytes,6,opt,name=account_grant_id,json=accountGrantId,proto3" json:"account_grant_id,omitempty"`
}

func (x *AccountAccess) Reset() {
	*x = AccountAccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_access_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountAccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountAccess) ProtoMessage() {}

func (x *AccountAccess) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_access_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountAccess.ProtoReflect.Descriptor instead.
func (*AccountAccess) Descriptor() ([]byte, []int) {
	return file_accounts_access_proto_rawDescGZIP(), []int{2}
}

func (x *AccountAccess) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *AccountAccess) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *AccountAccess) GetGrantedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.GrantedAt
	}
	return nil
}

func (x *AccountAccess) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *AccountAccess) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *AccountAccess) GetAccountGrantId() string {
	if x != nil {
		return x.AccountGrantId
	}
	return ""
}

var File_accounts_access_proto protoreflect.FileDescriptor

var file_accounts_access_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x73, 0x70, 0x65,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x18, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x12, 0x22, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x3a, 0x14, 0xfa, 0xf8, 0xb3, 0x07, 0x0f, 0xd2, 0xf3, 0xb3, 0x07,
	0x0a, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22, 0xfe, 0x01, 0x0a, 0x19,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x0a, 0xf2, 0xf8,
	0xb3, 0x07, 0x05, 0xb8, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x08, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x12, 0x62, 0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x2c, 0xf2,
	0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xf2, 0xf8, 0xb3, 0x07, 0x06, 0xb2, 0xf4,
	0xb3, 0x07, 0x01, 0x2a, 0xf2, 0xf8, 0xb3, 0x07, 0x12, 0xb2, 0xf4, 0xb3, 0x07, 0x0d, 0x21, 0x6a,
	0x73, 0x6f, 0x6e, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x09, 0x72, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x3a, 0x14, 0xfa, 0xf8, 0xb3, 0x07, 0x0f, 0xd2, 0xf3, 0xb3,
	0x07, 0x0a, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22, 0xf0, 0x02, 0x0a,
	0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x29,
	0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x0b, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3,
	0x07, 0x01, 0x52, 0x09, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x45, 0x0a,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a, 0xf2,
	0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x41, 0x74, 0x12, 0x23, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07,
	0x01, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x10, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52,
	0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x3a,
	0x1e, 0xfa, 0xf8, 0xb3, 0x07, 0x05, 0xa8, 0xf3, 0xb3, 0x07, 0x01, 0xfa, 0xf8, 0xb3, 0x07, 0x0f,
	0xd2, 0xf3, 0xb3, 0x07, 0x0a, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x32,
	0xae, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x6f, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xf9, 0xb3, 0x07, 0x08, 0xa2, 0xf3,
	0xb3, 0x07, 0x03, 0x67, 0x65, 0x74, 0x82, 0xf9, 0xb3, 0x07, 0x18, 0xaa, 0xf3, 0xb3, 0x07, 0x13,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2d, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x1a, 0x2b, 0xca, 0xf9, 0xb3, 0x07, 0x12, 0xc2, 0xf9, 0xb3, 0x07, 0x0d, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0xca, 0xf9, 0xb3, 0x07,
	0x0f, 0xca, 0xf9, 0xb3, 0x07, 0x0a, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x42, 0x7e, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6e, 0x67, 0x42, 0x16, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6c, 0x75,
	0x6d, 0x62, 0x69, 0x6e, 0x67, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e,
	0x67, 0x64, 0x6d, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0xc2, 0x92, 0xb4, 0x07,
	0x0f, 0xa2, 0x8c, 0xb4, 0x07, 0x0a, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accounts_access_proto_rawDescOnce sync.Once
	file_accounts_access_proto_rawDescData = file_accounts_access_proto_rawDesc
)

func file_accounts_access_proto_rawDescGZIP() []byte {
	file_accounts_access_proto_rawDescOnce.Do(func() {
		file_accounts_access_proto_rawDescData = protoimpl.X.CompressGZIP(file_accounts_access_proto_rawDescData)
	})
	return file_accounts_access_proto_rawDescData
}

var file_accounts_access_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_accounts_access_proto_goTypes = []interface{}{
	(*AccountAccessListRequest)(nil),  // 0: v1.AccountAccessListRequest
	(*AccountAccessListResponse)(nil), // 1: v1.AccountAccessListResponse
	(*AccountAccess)(nil),             // 2: v1.AccountAccess
	(*ListRequestMetadata)(nil),       // 3: v1.ListRequestMetadata
	(*ListResponseMetadata)(nil),      // 4: v1.ListResponseMetadata
	(*RateLimitMetadata)(nil),         // 5: v1.RateLimitMetadata
	(*timestamppb.Timestamp)(nil),     // 6: google.protobuf.Timestamp
}
var file_accounts_access_proto_depIdxs = []int32{
	3, // 0: v1.AccountAccessListRequest.meta:type_name -> v1.ListRequestMetadata
	4, // 1: v1.AccountAccessListResponse.meta:type_name -> v1.ListResponseMetadata
	2, // 2: v1.AccountAccessListResponse.accesses:type_name -> v1.AccountAccess
	5, // 3: v1.AccountAccessListResponse.rate_limit:type_name -> v1.RateLimitMetadata
	6, // 4: v1.AccountAccess.granted_at:type_name -> google.protobuf.Timestamp
	6, // 5: v1.AccountAccess.expires_at:type_name -> google.protobuf.Timestamp
	0, // 6: v1.AccountsAccess.List:input_type -> v1.AccountAccessListRequest
	1, // 7: v1.AccountsAccess.List:output_type -> v1.AccountAccessListResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_accounts_access_proto_init() }
func file_accounts_access_proto_init() {
	if File_accounts_access_proto != nil {
		return
	}
	file_options_proto_init()
	file_spec_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_accounts_access_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountAccessListRequest); i {
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
		file_accounts_access_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountAccessListResponse); i {
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
		file_accounts_access_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountAccess); i {
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
			RawDescriptor: file_accounts_access_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_accounts_access_proto_goTypes,
		DependencyIndexes: file_accounts_access_proto_depIdxs,
		MessageInfos:      file_accounts_access_proto_msgTypes,
	}.Build()
	File_accounts_access_proto = out.File
	file_accounts_access_proto_rawDesc = nil
	file_accounts_access_proto_goTypes = nil
	file_accounts_access_proto_depIdxs = nil
}