package tfgen

import "fmt"

// DataSources is a collection of Data configurations.
type DataSources []*Data

func (ds DataSources) generate() map[string]map[string]any {
	if len(ds) == 0 {
		return nil
	}

	out := make(map[string]map[string]any)
	for _, d := range ds {
		if out[d.Type] == nil {
			out[d.Type] = make(map[string]any)
		}
		out[d.Type][d.Name] = d.generate()
	}

	return out
}

// Data represents a data source block in a Terraform configuration.
type Data struct {
	Type   string // data source type, e.g. "aws_ami"
	Name   string // local name, unique per type within the stack
	Config Config // the data source's arguments
}

// Ref returns a Terraform expression referencing an attribute of this data
// source, e.g. "${data.aws_acm_certificate.acm_certificate.arn}".
func (d *Data) Ref(attribute string) string {
	return fmt.Sprintf("${data.%s.%s.%s}", d.Type, d.Name, attribute)
}

func (d *Data) generate() map[string]any {
	return d.Config
}
