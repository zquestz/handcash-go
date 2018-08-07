package api_test

import (
	"testing"

	"github.com/ad2games/vcr-go"
	"github.com/zquestz/handcash-go/api"
)

// TestReceive tests receiving handle data from
// the HandCash API.
func TestReceive(t *testing.T) {
	vcr.Start("handcash", nil)
	defer vcr.Stop()

	validUser := "rjseibane"

	api.SetNetwork("mainnet")

	resp, _ := api.Receive(validUser)
	if resp.Error != "" {
		t.Fatalf("Valid user %s returned error %s", validUser, resp.Error)
	}

	address := "1KArpdjbJ52xuydEWePAhjbZU8HSxxVTuo"
	if resp.ReceivingAddress != address {
		t.Errorf("Invalid ReceivingAddress %s expected %s", resp.ReceivingAddress, address)
	}

	pubKey := "020d82ae1be2e3b6e59446d476d7012fe240730e9cb801265969a535eee332034f"
	if resp.PublicKey != pubKey {
		t.Errorf("Invalid PubKey %s expected %s", resp.PublicKey, pubKey)
	}

	api.SetNetwork("testnet")

	testResp, _ := api.Receive(validUser)
	if testResp.Error != "" {
		t.Fatalf("Valid user %s returned error %s", validUser, testResp.Error)
	}

	testAddress := "mxszqDyaNGFcmTkPjJ2BGRpSTChdVWaNPZ"
	if testResp.ReceivingAddress != testAddress {
		t.Errorf("Invalid ReceivingAddress %s expected %s", testResp.ReceivingAddress, testAddress)
	}

	testPubKey := "03d193439a2f06ed1121be5b4e61381386ffee5ec5bec33daf17e33ccb34622753"
	if testResp.PublicKey != testPubKey {
		t.Errorf("Invalid PubKey %s expected %s", testResp.PublicKey, testPubKey)
	}

	invalidUser := "nope"
	expectedErr := "not found"

	invalidResp, _ := api.Receive(invalidUser)
	if invalidResp.Error != expectedErr {
		t.Fatalf("Invalid Error %s expected %s", invalidResp.Error, expectedErr)
	}
}

// TestSetNetwork tests setting the network.
// We only test the edge case here.
func TestSetNetwork(t *testing.T) {
	// Should return an error for anything else.
	invalidNetwork := "anything"

	aErr := api.SetNetwork(invalidNetwork)
	if aErr == nil {
		t.Errorf("error should have been returned for invalid network: %s", invalidNetwork)
	}
}
