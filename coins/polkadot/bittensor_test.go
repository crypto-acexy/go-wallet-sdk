package polkadot

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"testing"
)

var prikey = "45d3bd794c5bc6ed91ae41c93c0baed679935703dfac72c48d27f8321b8d3a40"

func TestBittensorTransfer(t *testing.T) {
	tx := TxStruct{
		From:         "5CLgTrXNgtS24q8TUKiLp9G3gMyuXz8MEUyky4XCfnHwALag",
		To:           "5CLgTrXNgtS24q8TUKiLp9G3gMyuXz8MEUyky4XCfnHwALag",
		Amount:       1,
		Nonce:        2,
		Tip:          0,
		BlockHeight:  1824833,
		BlockHash:    "0xd83a892d64ac2587fb894a7687fe600d51d217295ade6f637b82ed7707cb9478",
		GenesisHash:  "0x2f0555cc76fc2840a25a6ea3b9637146806f1f44b090c175ffde2a7e5ab36c03",
		SpecVersion:  136,
		TxVersion:    1,
		ModuleMethod: "0500",
		Version:      "84",
	}
	signed, _ := SignTx(tx, Transfer, prikey)
	fmt.Println(signed)
}

func TestCreateAddress(t *testing.T) {
	priKey, _ := hex.DecodeString(prikey)
	p := ed25519.NewKeyFromSeed(priKey)
	publicKey := p.Public().(ed25519.PublicKey)
	address, _ := PubKeyToAddress(publicKey, SubstratePrefix)
	fmt.Println(address)
}
