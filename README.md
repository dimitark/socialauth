#Social Auth .go

The SocialAuth package enables you to verify the OAuth2 authentication tokens issued by the social providers. 

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

Before using the library - it needs to be configured. The library expects to have valid application IDs for every configured provider. 

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

After setting up the configuration, the library can be used in two ways:

* independently
* as a middleware


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

Usually this library will be used as part of a middleware. For that reason, the library provides a middleware implementation, that plays nicelly with the **negroni & gorilla** libraries.

```go
n := negroni.New()
n.Use(socialauth.NewMiddleware(config, context.Set))
```