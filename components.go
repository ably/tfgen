package tfgen

// TerraformConfig holds the building blocks of a single Terraform
// configuration: the terraform block, providers, resources, data sources,
// modules, variables, and outputs.
type TerraformConfig struct {
	Data           DataSources     // data source blocks
	Modules        Modules         // module blocks
	Outputs        Outputs         // output blocks
	Providers      Providers       // provider blocks
	Resources      Resources       // resource blocks
	TerraformBlock *TerraformBlock // the terraform block, if any
	Variables      Variables       // variable blocks
}
