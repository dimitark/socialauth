package socialauth

import "errors"

// The GoogleAuthProvider validates the token against the google api
type GoogleAuthProvider struct {
	validator *GoogleTokenValidator
}

func (p GoogleAuthProvider) Key() string {
	return "google"
}

// Creates a new Google Authentication Provider
// The required appClientId can be found in your Google API's Dashboard
// https://developers.google.com/identity/protocols/OAuth2
func NewGoogleAuthProvider(appClientId string) *GoogleAuthProvider {
	return &GoogleAuthProvider{
		validator: NewGoogleTokenValidator(appClientId),
	}
}

// Returns the identifier of this provider
func (p *GoogleAuthProvider) Identifier() string {
	return "google"
}

// VerifyToken verifies the given token against the server's provider (Facebook, Google...)
// And returns the user ID or an error
func (p *GoogleAuthProvider) VerifyToken(userToken string) (string, error) {
	if p.validator == nil {
		return "", errors.New("The token validator is not initialized properly!")
	}

	userID, err := p.validator.Validate(userToken)

	if err != nil {
		return "", err
	}

	return userID, nil
}
