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

// RolesAccessClient is the client API for RolesAccess service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RolesAccessClient interface {
	// List gets a list of RoleAccess records matching a given set of criteria.
	List(ctx context.Context, in *RoleAccessListRequest, opts ...grpc.CallOption) (*RoleAccessListResponse, error)
}

type rolesAccessClient struct {
	cc grpc.ClientConnInterface
}

func NewRolesAccessClient(cc grpc.ClientConnInterface) RolesAccessClient {
	return &rolesAccessClient{cc}
}

func (c *rolesAccessClient) List(ctx context.Context, in *RoleAccessListRequest, opts ...grpc.CallOption) (*RoleAccessListResponse, error) {
	out := new(RoleAccessListResponse)
	err := c.cc.Invoke(ctx, "/v1.RolesAccess/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RolesAccessServer is the server API for RolesAccess service.
// All implementations must embed UnimplementedRolesAccessServer
// for forward compatibility
type RolesAccessServer interface {
	// List gets a list of RoleAccess records matching a given set of criteria.
	List(context.Context, *RoleAccessListRequest) (*RoleAccessListResponse, error)
	mustEmbedUnimplementedRolesAccessServer()
}

// UnimplementedRolesAccessServer must be embedded to have forward compatible implementations.
type UnimplementedRolesAccessServer struct {
}

func (UnimplementedRolesAccessServer) List(context.Context, *RoleAccessListRequest) (*RoleAccessListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRolesAccessServer) mustEmbedUnimplementedRolesAccessServer() {}

// UnsafeRolesAccessServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RolesAccessServer will
// result in compilation errors.
type UnsafeRolesAccessServer interface {
	mustEmbedUnimplementedRolesAccessServer()
}

func RegisterRolesAccessServer(s grpc.ServiceRegistrar, srv RolesAccessServer) {
	s.RegisterService(&_RolesAccess_serviceDesc, srv)
}

func _RolesAccess_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleAccessListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesAccessServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.RolesAccess/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesAccessServer).List(ctx, req.(*RoleAccessListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RolesAccess_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.RolesAccess",
	HandlerType: (*RolesAccessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _RolesAccess_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "roles_access.proto",
}