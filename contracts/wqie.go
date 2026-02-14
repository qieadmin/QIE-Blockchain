


package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	evmtypes "github.com/puneet166/qie/x/evm/types"
)

var (
	//go:embed compiled_contracts/WQIE.json
	WQIEJSON []byte

	// WQIEContract is the compiled contract of WQIE
	WQIEContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(WQIEJSON, &WQIEContract)
	if err != nil {
		panic(err)
	}

	if len(WQIEContract.Bin) == 0 {
		panic("failed to load WQIE smart contract")
	}
}
