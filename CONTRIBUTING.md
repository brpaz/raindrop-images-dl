# Contributing

Thank you for your interest in contributing to this project. We appreciate your efforts in helping us improve and grow the project. Below are the guidelines to ensure that your contributions are effective and fit well with the project.

## How to Contribute

There multiple ways you can contribute to this project, and not everything requires coding.

### Reporting Bugs

If you find a bug, feel free to submit an [issue](https://github.com/brpaz/raindrop-images-dl/issues) on GitHub.
Please include:

- A clear and descriptive title.
- Steps to reproduce the issue.
- Expected and actual results.
- Any relevant logs or error messages.

### Suggesting Features

We welcome ideas for new features or improvements. If you have a suggestion, submit an issue and include:

- A clear description of the proposed feature.
- Why the feature is useful and how it benefits the project.
- Any implementation ideas, if applicable.


## Setup your development environment

This project is built with Go, so it´s very straightforward to run it the local envrionment.

Still, this guide will use [Nix Flakes](https://nixos.wiki/wiki/Flakes).

Flakes helps managing development dependencies and ensure a reproducible environment. By using Nix, you can avoid issues with different versions of tools or dependencies on your machine.

### Prerequisites

1. **Install Nix**
   Follow the official [Nix installation guide](https://nixos.org/download.html) to install Nix on your system.

2. **Enable Flakes**
   Make sure you have Flakes enabled by adding the following lines to your Nix configuration (`/etc/nix/nix.conf` or `~/.config/nix/nix.conf`):
   ```ini
   experimental-features = nix-command flakes
    ```
3. **Direnv** helps managing Envrionment variables and nix envrionments in your project. Follow the instructions at [Direnv](https://direnv.net/) website and make sure direnv is properly configured in your shell.

### Fork the Repository

Fork the repository to your GitHub account by clicking the "Fork" button at the top of the project page. This will create a personal copy of the repository where you can work on your changes.

### Clone Your Fork

After forking the repository, clone your fork locally:

```bash
git clone https://github.com/YOUR_USERNAME/raindrop-images-dl.git
cd raindrop-images-dl
```

### Enter the development shell

You can now start an new dev environment using Nix.
This will automatically build and install this project dependencies.

```shell
nix develop
```

### Run the application

From the nix shell, you can run the application using go run:

```shell
 go run main.go download \
    --collection<my_collection> \
    -k <api_key>
    -o /tmp
```

### Run common tasks

[Taskfile](https://taskfile.dev) is used as Task runner for this project. It´s automatically available if you are using Nix shell.

The following tasks are provided:

Here’s the updated Markdown table with the `aliases`, `summary`, and `commands` columns removed:

| Task Name           | Description                                                             |
| ------------------- | ----------------------------------------------------------------------- |
| `lint`              | Runs all linting tasks                                                  |
| `lint-go`           | Lints Go code using [golangci-lint](https://golangci-lint.run/)         |
| `lint-docker`       | Lints Dockerfile using [Hadolint](https://github.com/hadolint/hadolint) |
| `fmt`               | Formats all code                                                        |
| `gomarkdoc`         | Generates documentation from Go comments                                |
| `test`              | Runs application tests                                                  |
| `test-cover-report` | Opens the test coverage report in the browser                           |
| `snapshot`          | Builds a snapshot release using [GoReleaser](https://goreleaser.com/)   |
| `build-local`       | Builds the application locally                                          |
| `build-docker`      | Builds the application using Docker                                     |

## Build with Docker

A Dockerfile is provided to build the application as a Docker image.

To build a docker image, run:

```shell
task build-docker
```

And run it using the following command:

```shell
docker run --rm raindrop-images-dl:local-dev <args>
```

### Build with Nix

```shell
nix build
./result/bin/raindrop-images-dl
```

## Development Flow

This project uses [GitHub flow](https://docs.github.com/en/get-started/using-github/github-flow) where every change must be a Pull request.

This allows the maintainers to review the code and automated CI tests to run, before merging the code into main branch.

### Commit conventions

This project follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) standard for commit messages. This helps maintain a clean and organized commit history, making it easier to track changes and automate release processes.

The general format is:

```
<type>(<scope>): <description>
```

- type: Describes the type of change (e.g., feat, fix, docs, style, refactor, test, chore).
- scope: Optional, specifies the section or module of the project affected by the commit (e.g., api, ui, auth).
- description: A short summary of the change (imperative tense, no period).

Examples:

-   `feat(auth): add user authentication flow`
-   `fix(ui): resolve button alignment issue`
-   `docs(readme): update installation instructions`

## Commit hooks

This project uses [Lefthook](https://github.com/evilmartians/lefthook) as Git Hooks manager.

Git hooks allows to run tasks like linters and tests before each commit, ensure the quality of the code.

The following commit hooks are provided:

| **Hook**     | **Description**                                                                               | **Run When**                                                  |
| ------------ | --------------------------------------------------------------------------------------------- | ------------------------------------------------------------- |
| `commit-msg` | Lints commit messages to ensure they follow the commit message guidelines using `commitlint`. | Runs during `commit-msg` phase, checking each commit message. |
| `pre-commit` | Lints Go files using `golangci-lint`, fixing issues introduced since the last revision.       | Runs before each commit on Go files (`**/*.go`).              |
| `pre-commit` | Runs Dockerfile linting via a custom task.                                                    | Runs before each commit on Dockerfiles (`Dockerfile`).        |
| `pre-commit` | Formats Go files with `gofumpt` to ensure consistent formatting.                              | Runs before each commit on Go files (`**/*.go`).              |

This version now only includes the hook name, its description, and when it runs.


> [!TIP]
> Lefthook is automatically installed in the project, when using Nix. If you are running on your host machine directly, you can use `lefthook install` command to install the hooks manually. Follow [Lefthook install guide](https://github.com/evilmartians/lefthook?tab=readme-ov-file#install) first to install Lefthook on your machine.

## CI/CD

[GitHub Actions](https://github.com/features/actions) is used as CI/CD pipeline.

The following workflows are available:

- [ci](.github/workflows//ci.yml) - Runs the main CI pipeline.
- [pr-label-checker](.github/workflows/pr-label-checker.yml) - Ensure each Pull request contains the required labels for the release drafter to function. Check [labels.yml](.github/labels.yml) for the possible labels.
- [sync-labels](.github/workflows/sync-labels.yml) - Syncs the labels defined in [labels.yml](.github/labels.yml) with the GitHub Repo. This allows to manage the available labels as code.
- [update-changelog](.github/workflows/update-changelog.yml) - Runs after a release to update project [CHANGELOG](CHANGELOG) file.

### Release process

To streamline the release process, we use [Release Drafter](https://github.com/release-drafter/release-drafter).

This setup automatically drafts and generates release notes based on merged pull requests, making the release process more efficient and consistent.

When a new release is published, GitHub Actions will build the application and:

- publish new binaries and artifacts
- publish a docker image to GitHub Container Registry and Docker Hub.
- update [CHANGELOG](CHANGELOG.md) file.

