// Package cloudflare provides convenience constructors for the Terraform
// Cloudflare provider. They are optional sugar over tfgen types; you can always
// construct them directly.
package cloudflare

import "github.com/ably/tfgen"

const source = "cloudflare/cloudflare"

// ProviderVersion returns a ProviderRequirement for the Cloudflare provider
// pinned to the given version constraint, e.g. "~> 4.0".
func ProviderVersion(version string) tfgen.ProviderRequirement {
	return tfgen.ProviderRequirement{Source: source, Version: version}
}

// Provider creates a Cloudflare provider. Credentials are typically supplied
// via environment variables rather than in configuration.
func Provider() *tfgen.Provider {
	return &tfgen.Provider{
		Name:   "cloudflare",
		Config: tfgen.Config{},
	}
}
