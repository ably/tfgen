package tfgen

// Config is a free-form set of Terraform configuration attributes, used for the
// body of resources, data sources, modules, providers, and backends. The keys
// and values map directly onto Terraform's JSON syntax.
//
// Its underlying type is map[string]any, so a plain map literal works anywhere
// a Config is expected, and vice versa.
type Config map[string]any
