package bca

import "time"

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

type InquiryByCustomerNumberResp struct {
	TransactionData []struct {
		DetailBills []struct {
			BillReference string `json:"BillReference"`
			BillNumber    string `json:"BillNumber"`
		} `json:"DetailBills"`
		PaymentFlagStatus string    `json:"PaymentFlagStatus"`
		RequestID         string    `json:"RequestID"`
		Reference         string    `json:"Reference"`
		TotalAmount       string    `json:"TotalAmount"`
		TransactionDate   time.Time `json:"TransactionDate"`
		PaidAmount        string    `json:"PaidAmount"`
	} `json:"TransactionData"`
}
