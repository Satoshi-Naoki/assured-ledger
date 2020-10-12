// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

// +build functest,cloud,!cloud_with_consensus

package functest

import (
	"os"
	"testing"

	"github.com/google/gops/agent"

	"github.com/insolar/assured-ledger/ledger-core/application/testutils/launchnet"
	"github.com/insolar/assured-ledger/ledger-core/instrumentation/inslogger/instestlogger"
)

func init() {
	// starts gops agent https://github.com/google/gops on default addr (127.0.0.1:0)
	if err := agent.Listen(agent.Options{}); err != nil {
		panic(err)
	}
}

var (
	numNodes int
)

func getNodesCount() (int, error) {
	return numNodes, nil
}

func TestMain(m *testing.M) {
	instestlogger.SetTestOutputWithStub()

	launchnet.SetCloudFileLogging(true)

	numVirtual, numLight, numHeavy := getTestNodesSetup()

	numNodes = numVirtual + numLight + numHeavy

	cloudRunner := launchnet.PrepareCloudRunner(
		launchnet.WithNumVirtual(numVirtual),
		launchnet.WithNumLightMaterials(numLight),
		launchnet.WithNumHeavyMaterials(numHeavy))

	os.Exit(cloudRunner.Run(
		func(apiAddresses []string) int {
			return m.Run()
		},
	))
}