// Package providers offers convenience constructors for commonly used
// Terraform providers. They are optional sugar over tfgen.Provider; you can
// always construct a *tfgen.Provider directly.
package providers

import "github.com/ably/tfgen"

// AWS creates an AWS provider with the given alias, region, and default tags.
func AWS(alias, region string, tags map[string]any) *tfgen.Provider {
	return &tfgen.Provider{
		Name: "aws",
		Config: map[string]any{
			"alias":        alias,
			"region":       region,
			"default_tags": []map[string]any{{"tags": tags}},
		},
	}
}

// AWSWithAssumeRole creates an AWS provider that assumes the given role.
func AWSWithAssumeRole(alias, region, roleARN string, tags map[string]any) *tfgen.Provider {
	return &tfgen.Provider{
		Name: "aws",
		Config: map[string]any{
			"alias":        alias,
			"region":       region,
			"default_tags": []map[string]any{{"tags": tags}},
			"assume_role":  AssumeRole(roleARN),
		},
	}
}

// Cloudflare creates a Cloudflare provider.
func Cloudflare() *tfgen.Provider {
	return &tfgen.Provider{
		Name:   "cloudflare",
		Config: map[string]any{},
	}
}

// AssumeRole returns a Terraform assume_role block for the given role ARN.
func AssumeRole(roleARN string) []map[string]any {
	return []map[string]any{
		{"role_arn": roleARN},
	}
}
