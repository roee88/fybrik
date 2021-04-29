package connectors

import (
	"io"

	pb "github.com/ibm/the-mesh-for-data/pkg/connectors/protobuf"
)

// DataCatalog is an interface of a facade to a data catalog.
type DataCatalog interface {
	pb.DataCatalogServiceServer
	io.Closer
}
