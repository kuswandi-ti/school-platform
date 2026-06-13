package client

import (
	"context"
	"google.golang.org/grpc"
	schoolcorev1 "school-platform/packages/proto/gen/go/schoolcore/v1"
)

type SchoolCore interface {
	GetCurrentFoundation(context.Context, *schoolcorev1.GetCurrentFoundationRequest, ...grpc.CallOption) (*schoolcorev1.GetCurrentFoundationResponse, error)
	ListSchools(context.Context, *schoolcorev1.ListSchoolsRequest, ...grpc.CallOption) (*schoolcorev1.ListSchoolsResponse, error)
	CreateSchool(context.Context, *schoolcorev1.CreateSchoolRequest, ...grpc.CallOption) (*schoolcorev1.CreateSchoolResponse, error)
	UpdateSchool(context.Context, *schoolcorev1.UpdateSchoolRequest, ...grpc.CallOption) (*schoolcorev1.UpdateSchoolResponse, error)
}
