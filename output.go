package tfgen

// Outputs is a collection of Output configurations.
type Outputs []*Output

func (os Outputs) generate() map[string]any {
	if len(os) == 0 {
		return nil
	}

	out := make(map[string]any, len(os))
	for _, o := range os {
		out[o.Name] = o.generate()
	}

	return out
}

// Output represents an output block in a Terraform configuration.
type Output struct {
	Name        string // output name
	Value       any    // exported value, often a Ref to another block
	Description string // optional human-readable description
}

func (o *Output) generate() map[string]any {
	out := map[string]any{
		"value": o.Value,
	}

	if o.Description != "" {
		out["description"] = o.Description
	}

	return out
}
