package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strings"
)

type Header struct {
	Alg  string `json:"alg"`
	Type string `json:"type"`
}

type Payload struct {
	Sub    string `json:"sub"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

func CreateJwt(data Payload, secret string) (string, error) {
	header := Header{
		Alg:  "HS256",
		Type: "JWT",
	}

	byteArrHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerB64 := Base64UrlEncode(byteArrHeader)

	byteArrPayload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	payloadB64 := Base64UrlEncode(byteArrPayload)

	byteArrSecret := []byte(secret)
	message := headerB64 + "." + payloadB64
	byteArrMessage := []byte(message)

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)
	signature := Base64UrlEncode(h.Sum(nil))

	return headerB64 + "." + payloadB64 + "." + signature, nil
}

func VerifyJwt(jwtToken string, secret string) bool {
	// check token empty
	if jwtToken == "" {
		return false
	}

	// check bearer
	if !strings.HasPrefix(jwtToken, "Bearer") {
		return false
	}

	token := strings.TrimPrefix(jwtToken, "Bearer ")

	// check token format
	tokenParts := strings.Split(token, ".")
	if len(tokenParts) != 3 {
		return false
	}

	jwtHeader := tokenParts[0]
	jwtPayload := tokenParts[1]
	oldSignature := tokenParts[2]

	message := jwtHeader + "." + jwtPayload

	byteArrSecret := []byte(secret)
	byteArrMessage := []byte(message)

	hMac := hmac.New(sha256.New, byteArrSecret)
	hMac.Write(byteArrMessage)
	newSignature := Base64UrlEncode(hMac.Sum(nil))

	// verify new token with old
	if oldSignature != newSignature {
		return false
	}

	return true
}

func Base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
