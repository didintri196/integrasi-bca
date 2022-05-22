package bca

import (
	"fmt"
	"testing"
)

func initBca() Client {
	client := NewClient()
	client.BaseUrl = "https://devapi.klikbca.com:9443"
	client.ApiKey = "a16c5bb4-49d1-4a12-9194-db3df367d893"
	client.ApiSecret = "2ad77de8-7f0e-4379-bce5-71d70529a611"
	client.ClientID = "e305a76a-78d3-4f92-b734-c23ae58c97d8"
	client.ClientSecret = "04031743-9645-4bc7-84e5-c8943618c2c8"
	client.CompanyID = "uatcorp001"
	client.Origin = "https://dev-api.goldentalipodo.com"
	client.LogLevel = 3
	return client
}

func TestGetToken(t *testing.T) {
	core := CoreGateway{
		Client: initBca(),
	}

	res, err := core.GetToken()
	t.Log("RES :", res)
	t.Log("ERR :", err)
}

func TestInquiryCustomerNumber(t *testing.T) {
	core := CoreGateway{
		Client: initBca(),
	}

	respToken, err := core.GetToken()
	t.Log("ERR_TOKEN", err)

	req := InquiryByCustomerNumber{
		CompanyCode:    "54321",
		CustomerNumber: "123456789012345678901234567890",
	}

	resp, err := core.InquiryCustomerNumber(respToken.AccessToken, req)
	fmt.Println("RESP :", resp)
	fmt.Println("ERR :", err)
}

func TestInquiryRequestID(t *testing.T) {
	core := CoreGateway{
		Client: initBca(),
	}

	respToken, err := core.GetToken()
	t.Log("ERR_TOKEN", err)

	req := InquiryByRequestID{
		CompanyCode: "54321",
		RequestID:   "202104201031475432100000292222",
	}

	resp, err := core.InquiryRequestID(respToken.AccessToken, req)
	fmt.Println("RESP :", resp)
	fmt.Println("ERR :", err)
}
