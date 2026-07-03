// Package aws provides convenience constructors for the Terraform AWS provider.
// They are optional sugar over tfgen types; you can always construct them
// directly.
package aws

import "github.com/ably/tfgen"

const source = "hashicorp/aws"

// ProviderVersion returns a ProviderRequirement for the AWS provider pinned to
// the given version constraint, e.g. "~> 5.0".
func ProviderVersion(version string) tfgen.ProviderRequirement {
	return tfgen.ProviderRequirement{Source: source, Version: version}
}

// Provider creates an AWS provider with the given alias, region, and default
// tags.
func Provider(alias, region string, tags tfgen.Config) *tfgen.Provider {
	return &tfgen.Provider{
		Name: "aws",
		Config: tfgen.Config{
			"alias":        alias,
			"region":       region,
			"default_tags": []tfgen.Config{{"tags": tags}},
		},
	}
}

// ProviderWithAssumeRole creates an AWS provider that assumes the given role.
func ProviderWithAssumeRole(alias, region, roleARN string, tags tfgen.Config) *tfgen.Provider {
	return &tfgen.Provider{
		Name: "aws",
		Config: tfgen.Config{
			"alias":        alias,
			"region":       region,
			"default_tags": []tfgen.Config{{"tags": tags}},
			"assume_role":  AssumeRole(roleARN),
		},
	}
}

// AssumeRole returns a Terraform assume_role block for the given role ARN.
func AssumeRole(roleARN string) []tfgen.Config {
	return []tfgen.Config{
		{"role_arn": roleARN},
	}
}
