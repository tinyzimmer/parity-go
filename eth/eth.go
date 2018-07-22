package eth

import (
	"encoding/json"

	parity "github.com/tinyzimmer/parity-go"
)

type Client struct {
	Node parity.ParityNode
}

func NewClient(node parity.ParityNode) (c Client) {
	c.Node = node
	return
}

type EthSyncingInput []string

type EthSyncingOutput struct {
	Syncing       bool
	StartingBlock string `json:"startingBlock"`
	CurrentBlock  string `json:"currentBlock"`
	HighestBlock  string `json:"highestBlock"`
}

func (c Client) Syncing() (response *EthSyncingOutput, err error) {
	var raw EthSyncingOutput
	resp, suc, err := c.Node.Post("eth_syncing", EthSyncingInput{})
	if !suc {
		raw.Syncing = false
	} else {
		json.Unmarshal(resp.Result, &raw)
		raw.Syncing = true
	}
	response = &raw
	return
}
