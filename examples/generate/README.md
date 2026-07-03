# go generate example

A minimal, runnable example of using `tfgen` with `go generate` to produce
Terraform JSON from Go.

```sh
go generate ./...
```

This runs the generator in [`main.go`](main.go), which writes
`stacks/ably/main.tf.json`. Apply it with your usual tooling:

```sh
cd stacks/ably
terraform init && terraform apply   # or: tofu init && tofu apply
```

In a real project you would keep a generator like this in your repository,
commit the generated `*.tf.json`, and re-run `go generate ./...` whenever the
Go configuration changes. The generated output is git-ignored here to keep the
example repository clean.
