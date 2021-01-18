module github.com/ibm/the-mesh-for-data

go 1.13

require (
	emperror.dev/errors v0.8.0
	github.com/IBM/go-sdk-core/v4 v4.10.0
	github.com/IBM/satcon-client-go v0.0.0-20210107134702-1be071817792
	github.com/ghodss/yaml v1.0.0
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-chi/render v1.0.1
	github.com/go-logr/logr v0.3.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.3
	github.com/hashicorp/vault/api v1.0.4
	github.com/machinebox/graphql v0.2.2
	github.com/matryer/is v1.4.0 // indirect
	github.com/onsi/ginkgo v1.14.2
	github.com/onsi/gomega v1.10.3
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron v1.2.0
	github.com/stretchr/testify v1.6.1
	google.golang.org/grpc v1.28.1
	gopkg.in/yaml.v2 v2.3.0
	helm.sh/helm/v3 v3.5.0
	istio.io/client-go v1.8.2
	k8s.io/api v0.20.2
	k8s.io/apimachinery v0.20.2
	k8s.io/cli-runtime v0.20.2
	k8s.io/client-go v0.20.2
	rsc.io/letsencrypt v0.0.3 // indirect
	sigs.k8s.io/cli-utils v0.22.4
	sigs.k8s.io/controller-runtime v0.8.0
	sigs.k8s.io/yaml v1.2.0
)

// replace helm.sh/helm/v3 v3.5.0 => github.com/hunchback/helm/v3 v3.5.0-hunchback

replace github.com/golang/protobuf => github.com/golang/protobuf v1.3.5

replace github.com/onsi/gomega => github.com/onsi/gomega v1.10.0

replace github.com/google/addlicense => github.com/the-mesh-for-data/addlicense v0.0.0-20200913135744-636c44b42906
