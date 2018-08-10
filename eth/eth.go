package eth

import (
	"encoding/json"

	"github.com/tinyzimmer/parity-go"
)

const (
	ETH_ACCOUNTS            = "eth_accounts"
	ETH_BLOCK_NUMBER        = "eth_blockNumber"
	ETH_SYNCING             = "eth_syncing"
	ETH_CALL                = "eth_call"
	ETH_GET_BLOCK_BY_NUMBER = "eth_getBlockByNumber"
)

type Client struct {

	// struct representation of a client for the eth API

	Node parityrpc.ParityNode
}

func NewClient(node parityrpc.ParityNode) (c *Client) {

	// set the Parity node for the client

	c.Node = node
	return

}

func (c *Client) Accounts() (response AccountsOutput, err error) {

	response = make(AccountsOutput, 0)
	err = c.Node.GenericCall(ETH_ACCOUNTS, AccountsInput{}, &response)
	return

}

func (c *Client) BlockNumber() (response BlockNumberOutput, err error) {

	err = c.Node.GenericCall(ETH_BLOCK_NUMBER, BlockNumberInput{}, &response)
	return

}

func (c *Client) Syncing() (response SyncingOutput, err error) {

	// When the client is done syncing, it simply returns a false.
	// This is stupid on so many levels. I'd rather a response with a matching
	// currentBlock and highestBlock. That would actually make sense and not
	// require this insane conditional. See how easy the others are?!

	resp, suc, err := c.Node.Post(ETH_SYNCING, SyncingInput{})
	if err != nil {
		return
	}
	if !suc {
		response.Syncing = false
	} else {
		json.Unmarshal(resp.Result, &response)
		response.Syncing = true
	}
	return

}

func (c *Client) GetBlockByNumber(input GetBlockByNumberInput) (response GetBlockByNumberOutput, err error) {
	convertedInput := []interface{}{input.BlockNumber, input.ShowFullTxns}
	err = c.Node.GenericCall(ETH_GET_BLOCK_BY_NUMBER, convertedInput, &response)
	return
}
