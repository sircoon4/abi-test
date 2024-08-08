package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type sample func() (
	nonce *big.Int,
	from common.Address,
	to common.Address,
	amount *big.Int,
)

func samples() []sample {
	return []sample{
		sample1,
		sample2,
		sample3,
		sample4,
	}
}

func sample1() (nonce *big.Int, from common.Address, to common.Address, amount *big.Int) {
	nonce = big.NewInt(1)
	from = common.HexToAddress("0x1234567890123456789012345678901234567890")
	to = common.HexToAddress("0x1234567890123456789012345678901234567890")
	amount = big.NewInt(100)

	return
}

func sample2() (nonce *big.Int, from common.Address, to common.Address, amount *big.Int) {
	nonce = big.NewInt(10)
	from = common.HexToAddress("0x47E0Dd0B503C153D7FB78c43cc9aC135C60Dfd94")
	to = common.HexToAddress("0xCE70F2e49927D431234BFc8D439412eef3a6276b")
	amount = big.NewInt(1000)

	return
}

func sample3() (nonce *big.Int, from common.Address, to common.Address, amount *big.Int) {
	nonce = big.NewInt(9876543210)
	from = common.HexToAddress("0xCE70F2e49927D431234BFc8D439412eef3a6276b")
	to = common.HexToAddress("0x47E0Dd0B503C153D7FB78c43cc9aC135C60Dfd94")
	amount = big.NewInt(1234567890)

	return
}

func sample4() (nonce *big.Int, from common.Address, to common.Address, amount *big.Int) {
	nonce = big.NewInt(97)
	from = common.HexToAddress("0xCE70F2e49927D431234BFc8D439412eef3a6276b")
	to = common.HexToAddress("0xCE70F2e49927D431234BFc8D439412eef3a6276b")
	amount = big.NewInt(1234)

	return
}
