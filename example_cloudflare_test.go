package tfgen_test

import (
	"bytes"
	"fmt"

	"github.com/ably/tfgen"
	"github.com/ably/tfgen/providers/cloudflare"
)

// Example_cloudflare builds a stack using the Cloudflare provider helper.
func Example_cloudflare() {
	stack := tfgen.NewStack("cloudflare")

	block := &tfgen.TerraformBlock{RequiredVersion: ">= 1.9"}
	block.AddRequiredProvider("cloudflare", cloudflare.ProviderVersion("~> 4.0"))
	stack.SetTerraformBlock(block)

	stack.AddProvider(cloudflare.Provider())

	stack.AddResource(&tfgen.Resource{
		Type: "cloudflare_record",
		Name: "www",
		Config: tfgen.Config{
			"zone_id": "${var.zone_id}",
			"name":    "www",
			"type":    "CNAME",
			"content": "example.com",
			"proxied": true,
		},
	})

	var buf bytes.Buffer
	if err := stack.Write(&buf); err != nil {
		panic(err)
	}

	fmt.Print(buf.String())
	// Output:
	// {
	//   "provider": {
	//     "cloudflare": {}
	//   },
	//   "resource": {
	//     "cloudflare_record": {
	//       "www": {
	//         "content": "example.com",
	//         "name": "www",
	//         "proxied": true,
	//         "type": "CNAME",
	//         "zone_id": "${var.zone_id}"
	//       }
	//     }
	//   },
	//   "terraform": {
	//     "required_providers": {
	//       "cloudflare": {
	//         "source": "cloudflare/cloudflare",
	//         "version": "~> 4.0"
	//       }
	//     },
	//     "required_version": ">= 1.9"
	//   }
	// }
}
