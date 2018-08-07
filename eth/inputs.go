package eth

type AccountsInput []string

type BlockNumberInput []string

type SyncingInput []string

type GetBlockByNumberInput struct {
	BlockNumber  string
	ShowFullTxns bool
}
