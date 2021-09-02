package irb

import (
	"github.com/hyperledger-labs/fabric-smart-client/integration/nwo/api"
	"github.com/hyperledger-labs/fabric-smart-client/integration/nwo/fabric"
	"github.com/hyperledger-labs/fabric-smart-client/integration/nwo/fsc"
)

func Topology() []api.Topology {
	fabricTopology := fabric.NewDefaultTopology()
	fabricTopology.AddOrganizationsByName("Org1", "Org2", "Org3")

	fscTopology := fsc.NewTopology()

	// data provider
	provider := fscTopology.AddNodeByName("provider")
	provider.AddOptions(fabric.WithOrganization("Org1"))

	// investigator
	investigator := fscTopology.AddNodeByName("investigator")
	investigator.AddOptions(fabric.WithOrganization("Org2"))

	// experimenter
	experimenter := fscTopology.AddNodeByName("experimenter")
	experimenter.AddOptions(fabric.WithOrganization("Org3"))

	return []api.Topology{fabricTopology, fscTopology}
}
