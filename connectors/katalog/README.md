# Katalog

An inline data catalog and credentials manager powered by Kubernetes resources:
- [`Asset`](docs/README.md#asset)
- Kubernetes `Secret` wrapping a [`Credentials`](docs/README.md#credentials)

## Usage

Credenditals are stored in Kubernetes `Secret` resources. These `Secret` resources must be applied in the same namespace as `Asset` resources that reference them. Currently, the content of the `Secret` is restricted to include an embedded `Credentials` CRD. For example, to validate and apply a credentials YAML file:

```bash
kubectl apply --dry-run=client -f ${filepath} && kubectl create secret generic ${secretName} --from-file=main=${filepath}
```

Assets are stored in `Asset` resources. An `Asset` CRD includes a reference to a credentials `Secret`, connection information, and other metadata such as columns and associated security tags. Apply it like any other Kubernetes resource.

## Manage users

To manage `Asset` resources a Kubernetes user must be granted the `katalog-editor` cluster role.

To view  `Asset` resources a Kubernetes user must be granted the `katalog-viewer` cluster role. 

As always, you can create `RoleBinding` to grant these permissions to assets in a specific namespace.

## Develop, Build and Deploy

Katalog itself is just [`install/rbac.yaml`](install/rbac.yaml) and [`install/crds.gen.yaml`](install/crds.gen.yaml) that you can apply to your cluster directly. These files are generated from the files in the [`manifests`](manifests) directory with `make generate`.

The rest of the code is the connector and most of it is mapping to the _current_ connectors API. Use `make build docker-build` to build the connector and `make deploy` to deploy it. Cleanup with `make clean` and `make undeploy`.

Use `make all` to build and deploy everything.
