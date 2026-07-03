package tfgen

import (
	"fmt"
	"maps"
	"path/filepath"
)

// Modules is a collection of Module configurations.
type Modules []*Module

func (ms Modules) generate() map[string]any {
	if len(ms) == 0 {
		return nil
	}

	out := make(map[string]any, len(ms))
	for _, m := range ms {
		out[m.Name] = m.generate()
	}

	return out
}

// Module represents a module block in a Terraform configuration.
type Module struct {
	Name    string // local name of the module instance
	Source  string // module source address
	Version string // optional version constraint (registry modules only)
	Config  Config // input variables passed to the module
}

// ModuleSource derives a module's Source from a consistent naming convention,
// whether the module lives in a registry or on the local filesystem. It follows
// Terraform's terraform-<PROVIDER>-<NAME> convention:
//
//   - Registry: {registryPrefix}/{name}/{provider}
//   - Local:    {modulesDir}/{provider}/terraform-{provider}-{name}
//
// SubDir inserts an additional directory between the provider and the module
// directory when computing the local path:
//
//   - Local with SubDir: {modulesDir}/{provider}/{subDir}/terraform-{provider}-{name}
//
// Using ModuleSource is optional; you can always set Module.Source directly.
type ModuleSource struct {
	Name     string // module name, e.g. "vpc", "eks-cluster"
	Provider string // provider namespace, e.g. "aws", "google"
	SubDir   string // optional subdirectory within {provider}/
}

// RegistrySource returns the registry source path for the given registry
// prefix. The prefix is the leading portion of a registry module address, such
// as the host and namespace (e.g. "app.terraform.io/example-org").
func (ms ModuleSource) RegistrySource(registryPrefix string) string {
	return fmt.Sprintf("%s/%s/%s", registryPrefix, ms.Name, ms.Provider)
}

// LocalSource returns the local module path rooted at modulesDir.
func (ms ModuleSource) LocalSource(modulesDir string) string {
	dir := fmt.Sprintf("terraform-%s-%s", ms.Provider, ms.Name)
	if ms.SubDir != "" {
		return filepath.Join(modulesDir, ms.Provider, ms.SubDir, dir)
	}
	return filepath.Join(modulesDir, ms.Provider, dir)
}

// OutputRef returns a Terraform expression referencing an output of this module,
// e.g. "${module.my_module.vpc_id}".
func (m *Module) OutputRef(output string) string {
	return fmt.Sprintf("${module.%s.%s}", m.Name, output)
}

// Outputs creates Output blocks for each named output, referencing this module.
func (m *Module) Outputs(names []string) []*Output {
	outputs := make([]*Output, len(names))
	for i, name := range names {
		outputs[i] = &Output{
			Name:  name,
			Value: m.OutputRef(name),
		}
	}
	return outputs
}

func (m *Module) generate() map[string]any {
	out := map[string]any{
		"source": m.Source,
	}

	if m.Version != "" {
		out["version"] = m.Version
	}

	maps.Copy(out, m.Config)

	return out
}
