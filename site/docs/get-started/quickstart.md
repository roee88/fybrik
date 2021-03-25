# Quick Start Guide

Follow this guide to install Mesh for Data using default parameters that are suitable for experimentation.
<!-- For a full installation refer to the [full installation guide](./setup/install) instead. -->

## Before you begin

Ensure that you have the following:

- [Helm](https://helm.sh/) 3.3 or newer must be installed and configured on your machine.
- [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) 1.16 or newer must be installed on your machine.
- Access to a Kubernetes cluster such as [Kind](http://kind.sigs.k8s.io/) as a cluster administrator.

## Install cert-manager

Mesh for Data requires [cert-manager](https://cert-manager.io) to be installed to your cluster. 
Many clusters already include cert-manager. Run the following to install cert-manager only if it's missing:

```bash
kubectl get namespace cert-manager || kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.2.0/cert-manager.yaml
``` 

## Install Hashicorp Vault and plugins

[Hashicorp Vault](https://www.vaultproject.io/) and a [secrets-kubernetes-reader](https://github.com/mesh-for-data/vault-plugin-secrets-kubernetes-reader) plugin are used by Mesh for Data for credential management.

Run the following to install vault and the plugin in development mode:

=== "Kubernetes" 

    ```bash
    helm repo add hashicorp https://helm.releases.hashicorp.com
    helm install vault hashicorp/vault --version 0.9.1 --create-namespace -n m4d-system \
        --set "server.dev.enabled=true" \
        --values https://raw.githubusercontent.com/IBM/the-mesh-for-data/master/third_party/vault/plugin-secrets-kubernetes-reader/values.yaml \
        --wait --timeout 120s
    ```

=== "OpenShift"

    ```bash
    helm repo add hashicorp https://helm.releases.hashicorp.com
    helm install vault hashicorp/vault --version 0.9.1 --create-namespace -n m4d-system \
        --set "global.openshift=true" \
        --set "server.dev.enabled=true" \
        --values https://raw.githubusercontent.com/IBM/the-mesh-for-data/master/third_party/vault/plugin-secrets-kubernetes-reader/values.yaml \
        --wait --timeout 120s
    ```

## Install control plane

The control plane includes a `manager` service that connects to a data catalog and to a policy manager. 
Install the latest release of Mesh for Data with a built-in data catalog and with [Open Policy Agent](https://www.openpolicyagent.org) as the policy manager:

```bash
helm repo add m4d-charts https://mesh-for-data.github.io/charts
helm install m4d-crd m4d-charts/m4d-crd -n m4d-system --wait
helm install m4d m4d-charts/m4d -n m4d-system --wait
```


## Install modules

[Modules](../concepts/modules.md) are plugins that the control plane deploys whenever required. 

Install the [arrow flight module](https://github.com/ibm/the-mesh-for-data-flight-module):

```bash
kubectl apply -f https://raw.githubusercontent.com/IBM/the-mesh-for-data-flight-module/master/module.yaml
```