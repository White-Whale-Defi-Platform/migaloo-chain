# Migaloo

![](https://user-images.githubusercontent.com/94062656/215557558-6d0c39f1-9405-439a-aeb5-9baccdbd9df8.png)

[![Go Report Card](https://goreportcard.com/badge/White-Whale-Defi-Platform/migaloo-chain)](https://goreportcard.com/report/White-Whale-Defi-Platform/migaloo-chain)

Migaloo Chain is the home of the White Whale.

This chain began as a fork of wasmd, and is an exploration into better CosmWasm enabled chain templates that more
rigorously follow upstream standards. It began as the wasmd repository, and we're hoping that it will end up as a
feather/ignite/whatever template.

## Helpful Information

Our testnet has Alliance on it, but our mainnet won't have alliance until there's a more-stable release from <https://github.com/terra-money/alliance>.

Because of this, branching is like:

* `main` is the tip of the repository and may not reflect working testnet or mainnet state
* `release/v1.0.x` is the tip of the mainnet branch, and reflects working mainnet state
* `release/v2.0.x` is the tip of the testnet branch, and reflects working testnet state (and is alliance enabled)

## Resources

1. [Website](https://migaloo.zone)
2. [LitePaper]() - Coming Soon
3. [Docs]() - Coming Soon
4. [Roadmap](./docs/ROADMAP.md)
5. [Discord](https://discord.com/channels/908044702794801233/1069611972053712947)
6. [Twitter](https://twitter.com/WhiteWhaleDefi)
7. [Telegram](https://t.me/whitewhaleofficial)

## System Requirements

* Operating System: Linux or macOS
* Disk Space: At least 100GB of free space is recommended.
* CPU: Multi-core processor, 4+ cores recommended
* RAM: 8GB+ recommended
* Network: Good internet connectivity

## Quick start

Requires [Go 1.20](https://go.dev/doc/install) or higher.

```bash
make install
migalood version
```

## Contributing

[Contributing Guide](./docs/CONTRIBUTING.md)

[Code of Conduct](./docs/CODE_OF_CONDUCT.md)

[Security Policies and Procedures](./docs/SECURITY.md)

[License](./LICENSE)

## Learn More

* [White Whale Protocol](https://whitewhale.money/)
* [Cosmos SDK documentation](https://docs.cosmos.network/)
* [Cosmos SDK Tutorials](https://tutorials.cosmos.network/)

## Disclaimer

**Migaloo software is offered "as is", with the understanding that the user assumes all risks and no guarantees or warranties are provided.**
