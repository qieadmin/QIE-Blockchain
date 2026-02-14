


package app

import (
	"github.com/puneet166/qie/app/eips"
	evmconfig "github.com/puneet166/qie/x/evm/config"
	"github.com/puneet166/qie/x/evm/core/vm"
)

// The init function of the config file allows to setup the global
// configuration for the EVM, modifying the custom ones defined in qie.
func init() {
	err := evmconfig.NewEVMConfigurator().
		WithExtendedEips(qieActivators).
		Configure()
	if err != nil {
		panic(err)
	}
}

// QieActivators defines a map of opcode modifiers associated
// with a key defining the corresponding EIP.
var qieActivators = map[string]func(*vm.JumpTable){
	"qie_0": eips.Enable0000,
	"qie_1": eips.Enable0001,
	"qie_2": eips.Enable0002,
}
