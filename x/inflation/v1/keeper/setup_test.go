package keeper_test

import (
	"github.com/stretchr/testify/suite"

	"github.com/puneet166/qie/testutil/integration/qie/factory"
	"github.com/puneet166/qie/testutil/integration/qie/grpc"
	"github.com/puneet166/qie/testutil/integration/qie/keyring"
	"github.com/puneet166/qie/testutil/integration/qie/network"
	"github.com/puneet166/qie/x/inflation/v1/types"
)

var denomMint = types.DefaultInflationDenom

type KeeperTestSuite struct {
	suite.Suite

	network *network.UnitTestNetwork
	handler grpc.Handler
	keyring keyring.Keyring
	factory factory.TxFactory
}

func (suite *KeeperTestSuite) SetupTest() {
	keys := keyring.New(2)
	nw := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keys.GetAllAccAddrs()...),
	)
	gh := grpc.NewIntegrationHandler(nw)
	tf := factory.New(nw, gh)

	suite.network = nw
	suite.factory = tf
	suite.handler = gh
	suite.keyring = keys
}
