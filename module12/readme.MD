### Install Jaeger

```sh
kubectl apply -f jaeger.yaml
kubectl edit configmap istio -n istio-system
set tracing.sampling=100
```

### Deploy httpserver

```sh
kubectl create ns module12
kubectl label ns module12 istio-injection=enabled
kubectl create -f httpserver.yaml -n module12

openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=cncamp Inc./CN=*.cncamp.io' -keyout cncamp.io.key -out cncamp.io.crt
kubectl create -n istio-system secret tls cncamp-credential --key=cncamp.io.key --cert=cncamp.io.crt

kubectl apply -f istio-specs.yaml -n module12
```

### Check ingress ip

```sh
k get svc -nistio-system

istio-ingressgateway   LoadBalancer   $INGRESS_IP
```

### Access the tracing via ingress for 100 times(sampling rate is 1%)

```sh
curl --resolve httpsserver.cncamp.io:443:$INGRESS_IP https://httpsserver.cncamp.io/healthz -v -k
```

### Check tracing dashboard

```sh
istioctl dashboard jaeger
```
