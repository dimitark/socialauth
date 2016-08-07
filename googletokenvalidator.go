package socialauth

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/mendsley/gojwk"
)

// TokenInfo
type TokenInfo struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	AtHash        string `json:"at_hash"`
	Aud           string `json:"aud"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Local         string `json:"locale"`
	Iss           string `json:"iss"`
	Azp           string `json:"azp"`
	Iat           int64  `json:"iat"`
	Exp           int64  `json:"exp"`
}

type certs struct {
	Keys []gojwk.Key `json:"keys"`
}

// GoogleTokenValidator validates the given token
// using the Google's public keys
type GoogleTokenValidator struct {
	publicKeys []*rsa.PublicKey // Google's Public keys (https://www.googleapis.com/oauth2/v3/certs)
	aud        string           // The AppClient ID
}

// NewGoogleTokenValidator creates a new instance of the GoogleTokenValidator struct
// Configures it and returns it
func NewGoogleTokenValidator(appClientID string) *GoogleTokenValidator {
	return &GoogleTokenValidator{
		publicKeys: getGooglePublicKeys(),
		aud:        appClientID,
	}
}

// Validate validates the given token.
// If the token is invalid - it returns an error
func (v *GoogleTokenValidator) Validate(token string) (string, error) {
	payload, signature, messageToSign, err := divideToken(token)

	// Check for errors
	if err != nil {
		return "", errors.New("Invalid token format!")
	}

	// Get the token info
	tokenInfo := getTokenInfo(payload)

	// The AUD in the TokenInfo must match with the ApClientID
	if v.aud != tokenInfo.Aud {
		return "", errors.New("The AUD from the Token doesn't match with the AppClientID")
	}

	// Check the ISS
	if (tokenInfo.Iss != "accounts.google.com") && (tokenInfo.Iss != "https://accounts.google.com") {
		return "", errors.New("The ISS is not valid!")
	}

	// Check if the token has expired
	if (time.Now().Unix() < tokenInfo.Iat) || (time.Now().Unix() > tokenInfo.Exp) {
		return "", errors.New("The token has expired!")
	}

	// Check the signature
	for _, key := range v.publicKeys {
		if rsa.VerifyPKCS1v15(key, crypto.SHA256, messageToSign, signature) == nil {
			return tokenInfo.Sub, nil
		}
	}

	return "", errors.New("Invalid token format!")
}

func divideToken(token string) ([]byte, []byte, []byte, error) {
	args := strings.Split(token, ".")
	if len(args) != 3 {
		return nil, nil, nil, errors.New("Invalid token format!")
	}
	return safeDecode(args[1]), safeDecode(args[2]), toSHA256(args[0] + "." + args[1]), nil
}

func getGooglePublicKeys() []*rsa.PublicKey {
	res, _ := http.Get("https://www.googleapis.com/oauth2/v3/certs")
	certsResp, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	var certs certs
	json.Unmarshal(certsResp, &certs)

	var publicKeys = make([]*rsa.PublicKey, len(certs.Keys))
	for index, key := range certs.Keys {
		publicKey, _ := key.DecodePublicKey()
		switch publicKey.(type) {
		case *rsa.PublicKey:
			publicKeys[index] = publicKey.(*rsa.PublicKey)
		default:
			fmt.Println("The public key is not a RSA public key")
		}
	}
	return publicKeys
}

func safeDecode(str string) []byte {
	lenMod4 := len(str) % 4
	if lenMod4 > 0 {
		str = str + strings.Repeat("=", 4-lenMod4)
	}

	bytes, _ := base64.URLEncoding.DecodeString(str)
	return bytes
}

func toSHA256(str string) []byte {
	a := sha256.New()
	a.Write([]byte(str))
	return a.Sum(nil)
}

func getTokenInfo(bt []byte) *TokenInfo {
	var a *TokenInfo
	json.Unmarshal(bt, &a)
	return a
}
