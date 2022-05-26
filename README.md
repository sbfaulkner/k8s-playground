1. brew install podman
2. brew install kubectl
3. brew install kubectx
4. brew install kind
5. brew install mkcert
6. mkcert -install
7. podman machine init --rootful
8. podman machine start
9. kind create cluster --config cluster.yaml
10. kubectl apply -f ingress-nginx.yaml
11. mkcert -cert-file config/tls/hello.pem -key-file config/tls/hello-key.pem hello.k8s-playground.dev
12. kubectl create namespace hello
13. kubectl create --namespace hello secret tls hello-tls --cert=config/tls/hello.pem --key=config/tls/hello-key.pem
14. kubectl apply -f hello.yaml
15. kubectl create namespace datadog
16. kubectl create --namespace datadog secret generic datadog --from-env-file=config/secrets/datadog.env
17. kubectl apply -f datadog-agent.yaml
