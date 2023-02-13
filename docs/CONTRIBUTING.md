
# Contributing to Migaloo

Welcome! We are thrilled that you want to contribute to Migaloo! Consider that there are many ways in which you can contribute,
it's not only about writing code. In this document we go through different ways you can get involved with the project.

## Crafting Code

If you are a software developer and want to contribute writing code, the first step is to get familiar with
the Migaloo architecture, which you can learn about in our [docs]().

Before you can write any code, please take a look at the list of prerequisites below.

### Prerequisites

To download the necessary tools, clone the repository and so on, you need network access.

The following are the tools you'll need:

- [Git](https://git-scm.com/)
- [Go](https://go.dev/doc/install)
- [Docker](https://www.docker.com/)
- An IDE. We recommend either [GoLand](https://www.jetbrains.com/go/download/), [Intellij Idea](https://www.jetbrains.com/idea/download/) or [VS Code](https://code.visualstudio.com/download),
  but feel free to use your preferred code editor.
- [Go Plugin](https://plugins.jetbrains.com/plugin/9568-go) for Intellij Idea or the [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.go) for VS Code.

---
Then fork the code and read it through. We encourage you to make your own contributions, though you might look at the
[issue tracker](https://github.com/White-Whale-Defi-Platform/migaloo-chain/issues) if you want to solve something that has
been pointed out already.

Make a pull request to our repository once your work is complete. We will review it and discuss potential changes before
we merge it to the main code base.

### Forking the repository

The following are the steps to fork the repository to your GitHub account and clone it to your local machine.

1. Fork the repository.
    - Go to Migaloo's Chain repository, i.e. [Migaloo-Chain](https://github.com/White-Whale-Defi-Platform/migaloo-chain).
    - Click on the ["fork" button](https://github.com/White-Whale-Defi-Platform/migaloo-chain/fork) located in the top
      right corner of the page.
    - Click on your GitHub account where you want to fork the repo.
2. Clone your fork to your local machine, preferably using the SSH URL. If you have issues cloning this repo, look at the
   [GitHub docs](https://docs.github.com/en/repositories/creating-and-managing-repositories/cloning-a-repository).
    - `git clone git@github.com:$USER/migaloo-chain.git` or `git clone https://github.com/$USER/migaloo-chain.git`
3. Set up your git user locally if you haven't already.
    - `git config --global user.name "your name or alias"`
    - `git config --global user.email "your email address"`
4. Make your contributions locally. The following are recommendations so that it is easier for anyone to understand what
   you are trying to achieve:
    - Please use [conventional commits](https://conventionalcommits.org) syntax.
    - Please make sure to use clear commit messages.
    - Please favor small commits instead of large ones.
5. Make sure your code compiles.
    - `make build`
    - `go build ./...`
6. Test your code. We strive for high quality code, so any changes you introduce need to be tested. We know testing contracts
   can be difficult! If you are not sure how to create tests, please refer to existing ones or just ask us on our
   [discord](https://discord.com/channels/908044702794801233/1069611972053712947). Please note that **Untested code will be rejected**
    - `go test ./...`
7. Push your changes to your repository.
    - `git push --set-upstream $YOUR_ORIGIN $YOUR_BRANCH_NAME`
8. Create a pull request. Go to your [repository](https://github.com/$USER/migaloo-chain.git) and create a pull request
    against Migaloo's repository **main branch** as base.
    - Please fill in the template presented to you when creating the pull request, as it helps everybody understand what the PR is trying to achieve.
    - Follow the instructions on the PR template. Consider that a pull request that doesn't follow the template or is not filled in properly *will be considered incomplete*.
9. Follow up the discussions on the PR as there might be requests from other members.
10. Wait for your PR to be approved and merged.

## Helping out in the issue tracker

We use [Github issues](https://github.com/White-Whale-Defi-Platform/migaloo-chain/issues) to manage issues in our code.
You can help out by resolving or commenting on existing issues or creating new issues for what you find. Whether you want
to report an issue or have a feature request, please fill the issue template and provide as much information as possible.

### Look For an Existing Issue

Before you create a new issue, please search through the open issues and make sure the issue or feature request has not
been made by someone else already.

If the issue or feature request already exists, please add a 👍 reaction to show your support and leave your comments on it, that way we can prioritize accordingly.

## Quality Assurance

Code quality and security are two things we take seriously at White Whale. We strive for having high test coverage, and
we make sure our code is safe by auditing via third party security firms. Nevertheless, we believe there's always room for improvement.

If you find a critical vulnerability, please do not report it publicly on the Github issues tracker. Instead, reach out to us
in private where we will discuss it in details.

See how to [report security bugs](https://github.com/White-Whale-Defi-Platform/migaloo-chain/blob/main/docs/SECURITY.md).

## Engaging with the community

If you are interested in developer relations, a great way to contribute is answering people's questions on our
[Discord](https://discord.com/channels/908044702794801233/1069611972053712947), [Twitter](https://twitter.com/WhiteWhaleDefi)
or [Telegram](https://t.me/whitewhaleofficial) channels, creating documentation or even medium articles!.

## Providing Suggestions

Migaloo is a community project, we are curious to hear your ideas for the future of it! One way to provide feedback
is by doing to our [Discord channel](https://discord.com/channels/908044702794801233/1069611972053712947). You can also
submit a suggestion or feature request through [Github's issue tracker](https://github.com/White-Whale-Defi-Platform/migaloo-chain/issues).
When doing so, make sure to describe your idea as good as possible so that we understand what you envision.

## Bringing your expertise

Are you a graphic designer and want to help out with some NFTs? Are you a copywriter seeing potential improvements in our communications?
Whatever it is, we would love to hear from you and see how we can make White Whale better for everybody.

## Docs

When in doubt, please take a look at our [litepaper]() or our [documentation]([]()).

# Thank you

All contributions to Migaloo are of great value and make this chain possible. Thanks for taking the time to make
Migaloo better! We really appreciate it.
