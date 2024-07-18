package main

import "encoding/hex"

// stateRootHash []byte => sha256(bencoded)
// proof [][]byte => bencoded list
// key []byte => []byte
// value []byte => bencoded

func sample00() (stateRootHash []byte, proof [][]byte, key []byte, value []byte) {
	// FullTrieHash
	stateRootHash, _ = hex.DecodeString("979a00921d42d2ca63e98c1c2ac07f0eacbb99e363b8f2f7f8e4d19c854b6c20")

	// ProofRoot
	proofR, _ := hex.DecodeString("6c313a0033323a84cac50effba60ce08415ff3356839fa032a91658a7740d60ffba4af6245a0c565")
	// ProofNode0
	proof0, _ := hex.DecodeString("6c33323ab7ba99008cd22b95a8fce03c7d2d72ca2352732ad0a849b91653b12ab07cce4c6c6e75323a3031656e6e6e6e6e6e6e6e6e6e6e6e6e6e6e65")
	// ProofNode00
	proof00, _ := hex.DecodeString("6c6c313a006c6e75343a30303030656533323a3edb1a2151507f2c3853e1b3a8039ec5768b22e71c268f49d12addbb6806b3f16e6e6e6e6e6e6e6e6e6e6e6e6e6e6c6e75323a30306565")
	proof = [][]byte{
		proofR,
		proof0,
		proof00,
	}

	key, _ = hex.DecodeString("00")

	// V00
	value, _ = hex.DecodeString("75323a3030")

	return
}

func sample01() (stateRootHash []byte, proof [][]byte, key []byte, value []byte) {
	// FullTrieHash
	stateRootHash, _ = hex.DecodeString("979a00921d42d2ca63e98c1c2ac07f0eacbb99e363b8f2f7f8e4d19c854b6c20")

	// ProofRoot
	proofR, _ := hex.DecodeString("6c313a0033323a84cac50effba60ce08415ff3356839fa032a91658a7740d60ffba4af6245a0c565")
	// ProofNode0
	proof0, _ := hex.DecodeString("6c33323ab7ba99008cd22b95a8fce03c7d2d72ca2352732ad0a849b91653b12ab07cce4c6c6e75323a3031656e6e6e6e6e6e6e6e6e6e6e6e6e6e6e65")
	proof = [][]byte{
		proofR,
		proof0,
	}

	key, _ = hex.DecodeString("01")

	// V01
	value, _ = hex.DecodeString("75323a3031")

	return
}

func sample0000() (stateRootHash []byte, proof [][]byte, key []byte, value []byte) {
	// FullTrieHash
	stateRootHash, _ = hex.DecodeString("979a00921d42d2ca63e98c1c2ac07f0eacbb99e363b8f2f7f8e4d19c854b6c20")

	// ProofRoot
	proofR, _ := hex.DecodeString("6c313a0033323a84cac50effba60ce08415ff3356839fa032a91658a7740d60ffba4af6245a0c565")
	// ProofNode0
	proof0, _ := hex.DecodeString("6c33323ab7ba99008cd22b95a8fce03c7d2d72ca2352732ad0a849b91653b12ab07cce4c6c6e75323a3031656e6e6e6e6e6e6e6e6e6e6e6e6e6e6e65")
	// ProofNode00
	proof00, _ := hex.DecodeString("6c6c313a006c6e75343a30303030656533323a3edb1a2151507f2c3853e1b3a8039ec5768b22e71c268f49d12addbb6806b3f16e6e6e6e6e6e6e6e6e6e6e6e6e6e6c6e75323a30306565")
	proof = [][]byte{
		proofR,
		proof0,
		proof00,
	}

	key, _ = hex.DecodeString("0000")

	// V0000
	value, _ = hex.DecodeString("75343a30303030")

	return
}

func sample0010() (stateRootHash []byte, proof [][]byte, key []byte, value []byte) {
	// FullTrieHash
	stateRootHash, _ = hex.DecodeString("979a00921d42d2ca63e98c1c2ac07f0eacbb99e363b8f2f7f8e4d19c854b6c20")

	// ProofRoot
	proofR, _ := hex.DecodeString("6c313a0033323a84cac50effba60ce08415ff3356839fa032a91658a7740d60ffba4af6245a0c565")
	// ProofNode0
	proof0, _ := hex.DecodeString("6c33323ab7ba99008cd22b95a8fce03c7d2d72ca2352732ad0a849b91653b12ab07cce4c6c6e75323a3031656e6e6e6e6e6e6e6e6e6e6e6e6e6e6e65")
	// ProofNode00
	proof00, _ := hex.DecodeString("6c6c313a006c6e75343a30303030656533323a3edb1a2151507f2c3853e1b3a8039ec5768b22e71c268f49d12addbb6806b3f16e6e6e6e6e6e6e6e6e6e6e6e6e6e6c6e75323a30306565")
	// ProofNode001
	proof001, _ := hex.DecodeString("6c313a0033323ab16b9db1bce3fedf7dd114b02938f2476c21936a3f30fffecc8db14d377d3dd965")
	// ProofNode0010
	proof0010, _ := hex.DecodeString("6c6e7533323a303030303030303030303030303030303030303030303030303030303030313065")
	proof = [][]byte{
		proofR,
		proof0,
		proof00,
		proof001,
		proof0010,
	}

	key, _ = hex.DecodeString("0010")

	// V0010
	value, _ = hex.DecodeString("7533323a3030303030303030303030303030303030303030303030303030303030303130")

	return
}

// Wrong hash
func sampleFalseHash() (stateRootHash []byte, proof [][]byte, key []byte, value []byte) {
	// HalfTrieHash
	stateRootHash, _ = hex.DecodeString("6cc5c2ca1b7b146268f0d930c58c7e5441b807e72cf16d56f52c869a594b17bf")

	// ProofRoot
	proofR, _ := hex.DecodeString("6c313a0033323a84cac50effba60ce08415ff3356839fa032a91658a7740d60ffba4af6245a0c565")
	// ProofNode0
	proof0, _ := hex.DecodeString("6c33323ab7ba99008cd22b95a8fce03c7d2d72ca2352732ad0a849b91653b12ab07cce4c6c6e75323a3031656e6e6e6e6e6e6e6e6e6e6e6e6e6e6e65")
	// ProofNode00
	proof00, _ := hex.DecodeString("6c6c313a006c6e75343a30303030656533323a3edb1a2151507f2c3853e1b3a8039ec5768b22e71c268f49d12addbb6806b3f16e6e6e6e6e6e6e6e6e6e6e6e6e6e6c6e75323a30306565")
	proof = [][]byte{
		proofR,
		proof0,
		proof00,
	}

	key, _ = hex.DecodeString("00")

	// V00
	value, _ = hex.DecodeString("75323a3030")

	return
}

// Wrong value
func sampleFalseValue() (stateRootHash []byte, proof [][]byte, key []byte, value []byte) {
	// FullTrieHash
	stateRootHash, _ = hex.DecodeString("979a00921d42d2ca63e98c1c2ac07f0eacbb99e363b8f2f7f8e4d19c854b6c20")

	// ProofRoot
	proofR, _ := hex.DecodeString("6c313a0033323a84cac50effba60ce08415ff3356839fa032a91658a7740d60ffba4af6245a0c565")
	// ProofNode0
	proof0, _ := hex.DecodeString("6c33323ab7ba99008cd22b95a8fce03c7d2d72ca2352732ad0a849b91653b12ab07cce4c6c6e75323a3031656e6e6e6e6e6e6e6e6e6e6e6e6e6e6e65")
	// ProofNode00
	proof00, _ := hex.DecodeString("6c6c313a006c6e75343a30303030656533323a3edb1a2151507f2c3853e1b3a8039ec5768b22e71c268f49d12addbb6806b3f16e6e6e6e6e6e6e6e6e6e6e6e6e6e6c6e75323a30306565")
	proof = [][]byte{
		proofR,
		proof0,
		proof00,
	}

	key, _ = hex.DecodeString("00")

	// V01
	value, _ = hex.DecodeString("75323a3031")

	return
}

// Wrong key
func sampleFalseKey() (stateRootHash []byte, proof [][]byte, key []byte, value []byte) {
	// FullTrieHash
	stateRootHash, _ = hex.DecodeString("979a00921d42d2ca63e98c1c2ac07f0eacbb99e363b8f2f7f8e4d19c854b6c20")

	// ProofRoot
	proofR, _ := hex.DecodeString("6c313a0033323a84cac50effba60ce08415ff3356839fa032a91658a7740d60ffba4af6245a0c565")
	// ProofNode0
	proof0, _ := hex.DecodeString("6c33323ab7ba99008cd22b95a8fce03c7d2d72ca2352732ad0a849b91653b12ab07cce4c6c6e75323a3031656e6e6e6e6e6e6e6e6e6e6e6e6e6e6e65")
	// ProofNode00
	proof00, _ := hex.DecodeString("6c6c313a006c6e75343a30303030656533323a3edb1a2151507f2c3853e1b3a8039ec5768b22e71c268f49d12addbb6806b3f16e6e6e6e6e6e6e6e6e6e6e6e6e6e6c6e75323a30306565")
	proof = [][]byte{
		proofR,
		proof0,
		proof00,
	}

	key, _ = hex.DecodeString("01")

	// V00
	value, _ = hex.DecodeString("75323a3030")

	return
}

// Wrong proof
func sampleFalseProof() (stateRootHash []byte, proof [][]byte, key []byte, value []byte) {
	// FullTrieHash
	stateRootHash, _ = hex.DecodeString("979a00921d42d2ca63e98c1c2ac07f0eacbb99e363b8f2f7f8e4d19c854b6c20")

	// ProofRoot
	proofR, _ := hex.DecodeString("6c313a0033323a84cac50effba60ce08415ff3356839fa032a91658a7740d60ffba4af6245a0c565")
	// ProofNode0
	proof0, _ := hex.DecodeString("6c33323ab7ba99008cd22b95a8fce03c7d2d72ca2352732ad0a849b91653b12ab07cce4c6c6e75323a3031656e6e6e6e6e6e6e6e6e6e6e6e6e6e6e65")
	// ProofNode00
	proof00, _ := hex.DecodeString("6c6c313a006c6e75343a30303030656533323a3edb1a2151507f2c3853e1b3a8039ec5768b22e71c268f49d12addbb6806b3f16e6e6e6e6e6e6e6e6e6e6e6e6e6e6c6e75323a30306565")
	proof = [][]byte{
		proofR,
		proof0,
		proof00,
	}

	key, _ = hex.DecodeString("01")

	// V01
	value, _ = hex.DecodeString("75323a3031")

	return
}
