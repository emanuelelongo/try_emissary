kubectl create namespace emissary && \
kubectl apply -f emissary/emissary-crds.yaml && \
kubectl wait --timeout=90s --for=condition=available deployment emissary-apiext -n emissary-system

kubectl apply -f emissary/config.yaml

kubectl apply -f emissary/emissary-emissaryns.yaml && \
kubectl -n emissary wait --for condition=available --timeout=90s deploy -lproduct=aes

kubectl apply -f emissary/listener.yaml