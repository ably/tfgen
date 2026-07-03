package tfgen

// TerraformConfig represents the entire Terraform configuration, including the
// terraform block, providers, resources, data sources, modules, variables, and
// outputs.
type TerraformConfig struct {
	Data           DataSources
	Modules        Modules
	Outputs        Outputs
	Providers      Providers
	Resources      Resources
	TerraformBlock *TerraformBlock
	Variables      Variables
}
