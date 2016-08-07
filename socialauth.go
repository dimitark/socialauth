package socialauth

import "errors"

// SocialAuth holds all the configured providers
// Use an instance of this struct to work with this library
type SocialAuth struct {
	// The configured providers
	providers map[Provider]AuthProvider
}

// NewSocialAuthWithConfigs creates and returns a new instance of the SocialAuth
// with the given configurations
func NewSocialAuthWithConfigs(configs map[Provider]map[string]string) *SocialAuth {
	sa := &SocialAuth{providers: make(map[Provider]AuthProvider)}

	// Configure everything
	for provider, config := range configs {
		sa.ConfigureProvider(provider, config)
	}

	// Return
	return sa
}

// ConfigureProvider configures the given provider
func (sa *SocialAuth) ConfigureProvider(provider Provider, config map[string]string) error {
	configuredProvider := getProvider(provider, config)
	if configuredProvider == nil {
		return errors.New("Cannot configure an unknown provider!")
	}

	// The provider exists and it's successfully configured
	sa.providers[provider] = configuredProvider
	return nil
}

// Get returns the asked provider. Nil if such provider is not configured.
func (sa *SocialAuth) Get(provider Provider) AuthProvider {
	return sa.providers[provider]
}
