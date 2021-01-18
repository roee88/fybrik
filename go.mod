module github.com/ibm/the-mesh-for-data

go 1.13

require (
	emperror.dev/errors v0.7.0
	github.com/IBM/go-sdk-core v1.1.0
	github.com/IBM/satcon-client-go v0.0.0-20210107134702-1be071817792
	github.com/Microsoft/go-winio v0.4.15 // indirect
	github.com/containerd/continuity v0.0.0-20201119173150-04c754faca46 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-chi/render v1.0.1
	github.com/go-logr/logr v0.2.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.3
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/hashicorp/vault/api v1.0.4
	github.com/machinebox/graphql v0.2.2
	github.com/matryer/is v1.4.0 // indirect
	github.com/onsi/ginkgo v1.14.2
	github.com/onsi/gomega v1.10.3
	github.com/opencontainers/runc v1.0.0-rc9 // indirect
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron v1.2.0
	github.com/stretchr/testify v1.6.1
	golang.org/x/sys v0.0.0-20201119102817-f84b799fce68 // indirect
	google.golang.org/grpc v1.28.1
	gopkg.in/yaml.v2 v2.3.0
	helm.sh/helm/v3 v3.5.0
	istio.io/api v0.0.0-20200723170824-3c2193e74947 // indirect
	istio.io/client-go v0.0.0-20200128004641-c87542c7dc1d
	istio.io/gogo-genproto v0.0.0-20191009201739-17d570f95998 // indirect
	k8s.io/api v0.20.1
	k8s.io/apimachinery v0.20.1
	k8s.io/cli-runtime v0.20.1
	k8s.io/client-go v0.20.1
	rsc.io/letsencrypt v0.0.3 // indirect
	sigs.k8s.io/cli-utils v0.20.4
	sigs.k8s.io/controller-runtime v0.6.2
	sigs.k8s.io/yaml v1.2.0
)

replace helm.sh/helm/v3 v3.5.0 => github.com/hunchback/helm/v3 v3.5.0-hunchback

replace github.com/onsi/gomega => github.com/onsi/gomega v1.10.0

replace github.com/google/addlicense => github.com/the-mesh-for-data/addlicense v0.0.0-20200913135744-636c44b42906
