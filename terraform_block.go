package tfgen

// TerraformBlock represents the terraform block in a Terraform configuration, which includes required_version, required_providers, and backend configuration.
type TerraformBlock struct {
	RequiredVersion   *string
	RequiredProviders map[string]ProviderRequirement
	Backend           *Backend
}

// ProviderRequirement represents the required provider configuration in the terraform block, which includes the source and version of the provider.
type ProviderRequirement struct {
	Source  string
	Version string
}

// Backend represents the backend configuration in the terraform block, which includes the type of backend and its configuration.
type Backend struct {
	Type   string
	Config map[string]any
}

func (block *TerraformBlock) generate() map[string]any {
	tfblock := make(map[string]any)

	if rv := block.RequiredVersion; rv != nil {
		tfblock["required_version"] = rv
	}

	if backend := block.Backend; backend != nil {
		tfblock["backend"] = map[string]any{
			backend.Type: backend.Config,
		}
	}

	for name, config := range block.RequiredProviders {
		if tfblock["required_providers"] == nil {
			tfblock["required_providers"] = make(map[string]any)
		}

		tfblock["required_providers"].(map[string]any)[name] = map[string]any{
			"source":  config.Source,
			"version": config.Version,
		}
	}

	return tfblock
}
