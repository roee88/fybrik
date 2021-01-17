# Example

Create credentials:
```bash
kubectl create secret generic data-csv-creds --from-file=main=example/credentials.yaml 
```

Create asset:
```bash
kubectl apply -f asset.yaml
```
