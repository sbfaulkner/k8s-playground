1. brew install podman
2. brew install kubectl
3. brew install kubectx
4. brew install kind
5. podman machine init --rootful
6. podman machine start
7. kind create cluster --config cluster.yaml
8. kubectl apply -f ingress-nginx.yaml
9. kubectl apply -f hello.yaml
