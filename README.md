# typegroupingcheck Linter

`typegroupingcheck` is a static analysis tool for Go that checks for grouped function parameter types in Go code. It is designed to be integrated with the `golangci-lint` toolset.

## Overview

The linter checks for grouped parameters, where the type is specified once for multiple parameters.

It can be used as a standalone tool or integrated into `golangci-lint` to enhance the existing suite of linters with this specific check.

## Installation

To install `typegroupingcheck` as a standalone linter, use the following command:

```sh
go install github.com/karimodm/typegroupingcheck@latest
```

Replace `github.com/karimodm/typegroupingcheck` with the actual import path of your linter.

## Building

To build `typegroupingcheck` as a plugin ready to be used by golangci-list, use the following command:

```sh
go build -buildmode=plugin -o typegroupingcheck.so
```

## Usage

After installation, you can run `typegroupingcheck` on your Go files or projects like so:

```sh
typegroupingcheck ./...
```

## Integration with `golangci-lint`

To integrate `typegroupingcheck` with `golangci-lint`, you will need to add it to the configuration file `.golangci.yml` in the root directory of your project. Here is a sample configuration:

```yaml
linters-settings:
  custom:
    typegroupingcheck:
      path: typegroupingcheck
      description: Checks for grouped function parameter types
      original-url: github.com/karimodm/typegroupingcheck
linters:
  enable:
    - typegroupingcheck
```

Please note that integration with `golangci-lint` requires the `typegroupingcheck` binary to be in your `PATH`.

## Contributing

Contributions to `typegroupingcheck` are welcome! You can contribute by:
- Reporting issues
- Suggesting new features or enhancements
- Submitting pull requests to improve the linter

Please refer to [CONTRIBUTING.md](CONTRIBUTING.md) for more details on how to contribute.

## License

`typegroupingcheck` is distributed under the [MIT License](LICENSE).
