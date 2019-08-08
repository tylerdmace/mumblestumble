package wallet

import (
	"github.com/tylerdmace/mumblestumble/ecc"
	"github.com/tylerdmace/mumblestumble/txn"
)

type Wallet struct {
	keyring []ecc.PrivateKey
}

func (w *Wallet) CreateTransaction() txn.Txn {
	panic("Not Implemented")
}

func (w *Wallet) CompleteTransaction() txn.Txn {
	panic("Not Implemented")
}

func (w *Wallet) SubmitTransaction() (txn.Txn, error) {
	panic("Not Implemented")
}

func (w *Wallet) SetWalletPassphrase(newPassphrase string) error {
	panic("Not Implemented")
}

func (w *Wallet) ChangeWalletPassphrase(currentPassphrase string, newPassphrase string) error {
	panic("Not Implemented")
}
