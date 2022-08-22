# k8s-wol-operator

Kubernetes Wake-On-LAN implementation

## Steps done

### Initializing project
The project was initialized with Kubebuilder

```
KUBEBUILDER_VERSION=latest
curl -L -o kubebuilder https://go.kubebuilder.io/dl/${KUBEBUILDER_VERSION}/$(go env GOOS)/$(go env GOARCH)
chmod +x kubebuilder && sudo mv kubebuilder /usr/local/bin/

go mod init github.com/aaletov/k8s-wol-operator
kubebuilder init --domain github.com --repo github.com/aaletov/k8s-wol-operator
```

API for custom resources was generated. Firstly, specs were generated from proto file, but due to some
issues with deepcopying types generated with `protoc-gen-go`, they were replaced by go types.

```
kubebuilder create api --group k8s-wol --version v1 --kind AttainableNode
kubebuilder create api --group k8s-wol --version v1 --kind WakeNodeUpRequest
```

### Creating basic design
User should create NodeWakeUpRequest object, which would be reconciled by controller within operator.
AttainableNode controller should discover nodes in his local network with `arp-scan` utility and create
corresponding objects in cluster. Operator should find AttainableNode for request using label/name and
invoke `WakeUp()` against AttainableNodeController, which should broadcast magic packet in his local network.  
