package connector

import (
	"context"
	"encoding/base64"
	"encoding/json"

	"github.com/pkg/errors"

	corev1 "k8s.io/api/core/v1"

	connectors "github.com/ibm/the-mesh-for-data/pkg/connectors/protobuf"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kclient "sigs.k8s.io/controller-runtime/pkg/client"
)

type DataCredentialsService struct {
	client kclient.Client
}

func (s *DataCredentialsService) GetCredentialsInfo(ctx context.Context, req *connectors.DatasetCredentialsRequest) (*connectors.DatasetCredentials, error) {
	namespace, name, err := splitNamespacedName(req.DatasetId)
	if err != nil {
		return nil, err
	}

	// Read the secret
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
	}
	objectKey, err := kclient.ObjectKeyFromObject(secret)
	if err != nil {
		return nil, err
	}
	err = s.client.Get(ctx, objectKey, secret)
	if err != nil {
		return nil, err
	}

	// Decode the secret data
	data, err := base64.StdEncoding.DecodeString(string(secret.Data["main"]))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to decode secret")
	}

	// Load the secret data as a Credentials object
	credentials := &Credentials{}
	err = json.Unmarshal(data, &credentials)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse credentials from secret")
	}

	creds := &connectors.Credentials{}

	switch credentials.Spec.Type {
	case "basic":
		creds.Username = credentials.Spec.Basic.Username
		creds.Password = credentials.Spec.Basic.Password
	case "accessKeys":
		creds.AccessKey = credentials.Spec.AccessKeys.AccessKey
		creds.SecretKey = credentials.Spec.AccessKeys.SecretKey
	case "apiKey":
		creds.ApiKey = *credentials.Spec.ApiKey
	default:
		return nil, errors.New("unknown credentials type")
	}

	return &connectors.DatasetCredentials{DatasetId: req.DatasetId, Creds: creds}, nil
}
