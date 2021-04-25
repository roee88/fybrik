package connectors

import (
	"context"
	"fmt"
	"io"
	"time"

	"emperror.dev/errors"
	"github.com/go-logr/logr"
	pb "github.com/ibm/the-mesh-for-data/pkg/connectors/protobuf"
	"go.uber.org/multierr"
	"google.golang.org/grpc"
)

// PolicyManager is an interface of a facade to connect to a policy manager.
// Implementations should derive from policyManagerBase.
type PolicyManager interface {
	pb.PolicyManagerServiceServer
	io.Closer
}

// NewMultiPolicyManager creates a PolicyManager facade that combines results from multiple policy managers
// You must call .Close() when you are done using the created instance
func NewMultiPolicyManager(log logr.Logger, managers ...PolicyManager) PolicyManager {
	return &multiPolicyManager{
		managers: managers,
		log:      log,
	}
}

// NewPolicyManagerGrpc creates a PolicyManager facade that connects to a GRPC service
// You must call .Close() when you are done using the created instance
func NewGrpcPolicyManager(log logr.Logger, name string, connectionURL string, connectionTimeout time.Duration) (PolicyManager, error) {
	ctx, cancel := context.WithTimeout(context.Background(), connectionTimeout)
	defer cancel()
	connection, err := grpc.DialContext(ctx, connectionURL, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("NewGrpcPolicyManager failed when connecting to %s", connectionURL))
	}
	return &grpcPolicyManager{
		name:              name,
		connectionURL:     connectionURL,
		connectionTimeout: connectionTimeout,
		log:               log,
		client:            pb.NewPolicyManagerServiceClient(connection),
		connection:        connection,
	}, nil
}

var _ PolicyManager = (*grpcPolicyManager)(nil)

type grpcPolicyManager struct {
	pb.UnimplementedPolicyManagerServiceServer

	name              string
	connectionURL     string
	connectionTimeout time.Duration
	log               logr.Logger

	connection *grpc.ClientConn
	client     pb.PolicyManagerServiceClient
}

func (m *grpcPolicyManager) GetPoliciesDecisions(ctx context.Context, in *pb.ApplicationContext) (*pb.PoliciesDecisions, error) {
	result, err := m.client.GetPoliciesDecisions(ctx, in)
	return result, errors.Wrap(err, fmt.Sprintf("get policies decisions from %s failed", m.name))
}

func (m *grpcPolicyManager) Close() error {
	return m.connection.Close()
}

var _ PolicyManager = (*multiPolicyManager)(nil)

type multiPolicyManager struct {
	pb.UnimplementedPolicyManagerServiceServer

	managers []PolicyManager
	log      logr.Logger
}

func (m *multiPolicyManager) GetPoliciesDecisions(ctx context.Context, in *pb.ApplicationContext) (*pb.PoliciesDecisions, error) {
	var resultError error
	result := &pb.PoliciesDecisions{}

	for _, manager := range m.managers {
		decisions, err := manager.GetPoliciesDecisions(ctx, in)
		if !multierr.AppendInto(&resultError, err) {
			result.ComponentVersions = append(result.ComponentVersions, decisions.ComponentVersions...)
			result.GeneralDecisions = append(result.GeneralDecisions, decisions.GeneralDecisions...)
			// TODO: replace with logic to combine by DatasetId and AccessOperation
			result.DatasetDecisions = append(result.DatasetDecisions, decisions.DatasetDecisions...)
		}
	}

	return result, errors.Wrap(resultError, fmt.Sprintf("multi policy manager returned %d errors", len(multierr.Errors(resultError))))
}

func (m *multiPolicyManager) Close() error {
	var err error
	for _, manager := range m.managers {
		multierr.AppendInto(&err, manager.Close())
	}
	return err
}
