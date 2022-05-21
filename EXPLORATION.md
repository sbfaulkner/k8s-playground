1. brew install podman
    * podman machine init
    * podman machine start
    * podman info
    * podman machine ssh
        * ps, etc.
        * exit
    * podman search ruby
    * podman search httpd --filter=is-official
    * podman pull docker.io/library/httpd
    * podman images
    * podman run -dt -p 8080:80/tcp docker.io/library/httpd
    * podman ps
    * curl http://localhost:8080
    * podman inspect <container-name> | grep IPAddress
    * podman logs <container-name>
    * podman top <container-name>
    * podman machine ssh
        * podman inspect -l
        * podman logs -l
        * podman top -l
    * podman stop <container-name>
    * podman ps -a
    * podman rm <container-name>
    * podman machine stop
    * podman machine rm
2. brew install kubectl
3. brew install kubectx
4. brew install kind
    * podman machine init --now
    * kind create cluster
    * kind create cluster --name kind-2
    * kubectl config get-contexts
    * kctx kind-kind
    * kubectl get all -A
    * kubectl get pods -A
    * kubectl cluster-info --context kind-kind-2
    * kind delete cluster --name kind-2
    * cat > kind-config.yaml <<EOF
        # three node (two workers) cluster config
        kind: Cluster
        apiVersion: kind.x-k8s.io/v1alpha4
        nodes:
        - role: control-plane
        - role: worker
        - role: worker
        EOF
    * kind create cluster --name k8s-exploration --config kind-config.yaml
    * kubectl get nodes
    * kubectl debug node/k8s-exploration-worker2 -it --image=busybox
    * kubectl get pods
    * kubectl delete pod <pod-name>
    * podman ps
    * podman exec -it k8s-exploration-worker2 /bin/sh
    * kind delete cluster
    * kind delete cluster --name k8s-exploration
