package socialauth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// AuthProvider is an interface that provides a method for token verification
type AuthProvider interface {
	// Identifier returns the string identifier of the provider
	Identifier() string

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
