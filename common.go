package bca

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

type Signature struct {
	HTTPMethod  string
	RelativeURL string
	Timestamp   string
	RequestBody string
	AccessToken string
	APISecret   string
}

func generateBcaSignature(bs Signature) (string, error) {
	var stringToSignature string

	encRequestBody, err := encodeRequestBody(bs.RequestBody)
	if err != nil {
		return "", err
	}

	// TODO: there's special rule for RelativeURL, we can forget it for now. more on:https://developer.bca.co.id/documentation/#signature
	stringToSignature = fmt.Sprintf("%v:%v:%v:%v:%v", bs.HTTPMethod, bs.RelativeURL, bs.AccessToken, encRequestBody, bs.Timestamp)

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(bs.APISecret))

	// Write Data to it
	h.Write([]byte(stringToSignature))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	return sha, nil
}

func getBcaTimestamp() string {
	// 'YYYY-MM-DDTHH24:MI:SS.ZZZ+07:00'
	jakarta := time.FixedZone("Asia/Jakarta", 7*60*60)
	t := time.Now().In(jakarta)
	return t.Format("2006-01-02T15:04:05.000-07:00")
}

func encodeRequestBody(requestBody string) (string, error) {
	hash := sha256.New()
	if _, err := hash.Write([]byte(cleanString(requestBody))); err != nil {
		return "", err
	}

	return strings.ToLower(hex.EncodeToString(hash.Sum(nil))), nil
}

func cleanString(any string) string {
	s := strings.ReplaceAll(any, " ", "")
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\t", "")
	return s
}

func bcaConvertTimestamp(req int64) string {
	timeReq := time.Unix(req, 0)
	return timeReq.Format("2006-01-02T15:04:05.000-07:00")
}
