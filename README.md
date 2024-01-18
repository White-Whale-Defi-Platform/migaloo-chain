# Migaloo

![](https://user-images.githubusercontent.com/94062656/215557558-6d0c39f1-9405-439a-aeb5-9baccdbd9df8.png)

[![Go Report Card](https://goreportcard.com/badge/White-Whale-Defi-Platform/migaloo-chain)](https://goreportcard.com/report/White-Whale-Defi-Platform/migaloo-chain)

[![OpenSSF Best Practices](https://bestpractices.coreinfrastructure.org/projects/7139/badge)](https://bestpractices.coreinfrastructure.org/projects/7139)

Migaloo Chain is the home of the White Whale.

This chain began as a fork of wasmd, and is an exploration into better CosmWasm enabled chain templates that more
rigorously follow upstream standards. It began as the wasmd repository, and we're hoping that it will end up as a
feather/ignite/whatever template.

## Helpful Information

Our testnet has Alliance on it, but our mainnet won't have alliance until there's a more-stable release from <https://github.com/terra-money/alliance>.

Because of this, branching is like:


* `release/v1.0.x` is the tip of the mainnet branch, and reflects working mainnet state until the launch of v2
* `release/v2.0.x` is the tip of the branch for v2, which enables alliance.
* `release/v3.0.x` is the tip of the branch for v3, which will add osmosis ibc hooks, async-icq.
* `release/v4.1.x` is the current development branch, which upgrades to ibc v7, sdk 47, cometbft 37, and wasmd v0.45.0


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

Requires [Go 1.21](https://go.dev/doc/install) or higher.

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
