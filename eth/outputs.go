package eth

import (
	"github.com/tinyzimmer/parity-go"
)

type AccountsOutput []string

type BlockNumberOutput string

func (o BlockNumberOutput) Decode() uint64 {
	return parityrpc.HexToInt(o)
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

type GetBlockByNumberOutput struct {
	Number           string      `json:"number"`
	Hash             string      `json:"hash"`
	ParentHash       string      `json:"parentHash"`
	SealFields       []string    `json:"sealFields"`
	Sha3Uncles       string      `json:"sha3Uncles"`
	LogsBloom        string      `json:"logsBloom"`
	TransactionsRoot string      `json:"transactionsRoot"`
	StateRoot        string      `json:"stateRoot"`
	Miner            string      `json:"miner"`
	Difficulty       string      `json:"difficulty"`
	TotalDifficulty  string      `json:"totalDifficulty"`
	ExtraData        string      `json:"extraData"`
	Size             string      `json:"size"`
	GasLimit         string      `json:"gasLimit"`
	MinGasPrice      string      `json:"minGasPrice"`
	GasUsed          string      `json:"gasUsed"`
	Timestamp        string      `json:"timestamp"`
	Transactions     interface{} `json:"transactions"`
	Uncles           []string    `json:"uncles"`
}

func (o SyncingOutput) DecodeAll() SyncingOutputDecoded {
	return SyncingOutputDecoded{
		Syncing:             o.Syncing,
		StartingBlock:       parityrpc.HexToInt(o.StartingBlock),
		CurrentBlock:        parityrpc.HexToInt(o.CurrentBlock),
		HighestBlock:        parityrpc.HexToInt(o.HighestBlock),
		WarpChunksAmount:    parityrpc.HexToInt(o.WarpChunksAmount),
		WarpChunksProcessed: parityrpc.HexToInt(o.WarpChunksProcessed),
	}
}
