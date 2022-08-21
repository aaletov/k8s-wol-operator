# k8s-wol-operator

## Initializing project
Install kubebuilder and init project

```
KUBEBUILDER_VERSION=latest
curl -L -o kubebuilder https://go.kubebuilder.io/dl/${KUBEBUILDER_VERSION}/$(go env GOOS)/$(go env GOARCH)
chmod +x kubebuilder && sudo mv kubebuilder /usr/local/bin/

go mod init github.com/aaletov/k8s-wol-operator
kubebuilder init --domain github.com --repo github.com/aaletov/k8s-wol-operator
```

Create API

```
kubebuilder create api --group k8s-wol --version v1 --kind AttainableNode
kubebuilder create api --group k8s-wol --version v1 --kind WakeNodeUpRequest
```