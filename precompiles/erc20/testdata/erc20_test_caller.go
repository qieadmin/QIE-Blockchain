


package testdata

import (
	contractutils "github.com/puneet166/qie/contracts/utils"
	evmtypes "github.com/puneet166/qie/x/evm/types"
)

func LoadERC20TestCaller() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("ERC20TestCaller.json")
}
