


package testdata

import (
	contractutils "github.com/puneet166/qie/contracts/utils"
	evmtypes "github.com/puneet166/qie/x/evm/types"
)

func LoadBankCallerContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("BankCaller.json")
}
