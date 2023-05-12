package BlockStructs

import (
	"crypto/sha256"
	"fmt"
)

func (tx *NFTTransaction) Hash() []byte {
	transactionData := fmt.Sprint(tx.NFTID) + tx.OwnerPubKey + tx.ReceiverPubKey + fmt.Sprintf("%f", tx.Amount)
	hash := sha256.Sum256([]byte(transactionData))
	return hash[:]
}
