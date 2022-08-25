package arp

import (
	"fmt"
	"strings" //

	"github.com/aaletov/k8s-wol-operator/pkg/cmd"
	"github.com/sirupsen/logrus"
)

const (
	// Device and subnet
	ArpScanCmd = "arp-scan --interface=%s --plain %s"
)

type ARP struct {
	e cmd.Executor
}

func NewARP(logger *logrus.Logger) ARP {
	e := cmd.GoExecutor{logger}
	return ARP{e}
}

type NetworkNode struct {
	MAC string
}

func (a ARP) GetNetworkNodes(logger *logrus.Logger, device, subnet string) ([]NetworkNode, error) {
	out, err := a.e.Run(fmt.Sprintf(ArpScanCmd, device, subnet))

	if err != nil {
		logger.Errorf("Error while running commands: %s", string(out))
		return nil, err
	}

	logger.Info(string(out))

	lines := strings.Split(string(out), "\n")
	nodes := make([]NetworkNode, 0)

	for _, line := range lines {
		mac := strings.Fields(line)[1]
		nodes = append(nodes, NetworkNode{mac})
	}

	return nodes, nil
}
