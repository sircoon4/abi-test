package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sircoon4/bencodex-go"
	"github.com/sircoon4/bencodex-go/bencodextype"
)

func main() {
	input, err := encodeWithdrawalTransactionInput(samples()[3]())
	if err != nil {
		fmt.Println(err)
		return
	}

	withdrawalTransaction := map[string]any{
		"nonce":  nil, // *big.Int
		"from":   nil, // common.Address
		"to":     nil, // common.Address
		"amount": nil, // *big.Int
	}
	withdrawalTransaction, err = parseWithdrawalTransactionInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	nonce := withdrawalTransaction["nonce"].(*big.Int)
	from := withdrawalTransaction["from"].(common.Address).Bytes()
	to := withdrawalTransaction["to"].(common.Address).Bytes()
	amount := withdrawalTransaction["amount"].(*big.Int)

	dict := bencodextype.NewDictionary()
	dict.Set("nonce", nonce)
	dict.Set("from", from)
	dict.Set("to", to)
	dict.Set("amount", amount)

	encoded, err := bencodex.Encode(dict)
	if err != nil {
		fmt.Println(err)
		return
	}

	sum := sha1.Sum(encoded)

	res := common.CopyBytes(addressAbi(sum))

	fmt.Println()
	fmt.Println(hex.EncodeToString(res))
}
