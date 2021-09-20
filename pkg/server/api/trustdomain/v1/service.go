package trustdomain

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/common/protoutil"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/pkg/server/api"
	"github.com/spiffe/spire/pkg/server/api/rpccontext"
	"github.com/spiffe/spire/pkg/server/datastore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	trustdomainv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/trustdomain/v1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
)

// Config is the service configuration.
type Config struct {
	DataStore   datastore.DataStore
	TrustDomain spiffeid.TrustDomain
}

// Service implements the v1 trustdomain service.
type Service struct {
	trustdomainv1.UnsafeTrustDomainServer

	ds datastore.DataStore
	td spiffeid.TrustDomain
}

// New creates a new trustdomain service.
func New(config Config) *Service {
	return &Service{
		ds: config.DataStore,
		td: config.TrustDomain,
	}
}

// RegisterService registers the trustdomain service on the gRPC server.
func RegisterService(s *grpc.Server, service *Service) {
	trustdomainv1.RegisterTrustDomainServer(s, service)
}

func (s *Service) ListFederationRelationships(ctx context.Context, req *trustdomainv1.ListFederationRelationshipsRequest) (*trustdomainv1.ListFederationRelationshipsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func (s *Service) GetFederationRelationship(ctx context.Context, req *trustdomainv1.GetFederationRelationshipRequest) (*types.FederationRelationship, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func (s *Service) BatchCreateFederationRelationship(ctx context.Context, req *trustdomainv1.BatchCreateFederationRelationshipRequest) (*trustdomainv1.BatchCreateFederationRelationshipResponse, error) {
	var results []*trustdomainv1.BatchCreateFederationRelationshipResponse_Result
	for _, eachRelationship := range req.FederationRelationship {
		r := s.createFederationRelationship(ctx, eachRelationship, req.OutputMask)
		results = append(results, r)
		rpccontext.AuditRPCWithTypesStatus(ctx, r.Status, func() logrus.Fields {
			return fieldsFromRelationshipProto(eachRelationship, nil)
		})
	}

	return &trustdomainv1.BatchCreateFederationRelationshipResponse{
		Results: results,
	}, nil
}

func (s *Service) BatchUpdateFederationRelationship(ctx context.Context, req *trustdomainv1.BatchUpdateFederationRelationshipRequest) (*trustdomainv1.BatchUpdateFederationRelationshipResponse, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func (s *Service) BatchDeleteFederationRelationship(ctx context.Context, req *trustdomainv1.BatchDeleteFederationRelationshipRequest) (*trustdomainv1.BatchDeleteFederationRelationshipResponse, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func (s *Service) RefreshBundle(ctx context.Context, req *trustdomainv1.RefreshBundleRequest) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func (s *Service) createFederationRelationship(ctx context.Context, f *types.FederationRelationship, outputMask *types.FederationRelationshipMask) *trustdomainv1.BatchCreateFederationRelationshipResponse_Result {
	log := rpccontext.Logger(ctx)

	dsFederationRelationship, err := api.ProtoToFederationRelationship(f)
	if err != nil {
		return &trustdomainv1.BatchCreateFederationRelationshipResponse_Result{
			Status: api.MakeStatus(log, codes.InvalidArgument, "failed to convert federation relationship", err),
		}
	}

	log = log.WithField(telemetry.TrustDomainID, dsFederationRelationship.TrustDomain.String())

	if s.td.Compare(dsFederationRelationship.TrustDomain) == 0 {
		return &trustdomainv1.BatchCreateFederationRelationshipResponse_Result{
			Status: api.MakeStatus(log, codes.InvalidArgument, "unable to create federation relationship for server trust domain", nil),
		}
	}

	resp, err := s.ds.CreateFederationRelationship(ctx, dsFederationRelationship)
	if err != nil {
		return &trustdomainv1.BatchCreateFederationRelationshipResponse_Result{
			Status: api.MakeStatus(log, codes.Internal, "failed to create federation relationship", err),
		}
	}

	tFederationRelationship, err := api.FederationRelationshipToProto(resp, outputMask)
	if err != nil {
		return &trustdomainv1.BatchCreateFederationRelationshipResponse_Result{
			Status: api.MakeStatus(log, codes.Internal, "failed to convert datastore response", err),
		}
	}

	log.Debug("federation relationship created")

	return &trustdomainv1.BatchCreateFederationRelationshipResponse_Result{
		Status:                 api.OK(),
		FederationRelationship: tFederationRelationship,
	}
}

func fieldsFromRelationshipProto(proto *types.FederationRelationship, mask *types.FederationRelationshipMask) logrus.Fields {
	fields := logrus.Fields{}

	if mask == nil {
		mask = protoutil.AllTrueFederationRelationshipMask
	}

	if proto == nil {
		return fields
	}

	if proto.TrustDomain != "" {
		fields[telemetry.TrustDomainID] = proto.TrustDomain
	}

	if mask.BundleEndpointUrl {
		fields[telemetry.BundleEndpointURL] = proto.BundleEndpointUrl
	}

	if mask.BundleEndpointProfile {
		switch profile := proto.BundleEndpointProfile.(type) {
		case *types.FederationRelationship_HttpsWeb:
			fields[telemetry.BundleEndpointProfile] = datastore.BundleEndpointWeb
		case *types.FederationRelationship_HttpsSpiffe:
			fields[telemetry.BundleEndpointProfile] = datastore.BundleEndpointSPIFFE
			fields[telemetry.EndpointSpiffeID] = profile.HttpsSpiffe.EndpointSpiffeId

			bundleFields := api.FieldsFromBundleProto(proto.GetHttpsSpiffe().Bundle, nil)
			for key, value := range bundleFields {
				fields[key] = value
			}
		}
	}

	return fields
}
