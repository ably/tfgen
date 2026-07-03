// Package ably provides convenience constructors for the Terraform Ably
// provider (https://github.com/ably/terraform-provider-ably).
//
// These are thin wrappers that emit tfgen types; Terraform validates the
// resulting configuration. Each constructor takes the Terraform resource label
// as its first argument, which also seeds the resource's primary identifier
// (name or id), and an optional options struct for the remaining attributes.
// Anything not covered here can be expressed with a plain tfgen.Resource.
package ably

import "github.com/ably/tfgen"

const source = "ably/ably"

// DefaultVersion is the Ably provider version pinned by ProviderVersion when no
// explicit constraint is given.
const DefaultVersion = "1.0.0"

// ProviderVersion returns a ProviderRequirement for the Ably provider. With no
// argument it pins DefaultVersion; pass a constraint (e.g. "~> 1.0") to
// override.
func ProviderVersion(version ...string) tfgen.ProviderRequirement {
	return tfgen.ProviderRequirement{Source: source, Version: firstOr(version, DefaultVersion)}
}

// Provider creates an Ably provider. Credentials are typically supplied via the
// ABLY_ACCOUNT_TOKEN environment variable rather than in configuration.
func Provider() *tfgen.Provider {
	return &tfgen.Provider{
		Name:   "ably",
		Config: tfgen.Config{},
	}
}

// AppOptions configures an ably_app resource. Field names mirror the
// terraform-provider-ably schema; zero-valued fields are omitted so Terraform
// applies its own defaults.
type AppOptions struct {
	Name    string // defaults to the resource label
	Status  string // "enabled" or "disabled"
	TLSOnly bool
	FCMKey  string
}

// App creates an ably_app resource. The label seeds the app's name.
func App(label string, opts ...AppOptions) *tfgen.Resource {
	o := first(opts)

	config := tfgen.Config{"name": label}
	if o.Name != "" {
		config["name"] = o.Name
	}
	if o.Status != "" {
		config["status"] = o.Status
	}
	if o.TLSOnly {
		config["tls_only"] = true
	}
	if o.FCMKey != "" {
		config["fcm_key"] = o.FCMKey
	}

	return &tfgen.Resource{Type: "ably_app", Name: label, Config: config}
}

// APIKeyOptions configures an ably_api_key resource. AppID and Capabilities are
// required by Terraform.
type APIKeyOptions struct {
	Name            string              // defaults to the resource label
	AppID           string              // required
	Capabilities    map[string][]string // required; channel/namespace -> capabilities
	RevocableTokens bool
}

// APIKey creates an ably_api_key resource. The label seeds the key's name.
func APIKey(label string, opts ...APIKeyOptions) *tfgen.Resource {
	o := first(opts)

	config := tfgen.Config{"name": label}
	if o.Name != "" {
		config["name"] = o.Name
	}
	if o.AppID != "" {
		config["app_id"] = o.AppID
	}
	if o.Capabilities != nil {
		config["capabilities"] = o.Capabilities
	}
	if o.RevocableTokens {
		config["revocable_tokens"] = true
	}

	return &tfgen.Resource{Type: "ably_api_key", Name: label, Config: config}
}

// NamespaceOptions configures an ably_namespace resource. AppID is required by
// Terraform. This covers the commonly used attributes; anything else can be set
// on the returned resource's Config.
type NamespaceOptions struct {
	ID               string // defaults to the resource label
	AppID            string // required
	Persisted        bool
	PersistLast      bool
	PushEnabled      bool
	TLSOnly          bool
	ExposeTimeserial bool
	BatchingEnabled  bool
	BatchingInterval int
}

// Namespace creates an ably_namespace resource. The label seeds the namespace
// id.
func Namespace(label string, opts ...NamespaceOptions) *tfgen.Resource {
	o := first(opts)

	config := tfgen.Config{"id": label}
	if o.ID != "" {
		config["id"] = o.ID
	}
	if o.AppID != "" {
		config["app_id"] = o.AppID
	}
	if o.Persisted {
		config["persisted"] = true
	}
	if o.PersistLast {
		config["persist_last"] = true
	}
	if o.PushEnabled {
		config["push_enabled"] = true
	}
	if o.TLSOnly {
		config["tls_only"] = true
	}
	if o.ExposeTimeserial {
		config["expose_timeserial"] = true
	}
	if o.BatchingEnabled {
		config["batching_enabled"] = true
	}
	if o.BatchingInterval != 0 {
		config["batching_interval"] = o.BatchingInterval
	}

	return &tfgen.Resource{Type: "ably_namespace", Name: label, Config: config}
}

// QueueOptions configures an ably_queue resource. AppID, Region, MaxLength, and
// TTL are all required by Terraform.
type QueueOptions struct {
	Name      string // defaults to the resource label
	AppID     string // required
	Region    string // required
	MaxLength int    // required
	TTL       int    // required
}

// Queue creates an ably_queue resource. The label seeds the queue's name.
func Queue(label string, opts ...QueueOptions) *tfgen.Resource {
	o := first(opts)

	config := tfgen.Config{"name": label}
	if o.Name != "" {
		config["name"] = o.Name
	}
	if o.AppID != "" {
		config["app_id"] = o.AppID
	}
	if o.Region != "" {
		config["region"] = o.Region
	}
	if o.MaxLength != 0 {
		config["max_length"] = o.MaxLength
	}
	if o.TTL != 0 {
		config["ttl"] = o.TTL
	}

	return &tfgen.Resource{Type: "ably_queue", Name: label, Config: config}
}

// first returns the first options value, or the zero value if none were given.
func first[T any](opts []T) (o T) {
	if len(opts) > 0 {
		o = opts[0]
	}
	return
}

// firstOr returns the first value in s, or def if s is empty.
func firstOr[T any](s []T, def T) T {
	if len(s) > 0 {
		return s[0]
	}
	return def
}
