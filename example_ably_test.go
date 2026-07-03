package tfgen_test

import (
	"bytes"
	"fmt"

	"github.com/ably/tfgen"
	"github.com/ably/tfgen/providers/ably"
)

// Example_ably provisions an Ably app with a namespace and an API key, wiring
// the child resources to the app via references.
func Example_ably() {
	stack := tfgen.NewStack("ably")

	block := &tfgen.TerraformBlock{RequiredVersion: ">= 1.9"}
	block.AddRequiredProvider("ably", ably.ProviderVersion())
	stack.SetTerraformBlock(block)

	stack.AddProvider(ably.Provider())

	app := ably.App("main", ably.AppOptions{TLSOnly: true})
	stack.AddResource(app)

	stack.AddResource(ably.Namespace("chat", ably.NamespaceOptions{
		AppID:     app.Ref("id"),
		Persisted: true,
	}))

	stack.AddResource(ably.APIKey("publisher", ably.APIKeyOptions{
		AppID:        app.Ref("id"),
		Capabilities: map[string][]string{"chat:*": {"publish", "subscribe"}},
	}))

	var buf bytes.Buffer
	if err := stack.Write(&buf); err != nil {
		panic(err)
	}

	fmt.Print(buf.String())
	// Output:
	// {
	//   "provider": {
	//     "ably": {}
	//   },
	//   "resource": {
	//     "ably_api_key": {
	//       "publisher": {
	//         "app_id": "${ably_app.main.id}",
	//         "capabilities": {
	//           "chat:*": [
	//             "publish",
	//             "subscribe"
	//           ]
	//         },
	//         "name": "publisher"
	//       }
	//     },
	//     "ably_app": {
	//       "main": {
	//         "name": "main",
	//         "tls_only": true
	//       }
	//     },
	//     "ably_namespace": {
	//       "chat": {
	//         "app_id": "${ably_app.main.id}",
	//         "id": "chat",
	//         "persisted": true
	//       }
	//     }
	//   },
	//   "terraform": {
	//     "required_providers": {
	//       "ably": {
	//         "source": "ably/ably",
	//         "version": "1.0.0"
	//       }
	//     },
	//     "required_version": ">= 1.9"
	//   }
	// }
}
