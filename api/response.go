package api

// Response represents the response
// from the HandCash receive API.
type Response struct {
	ReceivingAddress string `json:"receivingAddress,omitempty"`
	PublicKey        string `json:"publicKey,omitempty"`
	Error            string `json:"error,omitempty"`
}
