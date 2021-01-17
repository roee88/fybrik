// Package taxonomy provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package taxonomy

// Authentication defines model for Authentication.
type Authentication struct {

	// The access key is also known as AccessKeyId
	AccessKey *string `json:"accessKey,omitempty"`

	// API key used in various IAM enabled services
	ApiKey   *string `json:"apiKey,omitempty"`
	Password *string `json:"password,omitempty"`

	// The secret key is also known as SecretAccessKey
	SecretKey *string `json:"secretKey,omitempty"`
	Username  *string `json:"username,omitempty"`
}

// Connection defines model for Connection.
type Connection interface{}

// DB2 defines model for DB2.
type DB2 struct {
	Database *string `json:"database,omitempty"`
	Port     *string `json:"port,omitempty"`
	Ssl      *string `json:"ssl,omitempty"`
	Table    *string `json:"table,omitempty"`
	Url      *string `json:"url,omitempty"`
}

// Kafka defines model for Kafka.
type Kafka struct {
	BootstrapServers      *string `json:"bootstrap_servers,omitempty"`
	KeyDeserializer       *string `json:"key_deserializer,omitempty"`
	SaslMechanism         *string `json:"sasl_mechanism,omitempty"`
	SchemaRegistry        *string `json:"schema_registry,omitempty"`
	SecurityProtocol      *string `json:"security_protocol,omitempty"`
	SslTruststore         *string `json:"ssl_truststore,omitempty"`
	SslTruststorePassword *string `json:"ssl_truststore_password,omitempty"`
	TopicName             *string `json:"topic_name,omitempty"`
	ValueDeserializer     *string `json:"value_deserializer,omitempty"`
}

// S3 defines model for S3.
type S3 struct {
	Bucket    string  `json:"bucket"`
	Endpoint  string  `json:"endpoint"`
	ObjectKey string  `json:"object_key"`
	Region    *string `json:"region,omitempty"`
}
