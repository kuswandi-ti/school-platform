package grpctransport

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"net"
	schoolcorev1 "school-platform/packages/proto/gen/go/schoolcore/v1"
	"school-platform/services/school-core-service/internal/domain"
	"school-platform/services/school-core-service/internal/repository"
	"school-platform/services/school-core-service/internal/usecase"
	"testing"
)

type rpcStore struct{ foundation domain.Foundation }

func (s *rpcStore) GetFoundation(context.Context, uuid.UUID) (domain.Foundation, error) {
	return s.foundation, nil
}
func (*rpcStore) ListSchools(context.Context, uuid.UUID) ([]domain.School, error) { return nil, nil }
func (*rpcStore) GetSchool(context.Context, uuid.UUID, uuid.UUID) (domain.School, error) {
	return domain.School{}, nil
}
func (*rpcStore) CreateSchool(context.Context, repository.SchoolWrite, repository.AuditContext) (domain.School, error) {
	return domain.School{}, nil
}
func (*rpcStore) UpdateSchool(context.Context, repository.SchoolWrite, repository.AuditContext) (domain.School, error) {
	return domain.School{}, nil
}
func TestSchoolCoreGRPCRoundTrip(t *testing.T) {
	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	foundationID := uuid.New()
	schoolcorev1.RegisterSchoolCoreServiceServer(server, NewSchoolCoreServer(usecase.NewService(&rpcStore{foundation: domain.Foundation{ID: foundationID, FoundationCode: "TEST", Name: "Test", Status: "active"}})))
	go func() { _ = server.Serve(listener) }()
	defer server.Stop()
	connection, err := grpc.NewClient("passthrough:///bufnet", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return listener.Dial() }), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	defer connection.Close()
	result, err := schoolcorev1.NewSchoolCoreServiceClient(connection).GetCurrentFoundation(context.Background(), &schoolcorev1.GetCurrentFoundationRequest{Actor: &schoolcorev1.ActorContext{UserId: uuid.NewString(), FoundationId: foundationID.String(), Roles: []string{"admin_yayasan"}, Permissions: []string{"school.foundation.view"}, RequestId: "req", CorrelationId: "corr"}})
	require.NoError(t, err)
	require.Equal(t, "TEST", result.Foundation.FoundationCode)
}
