package mockup

import (
	"context"

	"github.com/ibm/the-mesh-for-data/pkg/connectors"
	pb "github.com/ibm/the-mesh-for-data/pkg/connectors/protobuf"
)

func CreateDataCatalogMock() connectors.DataCatalog {

	catalog := connectors.NewMockDataCatalog()

	catalog.RegisterDatasetInfo(context.TODO(), &pb.RegisterAssetRequest{
		DestinationCatalogId: "s3-external",
		DatasetDetails: &pb.DatasetDetails{
			Name:       "default-dataset",
			DataFormat: "parquet",
			Geo:        "Germany",
			DataStore: &pb.DataStore{
				Type: pb.DataStore_S3,
				Name: "cos",
				S3: &pb.S3DataStore{
					Endpoint:  "s3.eu-gb.cloud-object-storage.appdomain.cloud",
					Bucket:    "m4d-test-bucket",
					ObjectKey: "small.parq",
				},
			},
			CredentialsInfo: &pb.CredentialsInfo{
				VaultSecretPath: "/v1/kubernetes-secrets/creds-secret-name?namespace=m4d-system",
			},
			Metadata: &pb.DatasetMetadata{DatasetTags: []string{"PI"}},
		},
	})

	catalog.RegisterDatasetInfo(context.TODO(), &pb.RegisterAssetRequest{
		DestinationCatalogId: "s3",
		DatasetDetails: &pb.DatasetDetails{
			Name:       "allow-dataset",
			DataFormat: "parquet",
			Geo:        "theshire",
			DataStore: &pb.DataStore{
				Type: pb.DataStore_S3,
				Name: "cos",
				S3: &pb.S3DataStore{
					Endpoint:  "s3.eu-gb.cloud-object-storage.appdomain.cloud",
					Bucket:    "m4d-test-bucket",
					ObjectKey: "small.parq",
				},
			},
			CredentialsInfo: &pb.CredentialsInfo{
				VaultSecretPath: "/v1/kubernetes-secrets/creds-secret-name?namespace=m4d-system",
			},
			Metadata: &pb.DatasetMetadata{DatasetTags: []string{"PI"}},
		},
	})

	catalog.RegisterDatasetInfo(context.TODO(), &pb.RegisterAssetRequest{
		DestinationCatalogId: "s3",
		DatasetDetails: &pb.DatasetDetails{
			Name:       "deny-dataset",
			DataFormat: "parquet",
			Geo:        "theshire",
			DataStore: &pb.DataStore{
				Type: pb.DataStore_S3,
				Name: "cos",
				S3: &pb.S3DataStore{
					Endpoint:  "s3.eu-gb.cloud-object-storage.appdomain.cloud",
					Bucket:    "m4d-test-bucket",
					ObjectKey: "small.parq",
				},
			},
			CredentialsInfo: &pb.CredentialsInfo{
				VaultSecretPath: "/v1/kubernetes-secrets/creds-secret-name?namespace=m4d-system",
			},
			Metadata: &pb.DatasetMetadata{DatasetTags: []string{"PI"}},
		},
	})

	catalog.RegisterDatasetInfo(context.TODO(), &pb.RegisterAssetRequest{
		DestinationCatalogId: "s3-csv",
		DatasetDetails: &pb.DatasetDetails{
			Name:       "allow-dataset",
			DataFormat: "csv",
			Geo:        "theshire",
			DataStore: &pb.DataStore{
				Type: pb.DataStore_S3,
				Name: "cos",
				S3: &pb.S3DataStore{
					Endpoint:  "s3.eu-gb.cloud-object-storage.appdomain.cloud",
					Bucket:    "m4d-test-bucket",
					ObjectKey: "small.csv",
				},
			},
			CredentialsInfo: &pb.CredentialsInfo{
				VaultSecretPath: "/v1/kubernetes-secrets/creds-secret-name?namespace=m4d-system",
			},
			Metadata: &pb.DatasetMetadata{DatasetTags: []string{"PI"}},
		},
	})

	catalog.RegisterDatasetInfo(context.TODO(), &pb.RegisterAssetRequest{
		DestinationCatalogId: "db2",
		DatasetDetails: &pb.DatasetDetails{
			Name:       "default-dataset",
			DataFormat: "table",
			Geo:        "theshire",
			DataStore: &pb.DataStore{
				Type: pb.DataStore_DB2,
				Name: "db2",
				Db2: &pb.Db2DataStore{
					Database: "BLUDB",
					Table:    "NQD60833.SMALL",
					Url:      "dashdb-txn-sbox-yp-lon02-02.services.eu-gb.bluemix.net",
					Port:     "50000",
					Ssl:      "false",
				},
			},
			CredentialsInfo: &pb.CredentialsInfo{
				VaultSecretPath: "/v1/kubernetes-secrets/creds-secret-name?namespace=m4d-system",
			},
			Metadata: &pb.DatasetMetadata{},
		},
	})

	catalog.RegisterDatasetInfo(context.TODO(), &pb.RegisterAssetRequest{
		DestinationCatalogId: "db2",
		DatasetDetails: &pb.DatasetDetails{
			Name:       "allow-dataset",
			DataFormat: "table",
			Geo:        "theshire",
			DataStore: &pb.DataStore{
				Type: pb.DataStore_DB2,
				Name: "db2",
				Db2: &pb.Db2DataStore{
					Database: "BLUDB",
					Table:    "NQD60833.SMALL",
					Url:      "dashdb-txn-sbox-yp-lon02-02.services.eu-gb.bluemix.net",
					Port:     "50000",
					Ssl:      "false",
				},
			},
			CredentialsInfo: &pb.CredentialsInfo{
				VaultSecretPath: "/v1/kubernetes-secrets/creds-secret-name?namespace=m4d-system",
			},
			Metadata: &pb.DatasetMetadata{},
		},
	})

	catalog.RegisterDatasetInfo(context.TODO(), &pb.RegisterAssetRequest{
		DestinationCatalogId: "kafka",
		DatasetDetails: &pb.DatasetDetails{
			Name:       "allow-dataset",
			DataFormat: "json",
			Geo:        "theshire",
			DataStore: &pb.DataStore{
				Type: pb.DataStore_KAFKA,
				Name: "kafka",
				Kafka: &pb.KafkaDataStore{
					TopicName:             "topic",
					SecurityProtocol:      "SASL_SSL",
					SaslMechanism:         "SCRAM-SHA-512",
					SslTruststore:         "xyz123",
					SslTruststorePassword: "passwd",
					SchemaRegistry:        "kafka-registry",
					BootstrapServers:      "http://kafka-servers",
					KeyDeserializer:       "io.confluent.kafka.serializers.json.KafkaJsonSchemaDeserializer",
					ValueDeserializer:     "io.confluent.kafka.serializers.json.KafkaJsonSchemaDeserializer",
				},
			},
			CredentialsInfo: &pb.CredentialsInfo{
				VaultSecretPath: "/v1/kubernetes-secrets/creds-secret-name?namespace=m4d-system",
			},
			Metadata: &pb.DatasetMetadata{},
		},
	})

	return catalog
}
