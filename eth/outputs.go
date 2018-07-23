package eth

import (
	"github.com/tinyzimmer/parity-go"
)

type AccountsOutput []string

type BlockNumberOutput string

func (o BlockNumberOutput) Decode() uint64 {
	return parity.HexToInt(o)
}

type SyncingOutput struct {
	Syncing             bool
	StartingBlock       string `json:"startingBlock"`
	CurrentBlock        string `json:"currentBlock"`
	HighestBlock        string `json:"highestBlock"`
	WarpChunksAmount    string `json:"warpChunksAmount"`
	WarpChunksProcessed string `json:"warpChunksProcessed"`
}

type SyncingOutputDecoded struct {
	Syncing             bool
	StartingBlock       uint64
	CurrentBlock        uint64
	HighestBlock        uint64
	WarpChunksAmount    uint64
	WarpChunksProcessed uint64
}

func (o SyncingOutput) DecodeAll() SyncingOutputDecoded {
	return SyncingOutputDecoded{
		Syncing:             o.Syncing,
		StartingBlock:       parity.HexToInt(o.StartingBlock),
		CurrentBlock:        parity.HexToInt(o.CurrentBlock),
		HighestBlock:        parity.HexToInt(o.HighestBlock),
		WarpChunksAmount:    parity.HexToInt(o.WarpChunksAmount),
		WarpChunksProcessed: parity.HexToInt(o.WarpChunksProcessed),
	}
}
