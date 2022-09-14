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
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// RemoteIdentityGroupsClient is the client API for RemoteIdentityGroups service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemoteIdentityGroupsClient interface {
	// Get reads one RemoteIdentityGroup by ID.
	Get(ctx context.Context, in *RemoteIdentityGroupGetRequest, opts ...grpc.CallOption) (*RemoteIdentityGroupGetResponse, error)
	// List gets a list of RemoteIdentityGroups matching a given set of criteria.
	List(ctx context.Context, in *RemoteIdentityGroupListRequest, opts ...grpc.CallOption) (*RemoteIdentityGroupListResponse, error)
}

type remoteIdentityGroupsClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoteIdentityGroupsClient(cc grpc.ClientConnInterface) RemoteIdentityGroupsClient {
	return &remoteIdentityGroupsClient{cc}
}

func (c *remoteIdentityGroupsClient) Get(ctx context.Context, in *RemoteIdentityGroupGetRequest, opts ...grpc.CallOption) (*RemoteIdentityGroupGetResponse, error) {
	out := new(RemoteIdentityGroupGetResponse)
	err := c.cc.Invoke(ctx, "/v1.RemoteIdentityGroups/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteIdentityGroupsClient) List(ctx context.Context, in *RemoteIdentityGroupListRequest, opts ...grpc.CallOption) (*RemoteIdentityGroupListResponse, error) {
	out := new(RemoteIdentityGroupListResponse)
	err := c.cc.Invoke(ctx, "/v1.RemoteIdentityGroups/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteIdentityGroupsServer is the server API for RemoteIdentityGroups service.
// All implementations must embed UnimplementedRemoteIdentityGroupsServer
// for forward compatibility
type RemoteIdentityGroupsServer interface {
	// Get reads one RemoteIdentityGroup by ID.
	Get(context.Context, *RemoteIdentityGroupGetRequest) (*RemoteIdentityGroupGetResponse, error)
	// List gets a list of RemoteIdentityGroups matching a given set of criteria.
	List(context.Context, *RemoteIdentityGroupListRequest) (*RemoteIdentityGroupListResponse, error)
	mustEmbedUnimplementedRemoteIdentityGroupsServer()
}

// UnimplementedRemoteIdentityGroupsServer must be embedded to have forward compatible implementations.
type UnimplementedRemoteIdentityGroupsServer struct {
}

func (UnimplementedRemoteIdentityGroupsServer) Get(context.Context, *RemoteIdentityGroupGetRequest) (*RemoteIdentityGroupGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRemoteIdentityGroupsServer) List(context.Context, *RemoteIdentityGroupListRequest) (*RemoteIdentityGroupListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRemoteIdentityGroupsServer) mustEmbedUnimplementedRemoteIdentityGroupsServer() {}

// UnsafeRemoteIdentityGroupsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemoteIdentityGroupsServer will
// result in compilation errors.
type UnsafeRemoteIdentityGroupsServer interface {
	mustEmbedUnimplementedRemoteIdentityGroupsServer()
}

func RegisterRemoteIdentityGroupsServer(s grpc.ServiceRegistrar, srv RemoteIdentityGroupsServer) {
	s.RegisterService(&_RemoteIdentityGroups_serviceDesc, srv)
}

func _RemoteIdentityGroups_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteIdentityGroupGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteIdentityGroupsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.RemoteIdentityGroups/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteIdentityGroupsServer).Get(ctx, req.(*RemoteIdentityGroupGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteIdentityGroups_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteIdentityGroupListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteIdentityGroupsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.RemoteIdentityGroups/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteIdentityGroupsServer).List(ctx, req.(*RemoteIdentityGroupListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RemoteIdentityGroups_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.RemoteIdentityGroups",
	HandlerType: (*RemoteIdentityGroupsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _RemoteIdentityGroups_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _RemoteIdentityGroups_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "remote_identity_groups.proto",
}
