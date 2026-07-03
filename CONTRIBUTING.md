# Contributing to tfgen

Thanks for your interest in contributing. This document explains how to set up
your environment, the conventions the project follows, and how releases are cut.

## Reporting issues

Please raise bugs and feature requests as [GitHub issues](https://github.com/ably/tfgen/issues).
For bugs, include the version of `tfgen`, your Go version, and a minimal example
that reproduces the problem, ideally the input configuration and the Terraform
JSON you got versus what you expected.

## Development setup

`tfgen` is a standard Go module with no external toolchain beyond Go itself.

1. Fork the repository and clone your fork.
2. Make sure you have a supported Go version installed (see the `go` directive
   in [`go.mod`](go.mod); CI runs against the two most recent Go releases).
3. Create a branch for your change:

   ```sh
   git checkout -b my-change
   ```

## Running the checks locally

The same checks run in CI. Run them before opening a pull request:

```sh
gofmt -l .        # should print nothing
go vet ./...
go build ./...
go test -race ./...
```

If you have [golangci-lint](https://golangci-lint.run/) installed, you can also
run `golangci-lint run`, which CI runs on every pull request.

## Conventions

- **Formatting.** All code must be `gofmt`-ed. CI fails on any unformatted file.
- **Documentation.** Every exported type, function, and method should carry a
  doc comment describing its contract. Comments explain the non-obvious *why*,
  not the *what*.
- **Tests.** New behaviour needs test coverage, and every bug fix should come
  with a regression test. Tests live alongside the code in `_test.go` files.
- **Runnable examples.** Where it aids understanding, add
  [testable examples](https://go.dev/blog/examples) so they appear on
  pkg.go.dev and are verified by `go test`.
- **Commits and pull requests.** Keep each pull request focused on a single
  logical change. Write a clear description of the problem you are solving and
  the approach you took.

## Changelog

Add a note to the `Unreleased` section of [`CHANGELOG.md`](CHANGELOG.md) for any
user-visible change (new features, behaviour changes, bug fixes, deprecations).

## Releasing

Releases are cut by maintainers. `tfgen` follows
[Semantic Versioning](https://semver.org/). Because it is a Go module, a release
is simply an annotated, signed Git tag of the form `vX.Y.Z`; pkg.go.dev and
`go get` pick up the tag automatically, so there is no artefact to publish.

1. Decide the new version from the changes since the last release: patch for
   bug fixes, minor for backwards-compatible additions, major for breaking
   changes. While the module is pre-1.0, breaking changes bump the minor
   version.
2. Update [`CHANGELOG.md`](CHANGELOG.md): rename the `Unreleased` section to the
   new version with today's date, and add the comparison link at the bottom.
3. Open a pull request with those changes and get it merged to `main`.
4. Tag the merge commit and push the tag:

   ```sh
   git checkout main
   git pull
   git tag -s vX.Y.Z -m "vX.Y.Z"
   git push origin vX.Y.Z
   ```

5. Create a [GitHub release](https://github.com/ably/tfgen/releases) from the
   tag, using the changelog entry as the release notes.

## Licence

By contributing, you agree that your contributions will be licensed under the
[Apache 2.0 licence](LICENSE) that covers this project.
