### Install Jaeger

```sh
kubectl apply -f jaeger.yaml
kubectl edit configmap istio -n istio-system
set tracing.sampling=100
```

### Deploy httpserver svc1,svc2,svc3

```sh
kubectl create ns model12
kubectl label ns model12 istio-injection=enabled
kubectl -n model12 apply -f service0.yaml
kubectl -n model12 apply -f service1.yaml
kubectl -n model12 apply -f service2.yaml

openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=cncamp Inc./CN=*.cncamp.io' -keyout cncamp.io.key -out cncamp.io.crt
kubectl create -n istio-system secret tls cncamp-credential --key=cncamp.io.key --cert=cncamp.io.crt

kubectl apply -f istio-specs.yaml -n model12
```

### Check ingress ip

```sh
k get svc -nistio-system

istio-ingressgateway   LoadBalancer   $INGRESS_IP
```

### Access the tracing via ingress for 100 times(sampling rate is 1%)

```sh
curl --resolve httpsserver.cncamp.io:443:$INGRESS_IP https://httpsserver.cncamp.io/service0 -v -k
```

### Check tracing dashboard

```sh
istioctl dashboard jaeger
```
