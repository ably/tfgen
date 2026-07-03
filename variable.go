package tfgen

import "fmt"

// Variables is a collection of Variable configurations.
type Variables []*Variable

func (vs Variables) generate() map[string]any {
	if len(vs) == 0 {
		return nil
	}

	out := make(map[string]any, len(vs))
	for _, v := range vs {
		out[v.Name] = v.generate()
	}

	return out
}

// Variable represents a variable block in a Terraform configuration.
type Variable struct {
	Name        string
	Type        string // e.g. "string", "bool", "map(string)"
	Description string
	Default     any // nil means required (no default key emitted)
}

// Ref returns a Terraform expression referencing this variable,
// e.g. "${var.name_prefix}".
func (v *Variable) Ref() string {
	return fmt.Sprintf("${var.%s}", v.Name)
}

// CountRef returns a Terraform conditional count expression based on this
// boolean variable, e.g. "${var.deploy_ipam ? 1 : 0}".
func (v *Variable) CountRef() string {
	return fmt.Sprintf("${var.%s ? 1 : 0}", v.Name)
}

func (v *Variable) generate() map[string]any {
	out := make(map[string]any)

	if v.Type != "" {
		out["type"] = v.Type
	}

	if v.Description != "" {
		out["description"] = v.Description
	}

	if v.Default != nil {
		out["default"] = v.Default
	}

	return out
}
