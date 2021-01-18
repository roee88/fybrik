module github.com/ibm/the-mesh-for-data/connectors/opa

require github.com/ibm/the-mesh-for-data v0.0.0

replace github.com/ibm/the-mesh-for-data v0.0.0 => ../..

go 1.13

require (
	github.com/golang/protobuf v1.4.3
	github.com/hashicorp/go-retryablehttp v0.6.6
	github.com/stretchr/testify v1.6.1
	google.golang.org/grpc v1.31.1
	gotest.tools v2.2.0+incompatible
)
