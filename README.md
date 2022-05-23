1. brew install podman
2. brew install kubectl
3. brew install kubectx
4. brew install kind
5. podman machine init --now
6. kind create cluster --config cluster.yaml --wait 60s
7. kubectl apply -f hello-node/

kubectl config view
kubectl get deployments
kubectl get pods
kubectl get events
kubectl get services

??? kubectl expose deployment hello-node --type=LoadBalancer --port=8080
