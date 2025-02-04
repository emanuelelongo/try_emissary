# Experimenting Ambassador Emissary

Simple experiments about routing http traffic to different containers with Emissary

## Prepare the cluster

### Install Emissary

From the [getting started](https://www.getambassador.io/docs/emissary/latest/tutorials/getting-started):

```bash
# Add the Repo:
helm repo add datawire https://app.getambassador.io
helm repo update
 
# Create Namespace and Install:
kubectl create namespace emissary && \
kubectl apply -f https://app.getambassador.io/yaml/emissary/3.9.1/emissary-crds.yaml
 
kubectl wait --timeout=90s --for=condition=available deployment emissary-apiext -n emissary-system
 
helm install emissary-ingress --namespace emissary datawire/emissary-ingress && \
kubectl -n emissary wait --for condition=available --timeout=90s deploy -lapp.kubernetes.io/instance=emissary-ingress
```

### Create the `listener`

```bash
kubectl apply -f listener.yaml
```

### Port-Forward the `emissary-ingress` service\

The `emissary-ingress` service is created by default with the LoadBalancer type. Since we are running the examples on a local cluster let's expose it by port-forwarding its port.

```bash
kubectl port-forward -n emissary svc/emissary-ingress 8080:80
```

## Examples

### Routing by path

In this example we create two μ-services and let emissary forward traffic to them depending on the requested path.

```bash
helm install example path-routing-example

curl localhost:8080/sushi/
🍣
curl localhost:8080/sushi/bar
🍣
curl localhost:8080/pizza/
🍕
curl localhost:8080/foo/
🍝

helm delete example
```

### Routing by host

In this example we create two μ-services and let emissary forward traffic to them depending on the request host.

```bash
helm install example host-routing-example

curl localhost:8080/ -H host:pizza.lcl
🍕
curl localhost:8080/foo -H host:pizza.lcl
🍕
curl localhost:8080/bar -H host:sushi.lcl
🍣
curl localhost:8080/foo/
🍝

helm delete example
```

### Mixing routes by a host and path

In this example we create two μ-services and let emissary forward traffic by matching the request host with a given list.

```bash
helm install example mixed-routing-example

curl localhost:8080/ -H host:pizza.lcl 
🍕
curl localhost:8080/sushi/bar
🍣
curl localhost:8080/foo
🍝

helm delete example
```

### Routing by a host list

In this example we create two μ-services and let emissary forward traffic by matching the request host with a given list.

```bash
helm install example host-list-example

curl localhost:8080/foo/bar -H host:host1.lcl
🍕
curl localhost:8080/foo/bar -H host:host2.lcl
🍣
curl localhost:8080/foo/bar -H host:host3.lcl
🍕
curl localhost:8080/foo/bar -H host:host4.lcl
🍣

helm delete example
```

## Warp Benchmark

```bash
helm install benchmark benchmark
```

This will expose create a Min.io instance and 5 different mappings

- `minio-1.lcl` 
- `minio-2.lcl`
- `minio-3.lcl`
- `minio-4.lcl`
- `*`


```bash
warp mixed --host=minio-{0...4}.lcl:8080 \
    --access-key=minioadmin \
    --secret-key=minioadmin \
    --objects=500 \
    --obj.size=100MiB \
    --concurrent=30 \
    --duration=60s

helm delete benchmark

```
