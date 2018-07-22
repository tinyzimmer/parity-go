package eth

import (
	"encoding/json"
)

type Client struct {
	Node ParityNode
}

type EthSyncingInput []string

type EthSyncingOutput struct {
	Syncing       bool
	StartingBlock string `json:"startingBlock"`
	CurrentBlock  string `json:"currentBlock"`
	HighestBlock  string `json:"highestBlock"`
}

func (e Eth) Syncing() (response *EthSyncingOutput, err error) {
	var raw EthSyncingOutput
	resp, suc, err := e.Post(e.Host, "eth_syncing", EthSyncingInput{})
	if !suc {
		raw.Syncing = false
	} else {
		json.Unmarshal(resp.Result, &raw)
		raw.Syncing = true
	}
	response = &raw
	return
}
