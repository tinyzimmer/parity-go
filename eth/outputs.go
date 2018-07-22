package eth

import (
  "github.com/tinyzimmer/parity-go"
)

type EthAccountsOutput []string

type EthBlockNumberOutput string

func (o EthBlockNumberOutput) Decode() uint64 {
	return parity.HexToInt(o)
}

type EthSyncingOutput struct {
	Syncing       bool
	StartingBlock string `json:"startingBlock"`
	CurrentBlock  string `json:"currentBlock"`
	HighestBlock  string `json:"highestBlock"`
}
