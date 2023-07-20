package evmos

import (
	"fmt"
	"github.com/okx/go-wallet-sdk/coins/cosmos"
	"testing"
	"time"
)

/*
https://rest.bd.evmos.org:1317/cosmos/auth/v1beta1/accounts/evmos1rvs5xph4l3px2efynqsthus8p6r4exyrue82uy
curl -X POST -d '{"tx_bytes":"CswBCskBCikvaWJjLmFwcGxpY2F0aW9ucy50cmFuc2Zlci52MS5Nc2dUcmFuc2ZlchKbAQoIdHJhbnNmZXISCWNoYW5uZWwtMxobCgZhZXZtb3MSETEwMDAwMDAwMDAwMDAwMDAwIixldm1vczF5YzRxNnN2c2w5eHk5ZzJncGxnbmxweHdobnpyM3k3M3dmczB4aCotY29zbW9zMXJ2czV4cGg0bDNweDJlZnlucXN0aHVzOHA2cjRleHlyN2NreXh2MgA4gOCBuJnk3oEXEn0KWQpPCigvZXRoZXJtaW50LmNyeXB0by52MS5ldGhzZWNwMjU2azEuUHViS2V5EiMKIQOcJMA96W11QpNEacdGblBLXYYIw5nd27SBSxlh+Pc6UxIECgIIARgBEiAKGgoGYWV2bW9zEhA0MDAwMDAwMDAwMDAwMDAwEMCaDBpBxW58piSUv3r+MRmwIe3xllBkJxgF5I0QIlHLjym8amohsZWmyYCUzaux/pO2RNbB4K9VmJ2m8Y3/56w6Gpeh5wE=","mode":"BROADCAST_MODE_SYNC"}' https://rest.bd.evmos.org:1317/cosmos/tx/v1beta1/txs
*/
func TestTransfer(t *testing.T) {
	privateKeyHex := "//todo please replace your hex cosmos key"
	address, err := NewAddress(privateKeyHex)
	if err != nil {
		t.Fatal(err)
	}
	// evmos1yc4q6svsl9xy9g2gplgnlpxwhnzr3y73wfs0xh
	fmt.Println(address)

	param := cosmos.TransferParam{}
	param.FromAddress = "evmos1yc4q6svsl9xy9g2gplgnlpxwhnzr3y73wfs0xh"
	param.ToAddress = "evmos1yc4q6svsl9xy9g2gplgnlpxwhnzr3y73wfs0xh"
	param.Demon = "aevmos"
	param.Amount = "10000000000000000" // 18
	param.CommonParam.ChainId = "evmos_9001-2"
	param.CommonParam.Sequence = 0
	param.CommonParam.AccountNumber = 2091572
	param.CommonParam.FeeDemon = "aevmos"
	param.CommonParam.FeeAmount = "3500000000000000"
	param.CommonParam.GasLimit = 140000
	param.CommonParam.Memo = ""
	param.CommonParam.TimeoutHeight = 0
	tt, _ := cosmos.TransferAction(param, privateKeyHex, true)
	t.Log(tt)
}

func TestIbcTransfer(t *testing.T) {
	privateKeyHex := "//todo please replace your hex cosmos key"
	p := cosmos.IbcTransferParam{}
	p.CommonParam.ChainId = "evmos_9001-2"
	p.CommonParam.Sequence = 1
	p.CommonParam.AccountNumber = 2091572
	p.CommonParam.FeeDemon = "aevmos"
	p.CommonParam.FeeAmount = "4000000000000000"
	p.CommonParam.GasLimit = 200000
	p.CommonParam.Memo = ""
	p.CommonParam.TimeoutHeight = 0
	p.FromAddress = "evmos1yc4q6svsl9xy9g2gplgnlpxwhnzr3y73wfs0xh"
	p.ToAddress = "cosmos1rvs5xph4l3px2efynqsthus8p6r4exyr7ckyxv"
	p.Demon = "aevmos"
	p.Amount = "10000000000000000"
	p.SourcePort = "transfer"
	p.SourceChannel = "channel-3"
	p.TimeOutInSeconds = uint64(time.Now().UnixMilli()/1000) + 300
	tt, _ := cosmos.IbcTransferAction(p, privateKeyHex, true)
	t.Log(tt)
}
