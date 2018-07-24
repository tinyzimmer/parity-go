package parity

import (
	"github.com/tinyzimmer/parity-go"
)

const (
	PARITY_MODE = "parity_mode"
)

type Client struct {

	// struct representation of a client for the eth API

	Node parityrpc.ParityNode
}

func NewClient(node parityrpc.ParityNode) (c Client) {

	// set the Parity node for the client

	c.Node = node
	return

}

func (c *Client) Mode() (response ModeOutput, err error) {

	response = ""
	err = c.Node.GenericCall(PARITY_MODE, ModeInput{}, &response)
	return

}
