package socialauth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// ProviderService is the enum that defines the available providers
type Provider int

const (
	unknown Provider = iota
	// Facebook is the Facebook auth provider
	Facebook Provider = iota
	// Google is the Google auth provider
	Google Provider = iota
)

// ProviderFromName returns the provider from a given name
// Returns unknown if the given string is gibberish
func ProviderFromName(name string) Provider {
	switch strings.ToLower(name) {
	case "facebook":
		return Facebook
	case "google":
		return Google
	default:
		return unknown
	}
}

// AuthProvider is an interface that provides a method for token verification
type AuthProvider interface {
	// Verifies the given token against the server's provider (Facebook, Google...)
	// And returns the user ID or an error
	VerifyToken(userToken string) (string, error)
}

// Fetches the given URL and returns the response
// as JSON (through the target parameter)
func getJSON(url string, target interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	byteData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteData, target)
	if err != nil {
		return err
	}
	return nil
}
