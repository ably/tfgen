// Command generate shows how to drive tfgen from `go generate`. In a real
// project you keep a small generator like this in your repo, run
// `go generate ./...` to (re)produce the Terraform JSON, commit the generated
// *.tf.json, and apply it with terraform or tofu.
//
//go:generate go run .
package main

import (
	"log"

	"github.com/ably/tfgen"
	"github.com/ably/tfgen/providers/ably"
)

func main() {
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

	// Writes ./stacks/ably/main.tf.json, creating directories as needed.
	if err := (tfgen.Stacks{stack}).Write("."); err != nil {
		log.Fatalf("generating Terraform: %v", err)
	}

	log.Println("wrote stacks/ably/main.tf.json")
}
