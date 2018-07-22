package eth

import (
	"encoding/json"

	"github.com/tinyzimmer/parity-go"
)

const (
	ETH_ACCOUNTS     = "eth_accounts"
	ETH_BLOCK_NUMBER = "eth_blockNumber"
	ETH_SYNCING      = "eth_syncing"
	ETH_CALL         = "eth_call"
)

type Client struct {
	Node parity.ParityNode
}

func NewClient(node parity.ParityNode) (c Client) {
	c.Node = node
	return
}

func (c *Client) Accounts() (response EthAccountsOutput, err error) {
	response = make(EthAccountsOutput, 0)
	err = c.Node.GenericCall(ETH_ACCOUNTS, EthAccountsInput{}, &response)
	return
}

func (c *Client) BlockNumber() (response EthBlockNumberOutput, err error) {
	err = c.Node.GenericCall(ETH_BLOCK_NUMBER, EthBlockNumberInput{}, &response)
	return
}

func (c *Client) Syncing() (response EthSyncingOutput, err error) {
	resp, suc, err := c.Node.Post(ETH_SYNCING, EthSyncingInput{})
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
