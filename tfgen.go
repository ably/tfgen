package tfgen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Generator interface {
	// Validate checks that the generator's configuration is valid. This should
	// be called before GenerateStacks to catch config errors early.
	Validate() error

	// GenerateStacks generates one or more Terraform stacks based on the
	// generator's configuration. The generated stacks can then be written to
	// files or applied using a Terraform-compatible CLI.
	GenerateStacks() (Stacks, error)
}

type Stacks []*Stack

func (s Stacks) Write() error {
	for _, stack := range s {
		json, err := stack.JSON()
		if err != nil {
			return err
		}

		if _, err := stack.Output.Write(json); err != nil {
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

	// Output is where the generated Terraform configuration should be written
	// to.
	Output io.Writer

	config TerraformConfig
}

// NewStack creates a new Stack with the given name and an empty Terraform
// configuration.
func NewStack(name string, output io.Writer) *Stack {
	return &Stack{Name: name, Output: output}
}

func NewRequiredVersion(version string) *string {
	return &version
}

func (s *Stack) SetTerraformBlock(block *TerraformBlock) {
	s.config.TerraformBlock = block
}

func (s *Stack) SetProviders(providers []*Provider) {
	s.config.Providers = providers
}

func (s *Stack) AddProvider(provider *Provider) {
	s.config.Providers = append(s.config.Providers, provider)
}

func (s *Stack) SetResources(resources []*Resource) {
	s.config.Resources = resources
}

func (s *Stack) AddResource(resource *Resource) {
	s.config.Resources = append(s.config.Resources, resource)
}

func (s *Stack) SetData(data []*Data) {
	s.config.Data = data
}

func (s *Stack) AddData(data *Data) {
	s.config.Data = append(s.config.Data, data)
}

func (s *Stack) SetModules(modules []*Module) {
	s.config.Modules = modules
}

func (s *Stack) AddModule(module *Module) {
	s.config.Modules = append(s.config.Modules, module)
}

func (s *Stack) SetOutputs(outputs []*Output) {
	s.config.Outputs = outputs
}

func (s *Stack) AddOutput(output *Output) {
	s.config.Outputs = append(s.config.Outputs, output)
}

func (s *Stack) SetVariables(variables []*Variable) {
	s.config.Variables = variables
}

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
