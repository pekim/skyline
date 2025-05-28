# skyline

skyline implements the skyline algorithm for packing 2D rectangles.

## jvernay's code

Most of the implementation is a Go port of
https://git.sr.ht/~jvernay/JV/tree/main/item/src/jv_pack2d.
The article at
https://jvernay.fr/en/blog/skyline-2d-packer/implementation/.
describes the code.

That code was placed in the [public domain](https://unlicense.org/) by jvernay.

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
