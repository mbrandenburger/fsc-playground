package irb

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hyperledger-labs/fabric-smart-client/integration"
	"github.com/hyperledger-labs/fabric-smart-client/integration/fabric/atsa/fsc"
)

func StartPort() int {
	return integration.AssetTransferSecuredAgreementWithApprovers.StartPortForNode()
}

func TestFlow(t *testing.T) {
	ii, err := integration.Generate(StartPort(), fsc.Topology()...)
	assert.NoError(t, err)
	ii.Start()
	defer ii.Stop()

}
