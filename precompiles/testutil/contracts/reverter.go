


package contracts

import (
	contractutils "github.com/puneet166/qie/contracts/utils"
	evmtypes "github.com/puneet166/qie/x/evm/types"
)

func LoadReverterContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("Reverter.json")
}
