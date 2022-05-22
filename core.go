package bca

import (
	"fmt"
	"io"
	"net/url"
	"strings"
)

const (
	TOKEN_PATH = "/api/oauth/token"
	VA_PATH    = "/va/payments"
)

// CoreGateway struct
type CoreGateway struct {
	Client Client
}

// Call : base method to call Core API
func (gateway *CoreGateway) Call(method, path string, header map[string]string, body io.Reader, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = gateway.Client.BaseUrl + path

	return gateway.Client.Call(method, path, header, body, v)
}

func (gateway *CoreGateway) GetToken() (res TokenResponse, err error) {

	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	err = gateway.Call("POST", TOKEN_PATH, headers, strings.NewReader(data.Encode()), &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (gateway *CoreGateway) InquiryCustomerNumber(token string, req InquiryByCustomerNumber) (resp InquiryByCustomerNumberResp, err error) {

	// query params
	q := url.Values{}
	q.Set("CustomerNumber", req.CustomerNumber)
	q.Set("CompanyCode", req.CompanyCode)
	path := fmt.Sprint(VA_PATH, "?", q.Encode())

	// bca signature
	signature := Signature{
		APISecret:   gateway.Client.ApiSecret,
		AccessToken: token,
		HTTPMethod:  "GET",
		RelativeURL: path,
		RequestBody: "",
		Timestamp:   getBcaTimestamp(),
	}
	bcaSignature, err := generateBcaSignature(signature)
	if err != nil {
		return
	}

	headers := map[string]string{
		"Authorization":   fmt.Sprintf("Bearer %v", token),
		"Content-Type":    "application/json",
		"Origin":          gateway.Client.Origin,
		"X-BCA-Key":       gateway.Client.ApiKey,
		"X-BCA-Timestamp": signature.Timestamp,
		"X-BCA-Signature": bcaSignature,
	}

	err = gateway.Call("GET", path, headers, nil, &resp)
	if err != nil {
		return
	}

	return
}

func (gateway *CoreGateway) InquiryRequestID(token string, req InquiryByRequestID) (resp InquiryByCustomerNumberResp, err error) {

	// query params
	q := url.Values{}
	q.Set("RequestID", req.RequestID)
	q.Set("CompanyCode", req.CompanyCode)
	path := fmt.Sprint(VA_PATH, "?", q.Encode())

	// bca signature
	signature := Signature{
		APISecret:   gateway.Client.ApiSecret,
		AccessToken: token,
		HTTPMethod:  "GET",
		RelativeURL: path,
		RequestBody: "",
		Timestamp:   getBcaTimestamp(),
	}
	bcaSignature, err := generateBcaSignature(signature)
	if err != nil {
		return
	}

	headers := map[string]string{
		"Authorization":   fmt.Sprintf("Bearer %v", token),
		"Content-Type":    "application/json",
		"Origin":          gateway.Client.Origin,
		"X-BCA-Key":       gateway.Client.ApiKey,
		"X-BCA-Timestamp": signature.Timestamp,
		"X-BCA-Signature": bcaSignature,
	}

	err = gateway.Call("GET", path, headers, nil, &resp)
	if err != nil {
		return
	}

	return
}
