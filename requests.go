package parityrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Generic struct for a POST request to a node's RPC interface
// A raw interface is used for Params to allow variable input and let json
// do it's thing.

type PostPayload struct {
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      uint64      `json:"id"`
	JsonRPC string      `json:"jsonrpc"`
}

// Generic struct for a response from a parity node

type ParityResponse struct {
	JsonRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
}

func (c *ParityNode) GenericCall(method string, input interface{}, output interface{}) (err error) {

	// Assuming an RPC endpoint that doesn't change it's types on you
	// (I'm looking at you eth_syncing)
	// this function can probably be used to fill out most of the endpoints

	if c.Debug {
		log.Printf("DEBUG: Using GenericCall for method: %s input: %+v\n", method, input)
	}
	resp, suc, err := c.Post(method, input)
	if err != nil {
		return
	}
	if !suc {
		output = false
		return
	}
	json.Unmarshal(resp.Result, &output)
	return
}

func (c *ParityNode) Post(method string, params interface{}) (response ParityResponse, success bool, err error) {

	// Do a post request to a Parity endpoint with a given struct of parameters

	payload, err := c.GenPayload(method, params)

	if c.Debug {
		log.Printf("DEBUG: Payload generated: %s\n", string(payload))
	}

	if err != nil {
		log.Printf("ERROR: Failed to generate payload from params: %v\n", params)
		return
	}

	req, err := GenRequest(c.Host, payload)
	if err != nil {
		log.Printf("ERROR: Failed to initialize request to %s with payload %s", c.Host, string(payload))
		return
	}

	client := &http.Client{}
	if c.Debug {
		log.Printf("DEBUG: Posting payload to %s\n", c.Host)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("ERROR: Failed to receive response from %s\n", c.Host)
		return
	}

	defer resp.Body.Close()
	if c.Debug {
		log.Printf("DEBUG: Received response from %s. Attemtping to unmarshal.\n", c.Host)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	json.Unmarshal(data, &response)
	if c.Debug {
		log.Printf("DEBUG: Raw result: %s\n", string(response.Result))
	}

	// eth_syncing (and I may find others) will change it's type from a map to
	// a bool when it is done syncing. That annoys me.

	if string(response.Result) == "false" {
		success = false
	} else {
		success = true
	}

	// increment the ID counter

	c.IdCounter += 1
	if c.Debug {
		log.Printf("DEBUG: ID counter incremented to %v\n", c.IdCounter)
	}
	return

}

func (c *ParityNode) GenPayload(method string, params interface{}) (payload []byte, err error) {

	// json encode a jsonrpc payload

	pstruct := PostPayload{
		Method:  method,
		Params:  params,
		Id:      c.IdCounter,
		JsonRPC: JSON_RPC_VERSION,
	}
	payload, err = json.Marshal(pstruct)
	return

}

func GenRequest(host string, payload []byte) (req *http.Request, err error) {

	// create an HTTP request and set the Content-Type header

	req, err = http.NewRequest("POST", host, bytes.NewBuffer(payload))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	return

}

func HexToInt(hex interface{}) uint64 {

	hexStr := fmt.Sprintf("%v", hex)
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexStr, "0x", "", -1)

	// base 16 for hexadecimal
	result, _ := strconv.ParseUint(cleaned, 16, 64)
	return uint64(result)
}

func IntToHex(i int64) string {
	return fmt.Sprintf("0x%s", strconv.FormatInt(i, 16))
}
