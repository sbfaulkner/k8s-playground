
## Deploy ingress controller

```
kubectl apply -f ingress-nginx.yaml
```

## Install Datadog Agent

### Create Namespace
```
kubectl create namespace datadog-agent
```

### Create ClusterRole
```
kubectl apply -f datadog-agent/clusterrole.yaml
```

### Create ServiceAccount
```
kubectl apply -f datadog-agent/serviceaccount.yaml
```

### Create ClusterRoleBinding
```
kubectl apply -f datadog-agent/clusterrolebinding.yaml
```

### Create API Key Secret
```
kubectl create secret generic datadog -n datadog-agent --from-env-file config/secrets/datadog.env
```

### Create install information ConfigMap
```
kubectl apply -f datadog-agent/install_info-configmap.yaml
```

### Create confd ConfigMap
```
kubectl apply -f datadog-agent/confd-configmap.yaml
```

### Deploy DaemonSet
```
kubectl apply -f datadog-agent/daemonset.yaml
```

## Build container image
```
dev deploy
```

## Install Application

### Create namespace
```
kubectl create namespace hello
```

### Create TLS certificate
```
kubectl create --namespace hello secret tls hello-tls --cert=config/tls/hello.pem --key=config/tls/hello-key.pem
```

### Deploy application
```
kubectl apply -f hello.yaml
```
