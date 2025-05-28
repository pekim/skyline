# skyline

[![PkgGoDev](https://pkg.go.dev/badge/github.com/pekim/skyline)](https://pkg.go.dev/github.com/pekim/skyline)
[![golangci-lint](https://github.com/pekim/skyline/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/pekim/skyline/actions/workflows/golangci-lint.yml)

skyline implements the skyline algorithm for packing 2D rectangles.

## Acknowledgment

Most of the implementation is a Go port of jvernay's
[jv_pack2d](https://git.sr.ht/~jvernay/JV/tree/main/item/src/jv_pack2d).
There is a corresponding
[implementation article](https://jvernay.fr/en/blog/skyline-2d-packer/implementation/)
that describes the code.

That code was placed in the [public domain](https://unlicense.org/) by jvernay.

## License

skyline is licensed under the terms of the MIT license.

## Development

### pre-commit hook

There are configuration files for linting and other checks.
To use a git pre-commit hook for the checks

- install `goimports` if not already installed
  - https://pkg.go.dev/golang.org/x/tools/cmd/goimports
- install `golangci-lint` if not already installed
  - https://golangci-lint.run/usage/install/#local-installation
- install the `pre-commit` application if not already installed
  - https://pre-commit.com/index.html#install
- install a git pre-commit hook in this repo's workspace
  - `pre-commit install`
