package socialauth

import "strings"

// SocialAuth holds all the configured providers
// Use an instance of this struct to work with this library
type SocialAuth struct {
	// The configured providers
	providers map[string]AuthProvider
}

// NewSocialAuth creates and returns a new instance
// of the SocialAuth with the given providers
func NewSocialAuth(providers ...AuthProvider) *SocialAuth {
	// Configure everything
	mapOfProviders := make(map[string]AuthProvider, len(providers))
	for _, provider := range providers {
		mapOfProviders[strings.ToLower(provider.Identifier())] = provider
	}

	// Create and return the instance
	return &SocialAuth{
		providers: mapOfProviders,
	}
}

// Get returns the asked provider. Nil if such provider is not configured.
func (sa *SocialAuth) Get(provider string) AuthProvider {
	return sa.providers[strings.ToLower(provider)]
}
