package arp

import (
	"testing"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/aaletov/k8s-wol-operator/pkg/cmd"
	"strings"
	"github.com/stretchr/testify/assert"
)

func TestGetNetworkNodes(t *testing.T) {
	logger, hook := test.NewNullLogger()
	builder := strings.Builder{}
	builder.WriteString("172.18.0.1      02:42:a9:14:d0:da       (Unknown: locally administered)\n")
	builder.WriteString("172.18.0.2      02:42:ac:12:00:02       (Unknown: locally administered)\n")
	builder.WriteString("172.18.0.3      02:42:ac:12:00:03       (Unknown: locally administered)\n")
	builder.WriteString("172.18.0.5      02:42:ac:12:00:05       (Unknown: locally administered)")
	macs := []string{
		"02:42:a9:14:d0:da",
		"02:42:ac:12:00:02",
		"02:42:ac:12:00:03",
		"02:42:ac:12:00:05",
	}
	executor := cmd.MockExecutor{
		[]byte(builder.String()),
		nil,
	}
	arp := ARP{executor}
	nodes, err := arp.GetNetworkNodes(logger, "br-20b650c4ffa5", "172.18.0.0/24")

	if err != nil {
		t.Fatal(err)
	}

	for i, node := range nodes {
		assert.Equal(t, macs[i], node.MAC)
	}

	hook.Reset()
}