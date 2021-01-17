package connector

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/ibm/the-mesh-for-data/connectors/katalog/pkg/taxonomy"
	connectors "github.com/ibm/the-mesh-for-data/pkg/connectors/protobuf"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kclient "sigs.k8s.io/controller-runtime/pkg/client"
)

// TODO(roee88): This is a temporary implementation of a catalog connector to
// Katalog. It is here to map between Katalog CRDs to the connectors proto
// definitions. Eventually, the connectors proto definitions won't hardcode so
// much and rely on validating against a configured OpenAPI spec instead, making
// most of the code in this file unnecessary.

type DataCatalogService struct {
	client kclient.Client
}

func (s *DataCatalogService) GetDatasetInfo(ctx context.Context, req *connectors.CatalogDatasetRequest) (*connectors.CatalogDatasetInfo, error) {
	namespace, name, err := splitNamespacedName(req.DatasetId)
	if err != nil {
		return nil, err
	}
	asset, err := s.getAsset(ctx, namespace, name)
	if err != nil {
		return nil, err
	}

	datastore, err := builDataStore(asset)
	if err != nil {
		return nil, err
	}

	return &connectors.CatalogDatasetInfo{
		DatasetId: req.DatasetId,
		Details: &connectors.DatasetDetails{
			Name:       req.DatasetId,
			DataOwner:  emptyIfNil(asset.Spec.Security.Owner),
			DataFormat: emptyIfNil(asset.Spec.Details.DataFormat),
			Geo:        emptyIfNil(asset.Spec.Security.Geography),
			DataStore:  datastore,
			Metadata:   buildDatasetMetadata(asset),
		},
	}, nil
}

func buildDatasetMetadata(asset *Asset) *connectors.DatasetMetadata {
	security := asset.Spec.Security

	var namedMetadata map[string]string
	if security.NamedMetadata != nil {
		namedMetadata = security.NamedMetadata.AdditionalProperties
	}

	componentsMetadata := map[string]*connectors.DataComponentMetadata{}
	for componentName, componentValue := range security.ComponentsMetadata.AdditionalProperties {
		var componentNamedMetadata map[string]string
		if componentValue.NamedMetadata != nil {
			componentNamedMetadata = componentValue.NamedMetadata.AdditionalProperties
		}
		componentsMetadata[componentName] = &connectors.DataComponentMetadata{
			ComponentType: "column",
			Tags:          emptyArrayIfNil(componentValue.Tags),
			NamedMetadata: componentNamedMetadata,
		}
	}

	return &connectors.DatasetMetadata{
		DatasetTags:          emptyArrayIfNil(security.Tags),
		DatasetNamedMetadata: namedMetadata,
		ComponentsMetadata:   componentsMetadata,
	}
}

func builDataStore(asset *Asset) (*connectors.DataStore, error) {
	connection := asset.Spec.Details.Connection.(map[string]interface{})
	switch connectionType := connection["type"].(string); connectionType {
	case "s3":
		s3 := &taxonomy.S3{}
		if err := decodeToStruct(connection[connectionType], s3); err != nil {
			return nil, err
		}
		return &connectors.DataStore{
			Type: connectors.DataStore_S3,
			Name: asset.Name,
			S3: &connectors.S3DataStore{
				Endpoint:  s3.Endpoint,
				Bucket:    s3.Bucket,
				ObjectKey: s3.ObjectKey,
				Region:    emptyIfNil(s3.Region),
			},
		}, nil
	case "kafka":
		kafka := &taxonomy.Kafka{}
		if err := decodeToStruct(connection[connectionType], kafka); err != nil {
			return nil, err
		}
		return &connectors.DataStore{
			Type: connectors.DataStore_KAFKA,
			Name: asset.Name,
			Kafka: &connectors.KafkaDataStore{
				TopicName:             emptyIfNil(kafka.TopicName),
				BootstrapServers:      emptyIfNil(kafka.BootstrapServers),
				SchemaRegistry:        emptyIfNil(kafka.SchemaRegistry),
				KeyDeserializer:       emptyIfNil(kafka.KeyDeserializer),
				ValueDeserializer:     emptyIfNil(kafka.ValueDeserializer),
				SecurityProtocol:      emptyIfNil(kafka.SecurityProtocol),
				SaslMechanism:         emptyIfNil(kafka.SaslMechanism),
				SslTruststore:         emptyIfNil(kafka.SslTruststore),
				SslTruststorePassword: emptyIfNil(kafka.SslTruststorePassword),
			},
		}, nil
	case "db2":
		db2 := &taxonomy.DB2{}
		if err := decodeToStruct(connection[connectionType], db2); err != nil {
			return nil, err
		}
		return &connectors.DataStore{
			Type: connectors.DataStore_DB2,
			Name: asset.Name,
			Db2: &connectors.Db2DataStore{
				Url:      emptyIfNil(db2.Url),
				Database: emptyIfNil(db2.Database),
				Table:    emptyIfNil(db2.Table),
				Port:     emptyIfNil(db2.Port),
				Ssl:      emptyIfNil(db2.Ssl),
			},
		}, nil
	default:
		return nil, errors.New("unknown datastore type")
	}
}

func (s *DataCatalogService) getAsset(ctx context.Context, namespace string, name string) (*Asset, error) {
	// Read asset as unstructured
	object := &unstructured.Unstructured{}
	object.SetGroupVersionKind(schema.GroupVersionKind{Group: GroupVersion.Group, Version: GroupVersion.Version, Kind: "Asset"})
	object.SetNamespace(namespace)
	object.SetName(name)

	objectKey, err := kclient.ObjectKeyFromObject(object)
	if err != nil {
		return nil, err
	}

	err = s.client.Get(ctx, objectKey, object)
	if err != nil {
		return nil, err
	}

	// Decode into an Asset object
	asset := &Asset{}
	bytes, err := object.MarshalJSON()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, asset)
	if err != nil {
		return nil, err
	}

	return asset, nil
}
