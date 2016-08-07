package socialauth

import (
	"fmt"
	"net/http"
)

// ContextUserID is the key under which the UserID
// of the current user is stored in the given context,
// through the ContextSetFunc
// Only for Authenticated requests
const ContextUserID = "SocialAuthUserID"

// ContextSetFunc is the function that sets a value for a given key and request in a context.
// Usually this is the Gorilla Context -> context.Set -> github.com/gorilla/context
type ContextSetFunc func(*http.Request, interface{}, interface{})

// Middleware is a middleware for the http package
// It verifies if the request has a valid token
// The token is verified by one of the Social AuthProviders (Facebook, Google,...)
type Middleware struct {
	socialAuth     *SocialAuth
	contextSetFunc ContextSetFunc
}

// NewMiddleware creates a new Middleware with the given configuration
func NewMiddleware(socialAuth *SocialAuth, contextSetFunc ContextSetFunc) *Middleware {
	return &Middleware{
		socialAuth:     socialAuth,
		contextSetFunc: contextSetFunc,
	}
}

func (am *Middleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	success := false
	defer func() {
		// Recover
		if r := recover(); r != nil {
			fmt.Println("Recovered from the error:", r)
		}

		if !success {
			rw.WriteHeader(http.StatusForbidden)
			fmt.Fprintln(rw, "403 Forbidden")
		}
	}()

	providerHeaders := r.Header["X-Auth-Provider"]
	tokenHeaders := r.Header["X-Auth-Token"]
	if len(providerHeaders) > 0 && len(tokenHeaders) > 0 {
		if provider := am.socialAuth.Get(ProviderFromName(providerHeaders[0])); provider != nil {
			userID, err := provider.VerifyToken(tokenHeaders[0])
			if err != nil {
				fmt.Println("Auth error:", err)
				return
			}

			fmt.Println("Authentication successfull. User ID:", userID)
			success = true
			am.contextSetFunc(r, ContextUserID, userID)
			next(rw, r)
			return
		} else {
			fmt.Println("Unknown provider:", providerHeaders[0])
			return
		}
	}

	fmt.Println("X-Auth-Provider & X-Auth-Token must be provided!")
}
