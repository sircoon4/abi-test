package vm

import (
	"github.com/sircoon4/abi-test/vm/libplanet"
)

func Run(input []byte) ([]byte, error) {
	action, err := libplanet.ExtractActionDictFromSerializedPayload(input)
	if err != nil {
		return nil, err
	}

	return libplanet.ExtractActionEthAbi(action)
}
