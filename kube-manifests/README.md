# How to apply the manifests

Run the following:
```
export IMAGE_URL="harbor.lqd.dk/liquid/hello-go"
export ISTIO_HTTP_URL="http-hello-go.istio-knet.lqd.dk"
export ISTIO_HTTPS_URL="https-hello-go.istio-knet.lqd.dk"
export NGINX_HOST="hello-go.knet.lqd.dk"

cat deployment.yaml | sed "s'{{IMAGE_URL}}'$IMAGE_URL'g" | kubectl apply -f -
cat istio-http.yaml | sed "s/{{ISTIO_HTTP_URL}}/$ISTIO_HTTP_URL/g" | kubectl apply -f -
cat istio-https.yaml | sed "s/{{ISTIO_HTTPS_URL}}/$ISTIO_HTTPS_URL/g" | kubectl apply -f -
cat nginx-https.yaml | sed "s/{{NGINX_HOST}}/$NGINX_HOST/g" | kubectl apply -f -
```

## Additional information

- The nginx relies on the certmanager running in the cluster.

- Istio uses Secure Gateways (https://istio.io/docs/tasks/traffic-management/ingress/secure-ingress-sds/?_ga=2.59346669.1523461124.1580805708-807816478.1567755843)
