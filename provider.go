package tfgen

// Providers is a collection of Provider configurations.
type Providers []*Provider

// generate groups providers by name. When multiple providers share the same
// name (aliases), they are grouped into an array per the Terraform JSON
// specification.
func (ps Providers) generate() map[string]any {
	if len(ps) == 0 {
		return nil
	}

	grouped := make(map[string][]map[string]any)
	for _, p := range ps {
		grouped[p.Name] = append(grouped[p.Name], p.generate())
	}

	out := make(map[string]any, len(grouped))
	for name, configs := range grouped {
		if len(configs) == 1 {
			out[name] = configs[0]
		} else {
			out[name] = configs
		}
	}

	return out
}

// Provider represents a provider block in a Terraform configuration.
type Provider struct {
	Name   string
	Config map[string]any // includes "alias" if needed
}

func (p *Provider) generate() map[string]any {
	return p.Config
}
