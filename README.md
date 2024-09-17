# Kuzco

<p align="center">
<img width=90% height=80% src="./docs/img/kuzco-logo.png">
</p>

<p align="center">
  <em>Enhance your Terraform configurations with intelligent analysis powered by local LLMs</em>
</p>

## Introduction

Here's the **problem**: You spin up a Terraform resource, pull a basic configuration from the registry, and start wondering what other parameters should be enabled to make it more secure and efficient. Sure, you could use tools like TLint or TFSec, but `kuzco` saves you time by avoiding the need to dig through the Terraform registry and decipher unclear options. It leverages local LLMs to recommend what **should** be enabled and configured. Simply put, `kuzco` reviews your Terraform resources, compares them to the provider schema to detect unused parameters, and uses AI to suggest improvements for a more secure, reliable, and optimized setup.

## Demo

<p align="center">
<img width=100% height=100% src="./docs/img/kuzco-demo.gif">
</p>

## Installation

> [!NOTE]
> To use `kuzco`, Ollama must be installed. You can do this by running `brew bundle install` or `brew install ollama`

### Go

If you have a functional Go environment, you can install with:

```sh
go install github.com/RoseSecurity/kuzco@latest
```

### Source

```sh
git clone git@github.com:RoseSecurity/Kuzco.git
cd Kuzco
make build
```

## Usage

The following configuration options are available:

```sh
❯ kuzco
  _  __
 | |/ /  _   _   ____   ___    ___
 | ' /  | | | | |_  /  / __|  / _ \
 | . \  | |_| |  / /  | (__  | (_) |
 |_|\_\  \__,_| /___|  \___|  \___/

Intelligently analyze your Terraform configurations to receive personalized recommendations for boosting efficiency, security, and performance.

Usage:
  kuzco [flags]
  kuzco [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -f, --file string    Path to the Terraform file (required)
```

> [!CAUTION]
> This tool is in active and early development. Opening issues, triaging bugs, feature requests, and code contributions are appreciated.
