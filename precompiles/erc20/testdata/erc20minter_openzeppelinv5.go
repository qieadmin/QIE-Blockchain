


package testdata

import (
	contractutils "github.com/puneet166/qie/contracts/utils"
	evmtypes "github.com/puneet166/qie/x/evm/types"
)

func LoadERC20MinterV5Contract() (evmtypes.CompiledContract, error) {
	return contractutils.LegacyLoadContractFromJSONFile("ERC20Minter_OpenZeppelinV5.json")
}
