package tfgen

import "fmt"

// Resources is a collection of Resource configurations.
type Resources []*Resource

func (rs Resources) generate() map[string]map[string]any {
	if len(rs) == 0 {
		return nil
	}

	out := make(map[string]map[string]any)
	for _, r := range rs {
		if out[r.Type] == nil {
			out[r.Type] = make(map[string]any)
		}
		out[r.Type][r.Name] = r.generate()
	}

	return out
}

// Resource represents a resource block in a Terraform configuration.
type Resource struct {
	Type   string // resource type, e.g. "aws_s3_bucket"
	Name   string // local name, unique per type within the stack
	Config Config // the resource's arguments
}

// Ref returns a Terraform expression referencing an attribute of this
// resource, e.g. "${aws_s3_bucket.example.arn}".
func (r *Resource) Ref(attribute string) string {
	return fmt.Sprintf("${%s.%s.%s}", r.Type, r.Name, attribute)
}

func (r *Resource) generate() map[string]any {
	return r.Config
}
