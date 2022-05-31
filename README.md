1. kubectl apply -f ingress-nginx.yaml
2. kubectl create namespace hello
3. kubectl create --namespace hello secret tls hello-tls --cert=config/tls/hello.pem --key=config/tls/hello-key.pem
4. kubectl apply -f hello.yaml
5. kubectl create namespace datadog
6. kubectl create --namespace datadog secret generic datadog --from-env-file=config/secrets/datadog.env
7. kubectl apply -f datadog-agent.yaml
