package main

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func main() {
	proofAbi()
	fmt.Println()
	boolAbi()
}

func proofAbi() {
	var stateRootHash []byte
	var proof []byte
	var key []byte
	var value []byte

	Bytes, _ := abi.NewType("bytes", "", nil)

	var arguments = abi.Arguments{
		abi.Argument{Name: "stateRootHash", Type: Bytes, Indexed: false},
		abi.Argument{Name: "proof", Type: Bytes, Indexed: false},
		abi.Argument{Name: "key", Type: Bytes, Indexed: false},
		abi.Argument{Name: "value", Type: Bytes, Indexed: false},
	}

	stateRootHash, proof, key, value = sample1()

	encoded, err := arguments.Pack(stateRootHash, proof, key, value)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(hex.EncodeToString(encoded))

	// decoded := map[string]any{
	// 	"stateRootHash": nil,
	// 	"proof":         nil,
	// 	"key":           nil,
	// 	"value":         nil,
	// }
	// err = arguments.UnpackIntoMap(decoded, encoded)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(decoded)
}

func boolAbi() {
	var result bool

	Bool, _ := abi.NewType("bool", "", nil)

	var arguments = abi.Arguments{
		abi.Argument{Name: "proofResult", Type: Bool, Indexed: false},
	}

	result = false

	encoded, err := arguments.Pack(result)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(hex.EncodeToString(encoded))
}
