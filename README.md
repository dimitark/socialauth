#Social Auth .go

The SocialAuth package enables you to verify the OAuth2 authentication tokens issued by social providers. 

At the moment only two providers are available (Facebook and Google).  

## Installation

```
go get github.com/dimitark/socialauth
```

```go
import "github.com/dimitark/socialauth"
```

## Usage

### Configuration

You need to configure the client, before you can use it. The library requires valid application IDs for each configured provider. 

```go
config := map[socialauth.Provider]map[string]string{
		socialauth.Facebook: map[string]string{
			"appAccessToken": "YOUR_APP_ACCESS_TOKEN",
		},
		socialauth.Google: map[string]string{
			"appClientId": "YOUR_GOOGLE_APP_CLIENT_ID",
		},
	}
```

The library can be used in two ways:

* independently
* as middleware


### Independent usage

```go
auth := socialAuth.NewSocialAuthWithConfigs(config)
if provider := auth.Get(socialauth.Facebook); provider != nil {
	userID, err := provider.VerifyToken("USER_TOKEN")
	if err != nil {
		// Authentication failed!
	} else {
		// Success
	}
}
```

### Middleware usage

You can use it as a middleware for both **negroni** OR **gorilla** 

An example using **negroni**:

```go
n := negroni.New()
n.Use(socialauth.NewMiddleware(config, context.Set))
```
