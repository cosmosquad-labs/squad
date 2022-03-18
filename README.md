[![codecov](https://codecov.io/gh/cosmosquad-labs/squad/branch/main/graph/badge.svg?token=07gYNeGo88)](https://codecov.io/gh/cosmosquad-labs/squad)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/cosmosquad-labs/squad)](https://pkg.go.dev/github.com/cosmosquad-labs/squad)

# Squad

The Squad containing below Cosmos SDK modules

- liquidity
- liquidstaking
- farming
- mint (constant inflation)
- claim

<!-- markdown-link-check-disable -->
- see the [main](https://github.com/cosmosquad-labs/squad/tree/main) branch for the latest 
- see [releases](https://github.com/cosmosquad-labs/squad/releases) for the latest release

## Dependencies

If you haven't already, install Golang by following the [official docs](https://golang.org/doc/install). Make sure that your `GOPATH` and `GOBIN` environment variables are properly set up.

This project uses customized cosmos-sdk, Please check the difference on [here](https://github.com/cosmosquad-labs/cosmos-sdk/compare/v0.44.5...v1.0.2-sdk-0.44.5).

| Requirement           | Notes             |
|-----------------------|-------------------|
| Go version            | Go1.16 or higher  |
| customized cosmos-sdk | v1.0.2-sdk-0.44.5 |

## Installation

```bash
# Use git to clone the source code and install `squad`
git clone https://github.com/cosmosquad-labs/squad.git
cd squad
make install
```

## Getting Started

To get started to the project, visit the [TECHNICAL-SETUP.md](./TECHNICAL-SETUP.md) docs.

## Documentation

The Squad documentation is available in [docs](./docs) folder and technical specification is available in `x/{module}/spec/` folder.

## Contributing

We welcome contributions from everyone. The [main](https://github.com/cosmosquad-labs/squad/tree/main) branch contains the development version of the code. You can branch of from main and create a pull request, or maintain your own fork and submit a cross-repository pull request. If you're not sure where to start check out [CONTRIBUTING.md](./CONTRIBUTING.md) for our guidelines & policies for how we develop squad. Thank you to all those who have contributed to squad!
