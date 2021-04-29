package connectors

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ibm/the-mesh-for-data/pkg/connectors/protobuf"
)

// Ensure that grpcDataCatalog implements the DataCatalog interface
var _ DataCatalog = (*MockDataCatalog)(nil)

type MockDataCatalog struct {
	dataDetails map[string]*pb.DatasetDetails
}

// NewMockDataCatalog creates an in memory catalog for testing
func NewMockDataCatalog() DataCatalog {
	return &MockDataCatalog{dataDetails: map[string]*pb.DatasetDetails{}}
}

// RegisterDatasetInfo registers an asset with the ID set as "<catalog ID>/<asset name>"
func (m *MockDataCatalog) RegisterDatasetInfo(ctx context.Context, in *pb.RegisterAssetRequest) (*pb.RegisterAssetResponse, error) {
	datasetID := fmt.Sprintf("%s/%s", in.DestinationCatalogId, in.DatasetDetails.Name)
	m.dataDetails[datasetID] = in.DatasetDetails
	return &pb.RegisterAssetResponse{AssetId: datasetID}, nil
}

func (m *MockDataCatalog) GetDatasetInfo(ctx context.Context, in *pb.CatalogDatasetRequest) (*pb.CatalogDatasetInfo, error) {
	datasetID := in.GetDatasetId()
	log.Printf("MockDataCatalog.GetDatasetInfo called with DataSetID " + datasetID)
	details, found := m.dataDetails[datasetID]
	if found {
		return &pb.CatalogDatasetInfo{DatasetId: datasetID, Details: details}, nil
	}

	return nil, fmt.Errorf("MockDataCatalog.GetDatasetInfo could not find asset %s", datasetID)
}

func (m *MockDataCatalog) Close() error {
	return nil
}
