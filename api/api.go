package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	// MAINNET is the main network.
	MAINNET = "mainnet"
	// TESTNET is the test network.
	TESTNET = "testnet"

	mainURL = "https://api.handcash.io/api"
	testURL = "https://test-api.handcash.io/api"
)

var network = MAINNET

// Receive takes a handle and returns the
// Response struct with the address.
func Receive(handle string) (*Response, error) {
	handle = strings.TrimLeft(handle, "$")

	if handle == "" {
		return nil, errors.New("handle is required")
	}

	url := fmt.Sprintf("%s/receivingAddress/%s", getURL(), handle)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := &Response{}

	if len(body) == 0 {
		r.Error = "not found"
		return r, nil
	}

	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// getURL gets the appropriate URL for the network.
func getURL() string {
	switch network {
	case MAINNET:
		return mainURL
	case TESTNET:
		return testURL
	}

	return ""
}

// SetNetwork allows you to set the network.
// Valid values are mainnet/testnet.
func SetNetwork(n string) error {
	switch n {
	case MAINNET:
		network = MAINNET
	case TESTNET:
		network = TESTNET
	default:
		return fmt.Errorf("invalid network - %s", n)
	}

	return nil
}
