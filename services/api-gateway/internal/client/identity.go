package client

import (
	"context"

	identityv1 "school-platform/packages/proto/gen/go/identity/v1"

	"google.golang.org/grpc"
)

type Identity interface {
	Login(ctx context.Context, in *identityv1.LoginRequest, opts ...grpc.CallOption) (*identityv1.LoginResponse, error)
	Refresh(ctx context.Context, in *identityv1.RefreshRequest, opts ...grpc.CallOption) (*identityv1.RefreshResponse, error)
}
