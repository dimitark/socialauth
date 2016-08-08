package socialauth

import "fmt"

// The URL for validating the token
const fBTokenValidationURL = "https://graph.facebook.com/debug_token?input_token=%s&access_token=%s"

// The DebugToken response
type fBDebugTokenResp struct {
	Data struct {
		UserID string `json:"user_id"`
	} `json:"data"`
}

// FBAuthProvider - The Facebook Auth provider implementation
type FacebookAuthProvider struct {
	appAccessToken string
}

// Creates a new Facebook Authentication Provider instance
// The required appAccessToken can be obtained by Facebook's API ->
// https://developers.facebook.com/docs/facebook-login/access-tokens/#apptokens
func NewFacebookAuthProvider(appAccessToken string) *FacebookAuthProvider {
	return &FacebookAuthProvider{
		appAccessToken: appAccessToken,
	}
}

// Returns the identifier of this provider
func (p *FacebookAuthProvider) Identifier() string {
	return "facebook"
}

// VerifyToken verifies the given token against the Facebook's server
// and returns the user ID or an error
func (p *FacebookAuthProvider) VerifyToken(userToken string) (string, error) {
	resp := &fBDebugTokenResp{}
	err := getJSON(fmt.Sprintf(fBTokenValidationURL, userToken, p.appAccessToken), resp)

	if err != nil {
		return "", err
	}
	return resp.Data.UserID, nil
}
