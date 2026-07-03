package tfgen_test

import (
	"bytes"
	"fmt"

	"github.com/ably/tfgen"
	"github.com/ably/tfgen/providers"
)

// Example builds a small stack in memory and writes it as Terraform JSON.
func Example() {
	var buf bytes.Buffer
	stack := tfgen.NewStack("example", &buf)

	stack.SetTerraformBlock(&tfgen.TerraformBlock{
		RequiredVersion: tfgen.NewRequiredVersion(">= 1.9"),
		RequiredProviders: map[string]tfgen.ProviderRequirement{
			"aws": {Source: "hashicorp/aws", Version: "~> 5.0"},
		},
	})

	stack.AddProvider(providers.AWS("primary", "eu-west-1", map[string]any{
		"managed_by": "tfgen",
	}))

	bucket := &tfgen.Resource{
		Type: "aws_s3_bucket",
		Name: "example",
		Config: map[string]any{
			"bucket": "my-example-bucket",
		},
	}
	stack.AddResource(bucket)

	// Reference the resource's ARN in an output.
	stack.AddOutput(&tfgen.Output{
		Name:  "bucket_arn",
		Value: bucket.Ref("arn"),
	})

	if err := (tfgen.Stacks{stack}).Write(); err != nil {
		panic(err)
	}

	fmt.Print(buf.String())
	// Output:
	// {
	//   "output": {
	//     "bucket_arn": {
	//       "value": "${aws_s3_bucket.example.arn}"
	//     }
	//   },
	//   "provider": {
	//     "aws": {
	//       "alias": "primary",
	//       "default_tags": [
	//         {
	//           "tags": {
	//             "managed_by": "tfgen"
	//           }
	//         }
	//       ],
	//       "region": "eu-west-1"
	//     }
	//   },
	//   "resource": {
	//     "aws_s3_bucket": {
	//       "example": {
	//         "bucket": "my-example-bucket"
	//       }
	//     }
	//   },
	//   "terraform": {
	//     "required_providers": {
	//       "aws": {
	//         "source": "hashicorp/aws",
	//         "version": "~> 5.0"
	//       }
	//     },
	//     "required_version": ">= 1.9"
	//   }
	// }
}
