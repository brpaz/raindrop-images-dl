# Contributing

Thank you for your interest in contributing to this project. We appreciate your efforts in helping us improve and grow the project. Below are the guidelines to ensure that your contributions are effective and fit well with the project.

## How to Contribute

### Reporting Bugs

If you find a bug, feel free to submit an [issue][https://github.com/brpaz/raindrop-images-dl/issues] on GitHub.
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

[Nix Flakes](https://nixos.wiki/wiki/Flakes) are used to manage development dependencies and ensure a reproducible environment. By using Nix, you can avoid issues with different versions of tools or dependencies on your machine.

### Prerequisites

1. **Install Nix**
   Follow the official [Nix installation guide](https://nixos.org/download.html) to install Nix on your system.

2. **Enable Flakes**
   Make sure you have Flakes enabled by adding the following lines to your Nix configuration (`/etc/nix/nix.conf` or `~/.config/nix/nix.conf`):
   ```ini
   experimental-features = nix-command flakes


### Fork the Repository

Fork the repository to your GitHub account by clicking the "Fork" button at the top of the project page. This will create a personal copy of the repository where you can work on your changes.

### Clone Your Fork

After forking the repository, clone your fork locally:

```bash
git clone https://github.com/YOUR_USERNAME/raindrop-images-dl.git
cd raindrop-images-dl
```

### Create a New Branch

Before starting any work, create a new branch:

```shell
git checkout -b feature/your-feature-name
```

Use a meaningful branch name that reflects the nature of your changes (e.g., fix/issue-42, feature/add-new-endpoint).

### Enter the development shell

```shell
nix develop
```
