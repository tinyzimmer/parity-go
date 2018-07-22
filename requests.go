package parity

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

type PostPayload struct {
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      int         `json:"id"`
	JsonRPC string      `json:"jsonrpc"`
}

type ParityResponse struct {
	JsonRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
}

func (c *ParityNode) GenericCall(method string, input interface{}, output interface{}) (err error) {
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
	if string(response.Result) == "false" {
		success = false
	} else {
		success = true
	}
	c.IdCounter += 1
	if c.Debug {
		log.Printf("DEBUG: ID counter incremented to %v\n", c.IdCounter)
	}
	return
}

func (c *ParityNode) GenPayload(method string, params interface{}) (payload []byte, err error) {
	pstruct := PostPayload{
		Method:  method,
		Params:  params,
		Id:      1,
		JsonRPC: JSON_RPC_VERSION,
	}
	payload, err = json.Marshal(pstruct)
	return
}

func GenRequest(host string, payload []byte) (req *http.Request, err error) {
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
