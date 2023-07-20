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

// PeeringGroupsClient is the client API for PeeringGroups service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeeringGroupsClient interface {
	// Create registers a new PeeringGroup.
	Create(ctx context.Context, in *PeeringGroupCreateRequest, opts ...grpc.CallOption) (*PeeringGroupCreateResponse, error)
	// Delete removes a PeeringGroup by ID.
	Delete(ctx context.Context, in *PeeringGroupDeleteRequest, opts ...grpc.CallOption) (*PeeringGroupDeleteResponse, error)
	// Get reads one PeeringGroup by ID. It will load all its dependencies.
	Get(ctx context.Context, in *PeeringGroupGetRequest, opts ...grpc.CallOption) (*PeeringGroupGetResponse, error)
	// List gets a list of Peering Groups.
	List(ctx context.Context, in *PeeringGroupListRequest, opts ...grpc.CallOption) (*PeeringGroupListResponse, error)
}

type peeringGroupsClient struct {
	cc grpc.ClientConnInterface
}

func NewPeeringGroupsClient(cc grpc.ClientConnInterface) PeeringGroupsClient {
	return &peeringGroupsClient{cc}
}

func (c *peeringGroupsClient) Create(ctx context.Context, in *PeeringGroupCreateRequest, opts ...grpc.CallOption) (*PeeringGroupCreateResponse, error) {
	out := new(PeeringGroupCreateResponse)
	err := c.cc.Invoke(ctx, "/v1.PeeringGroups/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peeringGroupsClient) Delete(ctx context.Context, in *PeeringGroupDeleteRequest, opts ...grpc.CallOption) (*PeeringGroupDeleteResponse, error) {
	out := new(PeeringGroupDeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.PeeringGroups/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peeringGroupsClient) Get(ctx context.Context, in *PeeringGroupGetRequest, opts ...grpc.CallOption) (*PeeringGroupGetResponse, error) {
	out := new(PeeringGroupGetResponse)
	err := c.cc.Invoke(ctx, "/v1.PeeringGroups/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peeringGroupsClient) List(ctx context.Context, in *PeeringGroupListRequest, opts ...grpc.CallOption) (*PeeringGroupListResponse, error) {
	out := new(PeeringGroupListResponse)
	err := c.cc.Invoke(ctx, "/v1.PeeringGroups/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeeringGroupsServer is the server API for PeeringGroups service.
// All implementations must embed UnimplementedPeeringGroupsServer
// for forward compatibility
type PeeringGroupsServer interface {
	// Create registers a new PeeringGroup.
	Create(context.Context, *PeeringGroupCreateRequest) (*PeeringGroupCreateResponse, error)
	// Delete removes a PeeringGroup by ID.
	Delete(context.Context, *PeeringGroupDeleteRequest) (*PeeringGroupDeleteResponse, error)
	// Get reads one PeeringGroup by ID. It will load all its dependencies.
	Get(context.Context, *PeeringGroupGetRequest) (*PeeringGroupGetResponse, error)
	// List gets a list of Peering Groups.
	List(context.Context, *PeeringGroupListRequest) (*PeeringGroupListResponse, error)
	mustEmbedUnimplementedPeeringGroupsServer()
}

// UnimplementedPeeringGroupsServer must be embedded to have forward compatible implementations.
type UnimplementedPeeringGroupsServer struct {
}

func (UnimplementedPeeringGroupsServer) Create(context.Context, *PeeringGroupCreateRequest) (*PeeringGroupCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPeeringGroupsServer) Delete(context.Context, *PeeringGroupDeleteRequest) (*PeeringGroupDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPeeringGroupsServer) Get(context.Context, *PeeringGroupGetRequest) (*PeeringGroupGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPeeringGroupsServer) List(context.Context, *PeeringGroupListRequest) (*PeeringGroupListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPeeringGroupsServer) mustEmbedUnimplementedPeeringGroupsServer() {}

// UnsafePeeringGroupsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeeringGroupsServer will
// result in compilation errors.
type UnsafePeeringGroupsServer interface {
	mustEmbedUnimplementedPeeringGroupsServer()
}

func RegisterPeeringGroupsServer(s grpc.ServiceRegistrar, srv PeeringGroupsServer) {
	s.RegisterService(&_PeeringGroups_serviceDesc, srv)
}

func _PeeringGroups_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeeringGroupCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeeringGroupsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PeeringGroups/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeeringGroupsServer).Create(ctx, req.(*PeeringGroupCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeeringGroups_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeeringGroupDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeeringGroupsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PeeringGroups/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeeringGroupsServer).Delete(ctx, req.(*PeeringGroupDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeeringGroups_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeeringGroupGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeeringGroupsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PeeringGroups/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeeringGroupsServer).Get(ctx, req.(*PeeringGroupGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeeringGroups_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeeringGroupListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeeringGroupsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PeeringGroups/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeeringGroupsServer).List(ctx, req.(*PeeringGroupListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PeeringGroups_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.PeeringGroups",
	HandlerType: (*PeeringGroupsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PeeringGroups_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PeeringGroups_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PeeringGroups_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _PeeringGroups_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peering_groups.proto",
}
