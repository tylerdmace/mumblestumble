package txn

import "math/big"

type Output struct { // vH + rG
	v int64   // Value in stumbles
	r big.Int // Blinding factors
}

type PedersonCommitment struct {
	input  txn.Output
	output txn.Output
}

type Txn struct {
	ID         int64
	commitment PedersonCommitment
}
