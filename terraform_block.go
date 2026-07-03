package tfgen

// TerraformBlock represents the terraform block in a Terraform configuration, which includes required_version, required_providers, and backend configuration.
type TerraformBlock struct {
	RequiredVersion   string
	RequiredProviders RequiredProviders
	Backend           *Backend
}

// RequiredProviders maps a provider's local name to its source and version
// requirement.
type RequiredProviders map[string]ProviderRequirement

// AddRequiredProvider records a provider requirement under the given local
// name, allocating the map if necessary.
func (block *TerraformBlock) AddRequiredProvider(name string, req ProviderRequirement) {
	if block.RequiredProviders == nil {
		block.RequiredProviders = make(RequiredProviders)
	}
	block.RequiredProviders[name] = req
}

// ProviderRequirement represents the required provider configuration in the terraform block, which includes the source and version of the provider.
type ProviderRequirement struct {
	Source  string
	Version string
}

// Backend represents the backend configuration in the terraform block, which includes the type of backend and its configuration.
type Backend struct {
	Type   string
	Config Config
}

func (block *TerraformBlock) generate() map[string]any {
	tfblock := make(map[string]any)

	if block.RequiredVersion != "" {
		tfblock["required_version"] = block.RequiredVersion
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
