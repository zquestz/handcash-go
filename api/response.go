package api

// Response represents the response
// from the HandCash receive API.
type Response struct {
	ReceivingAddress string `json:"receivingAddress,omitempty"`
	CashAddr         string `json:"cashAddr,omitempty"`
	PublicKey        string `json:"publicKey,omitempty"`
	Error            string `json:"error,omitempty"`
}
