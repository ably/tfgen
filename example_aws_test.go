package tfgen_test

import (
	"bytes"
	"fmt"

	"github.com/ably/tfgen"
	"github.com/ably/tfgen/providers/aws"
)

// Example_aws builds a stack using the AWS provider helpers.
func Example_aws() {
	stack := tfgen.NewStack("aws")

	block := &tfgen.TerraformBlock{RequiredVersion: ">= 1.9"}
	block.AddRequiredProvider("aws", aws.ProviderVersion("~> 5.0"))
	stack.SetTerraformBlock(block)

	stack.AddProvider(aws.Provider("primary", "eu-west-1", tfgen.Config{
		"managed_by": "tfgen",
	}))

	stack.AddResource(&tfgen.Resource{
		Type:   "aws_s3_bucket",
		Name:   "example",
		Config: tfgen.Config{"bucket": "my-example-bucket"},
	})

	var buf bytes.Buffer
	if err := stack.Write(&buf); err != nil {
		panic(err)
	}

	fmt.Print(buf.String())
	// Output:
	// {
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
