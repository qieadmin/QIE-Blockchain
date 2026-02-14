


package ante_test

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/puneet166/qie/testutil/integration/qie/network"
	evmante "github.com/puneet166/qie/x/evm/ante"
)

func (suite *EvmAnteTestSuite) TestBuildEvmExecutionCtx() {
	network := network.New()

	ctx := evmante.BuildEvmExecutionCtx(network.GetContext())

	suite.Equal(storetypes.GasConfig{}, ctx.KVGasConfig())
	suite.Equal(storetypes.GasConfig{}, ctx.TransientKVGasConfig())
}
