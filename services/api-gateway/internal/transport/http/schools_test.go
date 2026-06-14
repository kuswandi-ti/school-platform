package httptransport

import (
	"bytes"
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"net/http/httptest"
	schoolcorev1 "school-platform/packages/proto/gen/go/schoolcore/v1"
	"school-platform/services/api-gateway/internal/middleware"
	"school-platform/services/api-gateway/internal/response"
	"testing"
)

type schoolClientStub struct {
	createRequest  *schoolcorev1.CreateSchoolRequest
	createResponse *schoolcorev1.CreateSchoolResponse
	listResponse   *schoolcorev1.ListSchoolsResponse
	err            error
}

func (s *schoolClientStub) GetCurrentFoundation(context.Context, *schoolcorev1.GetCurrentFoundationRequest, ...grpc.CallOption) (*schoolcorev1.GetCurrentFoundationResponse, error) {
	return nil, s.err
}
func (s *schoolClientStub) ListSchools(context.Context, *schoolcorev1.ListSchoolsRequest, ...grpc.CallOption) (*schoolcorev1.ListSchoolsResponse, error) {
	return s.listResponse, s.err
}
func (s *schoolClientStub) CreateSchool(_ context.Context, r *schoolcorev1.CreateSchoolRequest, _ ...grpc.CallOption) (*schoolcorev1.CreateSchoolResponse, error) {
	s.createRequest = r
	return s.createResponse, s.err
}
func (s *schoolClientStub) UpdateSchool(context.Context, *schoolcorev1.UpdateSchoolRequest, ...grpc.CallOption) (*schoolcorev1.UpdateSchoolResponse, error) {
	return nil, s.err
}
func TestCreateSchoolHandlerUsesStandardResponseAndActor(t *testing.T) {
	c := &schoolClientStub{createResponse: &schoolcorev1.CreateSchoolResponse{School: &schoolcorev1.School{Id: "school", SchoolCode: "SD"}}}
	r := httptest.NewRequest(http.MethodPost, "/api/v1/schools", bytes.NewBufferString(`{"school_code":"SD","name":"SD Test","school_level":"elementary","status":"active"}`))
	r = withSchoolActor(r)
	w := httptest.NewRecorder()
	CreateSchoolHandler(c).ServeHTTP(w, r)
	requireStatus(t, w, http.StatusCreated)
	var envelope response.Envelope
	requireNoError(t, json.Unmarshal(w.Body.Bytes(), &envelope))
	if envelope.Error != nil || c.createRequest.Actor.FoundationId != "foundation-id" {
		t.Fatalf("unexpected response or actor: %#v %#v", envelope, c.createRequest)
	}
}
func TestCreateSchoolHandlerMapsForbidden(t *testing.T) {
	c := &schoolClientStub{err: status.Error(codes.PermissionDenied, "forbidden")}
	r := withSchoolActor(httptest.NewRequest(http.MethodPost, "/api/v1/schools", bytes.NewBufferString(`{"school_code":"SD","name":"SD Test","school_level":"elementary","status":"active"}`)))
	w := httptest.NewRecorder()
	CreateSchoolHandler(c).ServeHTTP(w, r)
	requireError(t, w, http.StatusForbidden, response.CodeForbidden)
}
func TestCreateSchoolHandlerRejectsInvalidBody(t *testing.T) {
	r := withSchoolActor(httptest.NewRequest(http.MethodPost, "/api/v1/schools", bytes.NewBufferString(`{"school_code":"","name":""}`)))
	w := httptest.NewRecorder()
	CreateSchoolHandler(&schoolClientStub{}).ServeHTTP(w, r)
	requireError(t, w, http.StatusUnprocessableEntity, response.CodeValidationError)
}
func TestListSchoolsHandlerReturnsMeta(t *testing.T) {
	c := &schoolClientStub{listResponse: &schoolcorev1.ListSchoolsResponse{Schools: []*schoolcorev1.School{{Id: "one"}, {Id: "two"}}}}
	w := httptest.NewRecorder()
	ListSchoolsHandler(c).ServeHTTP(w, withSchoolActor(httptest.NewRequest(http.MethodGet, "/api/v1/schools", nil)))
	requireStatus(t, w, http.StatusOK)
	if !bytes.Contains(w.Body.Bytes(), []byte(`"total":2`)) {
		t.Fatalf("missing total metadata: %s", w.Body.String())
	}
}
func withSchoolActor(r *http.Request) *http.Request {
	a := middleware.ActorContext{UserID: "user-id", FoundationID: "foundation-id", Roles: []string{"admin_yayasan"}, Permissions: []string{"school.school.manage"}, RequestID: "req", CorrelationID: "corr"}
	return r.WithContext(middleware.WithActorContext(r.Context(), a))
}
func requireNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
