// Package tfgen generates Terraform JSON configuration from Go.
//
// You describe infrastructure with plain Go values. A [Stack] holds one
// Terraform configuration, to which you add resources, data sources, modules,
// providers, outputs, and variables. [Stack.Write] then serialises it to
// Terraform JSON that any Terraform-compatible CLI (Terraform or OpenTofu) can
// apply.
//
// The providers subpackages offer optional, typed constructors for common
// providers.
package tfgen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Generator produces one or more Terraform stacks from some configuration.
// Implement it to organise larger configuration-generation logic behind a
// single, testable entry point.
type Generator interface {
	// Validate checks that the generator's configuration is valid. This should
	// be called before GenerateStacks to catch config errors early.
	Validate() error

	// GenerateStacks generates one or more Terraform stacks based on the
	// generator's configuration. The generated stacks can then be written to
	// files or applied using a Terraform-compatible CLI.
	GenerateStacks() (Stacks, error)
}

// Stacks is a collection of Stack values.
type Stacks []*Stack

// Write writes each stack to its conventional file at
// {baseDir}/stacks/{name}/main.tf.json, creating directories as needed.
func (s Stacks) Write(baseDir string) error {
	for _, stack := range s {
		f, err := StackOutputWriter(baseDir, stack.Name)
		if err != nil {
			return err
		}

		if err := stack.Write(f); err != nil {
			f.Close()
			return err
		}

		if err := f.Close(); err != nil {
			return err
		}
	}

	return nil
}

// Stack represents a Terraform stack which can eventually be written to be
// applied by a Terraform-compatible CLI, such as Terraform or Tofu.
type Stack struct {
	// Name is the name of the stack
	Name string

	config TerraformConfig
}

// NewStack creates a new Stack with the given name and an empty Terraform
// configuration.
func NewStack(name string) *Stack {
	return &Stack{Name: name}
}

// Write writes the stack's Terraform JSON configuration to w.
func (s *Stack) Write(w io.Writer) error {
	data, err := s.JSON()
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

// SetTerraformBlock sets the stack's terraform block.
func (s *Stack) SetTerraformBlock(block *TerraformBlock) {
	s.config.TerraformBlock = block
}

// SetProviders replaces the stack's providers.
func (s *Stack) SetProviders(providers []*Provider) {
	s.config.Providers = providers
}

// AddProvider appends a provider to the stack.
func (s *Stack) AddProvider(provider *Provider) {
	s.config.Providers = append(s.config.Providers, provider)
}

// SetResources replaces the stack's resources.
func (s *Stack) SetResources(resources []*Resource) {
	s.config.Resources = resources
}

// AddResource appends a resource to the stack.
func (s *Stack) AddResource(resource *Resource) {
	s.config.Resources = append(s.config.Resources, resource)
}

// SetData replaces the stack's data sources.
func (s *Stack) SetData(data []*Data) {
	s.config.Data = data
}

// AddData appends a data source to the stack.
func (s *Stack) AddData(data *Data) {
	s.config.Data = append(s.config.Data, data)
}

// SetModules replaces the stack's modules.
func (s *Stack) SetModules(modules []*Module) {
	s.config.Modules = modules
}

// AddModule appends a module to the stack.
func (s *Stack) AddModule(module *Module) {
	s.config.Modules = append(s.config.Modules, module)
}

// SetOutputs replaces the stack's outputs.
func (s *Stack) SetOutputs(outputs []*Output) {
	s.config.Outputs = outputs
}

// AddOutput appends an output to the stack.
func (s *Stack) AddOutput(output *Output) {
	s.config.Outputs = append(s.config.Outputs, output)
}

// SetVariables replaces the stack's variables.
func (s *Stack) SetVariables(variables []*Variable) {
	s.config.Variables = variables
}

// AddVariable appends a variable to the stack.
func (s *Stack) AddVariable(variable *Variable) {
	s.config.Variables = append(s.config.Variables, variable)
}

// JSON returns the JSON representation of the stack's Terraform configuration.
// This correctly marshals the stack elements into the correct Terraform
// specification.
func (s *Stack) JSON() ([]byte, error) {
	stack := make(map[string]any)

	if block := s.config.TerraformBlock; block != nil {
		stack["terraform"] = block.generate()
	}

	if variables := s.config.Variables.generate(); len(variables) > 0 {
		stack["variable"] = variables
	}

	if providers := s.config.Providers.generate(); len(providers) > 0 {
		stack["provider"] = providers
	}

	if resources := s.config.Resources.generate(); len(resources) > 0 {
		stack["resource"] = resources
	}

	if data := s.config.Data.generate(); len(data) > 0 {
		stack["data"] = data
	}

	if modules := s.config.Modules.generate(); len(modules) > 0 {
		stack["module"] = modules
	}

	if outputs := s.config.Outputs.generate(); len(outputs) > 0 {
		stack["output"] = outputs
	}

	return MarshalJSON(stack)
}

// MarshalJSON encodes v to indented JSON with HTML escaping disabled.
func MarshalJSON(v any) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)

	if err := enc.Encode(v); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// StackOutputWriter creates the output directory and file for a stack. The
// file is written to {baseDir}/stacks/{stackName}/main.tf.json. For a custom
// filename or layout, use OutputFile directly.
func StackOutputWriter(baseDir, stackName string) (*os.File, error) {
	return OutputFile(filepath.Join(baseDir, "stacks", stackName, "main.tf.json"))
}

// OutputFile creates the parent directories and file at the given path.
func OutputFile(path string) (*os.File, error) {
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("creating output directory: %w", err)
	}

	return os.Create(path)
}
