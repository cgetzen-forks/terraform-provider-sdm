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
// source: account_resources.proto

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

// AccountResourceListRequest specifies criteria for retrieving a list of AccountResource records.
type AccountResourceListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Paging parameters for the query.
	Meta *ListRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A human-readable filter query string.
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *AccountResourceListRequest) Reset() {
	*x = AccountResourceListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountResourceListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountResourceListRequest) ProtoMessage() {}

func (x *AccountResourceListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountResourceListRequest.ProtoReflect.Descriptor instead.
func (*AccountResourceListRequest) Descriptor() ([]byte, []int) {
	return file_account_resources_proto_rawDescGZIP(), []int{0}
}

func (x *AccountResourceListRequest) GetMeta() *ListRequestMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *AccountResourceListRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// AccountResourceListResponse returns a list of AccountResource records that meet
// the criteria of an AccountResourceListRequest.
type AccountResourceListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Paging information for the query.
	Meta *ListResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A single page of results matching the list request criteria.
	AccountResources []*AccountResource `protobuf:"bytes,2,rep,name=account_resources,json=accountResources,proto3" json:"account_resources,omitempty"`
	// Rate limit information.
	RateLimit *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
}

func (x *AccountResourceListResponse) Reset() {
	*x = AccountResourceListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountResourceListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountResourceListResponse) ProtoMessage() {}

func (x *AccountResourceListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountResourceListResponse.ProtoReflect.Descriptor instead.
func (*AccountResourceListResponse) Descriptor() ([]byte, []int) {
	return file_account_resources_proto_rawDescGZIP(), []int{1}
}

func (x *AccountResourceListResponse) GetMeta() *ListResponseMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *AccountResourceListResponse) GetAccountResources() []*AccountResource {
	if x != nil {
		return x.AccountResources
	}
	return nil
}

func (x *AccountResourceListResponse) GetRateLimit() *RateLimitMetadata {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

// AccountResource represents an individual access grant of a Account to a Resource.
type AccountResource struct {
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
	// The time this grant was created, distinct from 'granted at' in the case where access is scheduled
	// for the future. If access was granted, revoked, and granted again, this will reflect the later creation time.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *AccountResource) Reset() {
	*x = AccountResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_resources_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountResource) ProtoMessage() {}

func (x *AccountResource) ProtoReflect() protoreflect.Message {
	mi := &file_account_resources_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountResource.ProtoReflect.Descriptor instead.
func (*AccountResource) Descriptor() ([]byte, []int) {
	return file_account_resources_proto_rawDescGZIP(), []int{2}
}

func (x *AccountResource) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *AccountResource) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *AccountResource) GetGrantedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.GrantedAt
	}
	return nil
}

func (x *AccountResource) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *AccountResource) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *AccountResource) GetAccountGrantId() string {
	if x != nil {
		return x.AccountGrantId
	}
	return ""
}

func (x *AccountResource) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_account_resources_proto protoreflect.FileDescriptor

var file_account_resources_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x73,
	0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x1a, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07,
	0x01, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x3a, 0x28, 0xfa, 0xf8, 0xb3, 0x07, 0x06,
	0xd2, 0xf3, 0xb3, 0x07, 0x01, 0x2a, 0xfa, 0xf8, 0xb3, 0x07, 0x18, 0xd2, 0xf3, 0xb3, 0x07, 0x13,
	0x21, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x22, 0xa7, 0x02, 0x0a, 0x1b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x12, 0x4c, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb8, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x10, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12,
	0x62, 0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x2c, 0xf2, 0xf8, 0xb3, 0x07,
	0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xf2, 0xf8, 0xb3, 0x07, 0x06, 0xb2, 0xf4, 0xb3, 0x07, 0x01,
	0x2a, 0xf2, 0xf8, 0xb3, 0x07, 0x12, 0xb2, 0xf4, 0xb3, 0x07, 0x0d, 0x21, 0x6a, 0x73, 0x6f, 0x6e,
	0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x09, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x3a, 0x28, 0xfa, 0xf8, 0xb3, 0x07, 0x06, 0xd2, 0xf3, 0xb3, 0x07, 0x01, 0x2a,
	0xfa, 0xf8, 0xb3, 0x07, 0x18, 0xd2, 0xf3, 0xb3, 0x07, 0x13, 0x21, 0x74, 0x65, 0x72, 0x72, 0x61,
	0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0xcd, 0x03,
	0x0a, 0x0f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x29, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07,
	0x01, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x0b,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x0a, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x0a, 0x67, 0x72, 0x61,
	0x6e, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05,
	0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x09, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x45, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x09, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x23, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0,
	0xf3, 0xb3, 0x07, 0x01, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x10,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3,
	0x07, 0x01, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x61, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x45, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x3a, 0x32, 0xfa, 0xf8, 0xb3, 0x07, 0x05,
	0xa8, 0xf3, 0xb3, 0x07, 0x01, 0xfa, 0xf8, 0xb3, 0x07, 0x06, 0xd2, 0xf3, 0xb3, 0x07, 0x01, 0x2a,
	0xfa, 0xf8, 0xb3, 0x07, 0x18, 0xd2, 0xf3, 0xb3, 0x07, 0x13, 0x21, 0x74, 0x65, 0x72, 0x72, 0x61,
	0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x32, 0xcc, 0x01,
	0x0a, 0x10, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x12, 0x75, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82, 0xf9, 0xb3,
	0x07, 0x08, 0xa2, 0xf3, 0xb3, 0x07, 0x03, 0x67, 0x65, 0x74, 0x82, 0xf9, 0xb3, 0x07, 0x1a, 0xaa,
	0xf3, 0xb3, 0x07, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x41, 0xca, 0xf9, 0xb3, 0x07, 0x14,
	0xc2, 0xf9, 0xb3, 0x07, 0x0f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0xca, 0xf9, 0xb3, 0x07, 0x06, 0xca, 0xf9, 0xb3, 0x07, 0x01, 0x2a, 0xca,
	0xf9, 0xb3, 0x07, 0x18, 0xca, 0xf9, 0xb3, 0x07, 0x13, 0x21, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66,
	0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x42, 0x94, 0x01, 0x0a,
	0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6e, 0x67, 0x42, 0x18, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x50, 0x6c, 0x75, 0x6d,
	0x62, 0x69, 0x6e, 0x67, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67,
	0x64, 0x6d, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0xc2, 0x92, 0xb4, 0x07, 0x06,
	0xa2, 0x8c, 0xb4, 0x07, 0x01, 0x2a, 0xc2, 0x92, 0xb4, 0x07, 0x18, 0xa2, 0x8c, 0xb4, 0x07, 0x13,
	0x21, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_resources_proto_rawDescOnce sync.Once
	file_account_resources_proto_rawDescData = file_account_resources_proto_rawDesc
)

func file_account_resources_proto_rawDescGZIP() []byte {
	file_account_resources_proto_rawDescOnce.Do(func() {
		file_account_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_resources_proto_rawDescData)
	})
	return file_account_resources_proto_rawDescData
}

var file_account_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_account_resources_proto_goTypes = []interface{}{
	(*AccountResourceListRequest)(nil),  // 0: v1.AccountResourceListRequest
	(*AccountResourceListResponse)(nil), // 1: v1.AccountResourceListResponse
	(*AccountResource)(nil),             // 2: v1.AccountResource
	(*ListRequestMetadata)(nil),         // 3: v1.ListRequestMetadata
	(*ListResponseMetadata)(nil),        // 4: v1.ListResponseMetadata
	(*RateLimitMetadata)(nil),           // 5: v1.RateLimitMetadata
	(*timestamppb.Timestamp)(nil),       // 6: google.protobuf.Timestamp
}
var file_account_resources_proto_depIdxs = []int32{
	3, // 0: v1.AccountResourceListRequest.meta:type_name -> v1.ListRequestMetadata
	4, // 1: v1.AccountResourceListResponse.meta:type_name -> v1.ListResponseMetadata
	2, // 2: v1.AccountResourceListResponse.account_resources:type_name -> v1.AccountResource
	5, // 3: v1.AccountResourceListResponse.rate_limit:type_name -> v1.RateLimitMetadata
	6, // 4: v1.AccountResource.granted_at:type_name -> google.protobuf.Timestamp
	6, // 5: v1.AccountResource.expires_at:type_name -> google.protobuf.Timestamp
	6, // 6: v1.AccountResource.created_at:type_name -> google.protobuf.Timestamp
	0, // 7: v1.AccountResources.List:input_type -> v1.AccountResourceListRequest
	1, // 8: v1.AccountResources.List:output_type -> v1.AccountResourceListResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_account_resources_proto_init() }
func file_account_resources_proto_init() {
	if File_account_resources_proto != nil {
		return
	}
	file_options_proto_init()
	file_spec_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_account_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountResourceListRequest); i {
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
		file_account_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountResourceListResponse); i {
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
		file_account_resources_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountResource); i {
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
			RawDescriptor: file_account_resources_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_resources_proto_goTypes,
		DependencyIndexes: file_account_resources_proto_depIdxs,
		MessageInfos:      file_account_resources_proto_msgTypes,
	}.Build()
	File_account_resources_proto = out.File
	file_account_resources_proto_rawDesc = nil
	file_account_resources_proto_goTypes = nil
	file_account_resources_proto_depIdxs = nil
}
