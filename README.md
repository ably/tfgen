# tfgen

[![Go Reference](https://pkg.go.dev/badge/github.com/ably/tfgen.svg)](https://pkg.go.dev/github.com/ably/tfgen)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)

Generate Terraform configuration from Go.

`tfgen` is a small library for building [Terraform JSON
configuration](https://developer.hashicorp.com/terraform/language/syntax/json)
programmatically in Go. You describe your infrastructure with plain Go values
and `tfgen` emits `*.tf.json` that any Terraform-compatible CLI
([Terraform](https://developer.hashicorp.com/terraform) or
[OpenTofu](https://opentofu.org/)) can apply.

Unlike Hashicorp's now unsupported
[CDKTF](https://developer.hashicorp.com/terraform/cdktf), configuration is
entirely flexible and does not depend on generating strong types based off
provider versions. This means you can reliably and quickly generate
Terraform-compatible JSON using easy to understand Go.

## Quick start

Keep a small generator in your project and wire it to `go generate`:

```go
package main

//go:generate go run .

import (
	"log"
	"os"

	"github.com/ably/tfgen"
)

func main() {
	stack := tfgen.NewStack("example")

	bucket := &tfgen.Resource{
		Type: "aws_s3_bucket",
		Name: "example",
		Config: tfgen.Config{
			"bucket": "my-example-bucket",
		},
	}
	stack.AddResource(bucket)

	// Ref builds a Terraform interpolation referencing an attribute of the
	// resource, so you never hand-write "${...}" strings.
	stack.AddOutput(&tfgen.Output{
		Name:  "bucket_arn",
		Value: bucket.Ref("arn"),
	})

	f, err := os.Create("main.tf.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := stack.Write(f); err != nil {
		log.Fatal(err)
	}
}
```

Run `go generate ./...` to write `main.tf.json`, ready to apply with
`terraform` or `tofu`:

```json
{
  "output": {
    "bucket_arn": {
      "value": "${aws_s3_bucket.example.arn}"
    }
  },
  "resource": {
    "aws_s3_bucket": {
      "example": {
        "bucket": "my-example-bucket"
      }
    }
  }
}
```

See [`examples/generate`](examples/generate) for a complete, runnable example.

## Concepts

A `Stack` is one Terraform configuration. You add building blocks to it, a
`Resource`, `Data` source, `Module`, `Provider`, `Output`, `Variable`, or the
`terraform` block, then serialise it to JSON. `Config` is a `map[string]any`, so
any attribute maps directly onto Terraform's JSON syntax, and helpers such as
`Resource.Ref` and `Variable.Ref` build interpolation references so you never
hand-write `"${...}"` strings.

Write a single stack to any `io.Writer` with `Stack.Write`, or write a set of
stacks to their conventional files at `{baseDir}/stacks/{name}/main.tf.json`
with `Stacks.Write`. See the
[Go documentation](https://pkg.go.dev/github.com/ably/tfgen) for the full API.

## Provider helpers

The `providers/*` subpackages provide optional, typed constructors for common
providers. They are thin sugar over `tfgen` types, so you can always drop down
to a plain `tfgen.Provider` or `tfgen.Resource`.

- [`providers/aws`](providers/aws) — AWS provider
- [`providers/cloudflare`](providers/cloudflare) — Cloudflare provider
- [`providers/ably`](providers/ably) — the [Ably Terraform provider](https://github.com/ably/terraform-provider-ably): apps, API keys, namespaces, and queues

## Contributing

Contributions are welcome, see [CONTRIBUTING.md](CONTRIBUTING.md).

## License

Licensed under the [Apache License 2.0](LICENSE).
