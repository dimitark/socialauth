#SocialAuth.go

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

You need to configure the library, before you can use it.

```go
// Configure the SocialAuth
socialAuth := socialauth.NewSocialAuth(
	socialauth.NewFacebookAuthProvider("YOUR_APP_ACCESS_TOKEN"),
	socialauth.NewGoogleAuthProvider("YOUR_GOOGLE_APP_CLIENT_ID"),
)
```

The library can be used in two ways:

* independently
* as middleware


### Independent usage

```go
if provider := socialAuth.Get("facebook"); provider != nil {
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
n.Use(socialauth.NewMiddleware(socialAuth, context.Set))
```

License
=======

    The MIT License (MIT)

    Copyright (c) 2016, Dimitar Kotevski

    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:

    The above copyright notice and this permission notice shall be included in
    all copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
    THE SOFTWARE.