# Kuzco

<p align="center">
<img width=90% height=80% src="./docs/img/kuzco-logo.png">
</p>

<p align="center">
  <em>Enhance your Terraform and OpenTofu configurations with intelligent analysis powered by local LLMs</em>
</p>

## Introduction

Here's the **problem**: You spin up a Terraform or OpenTofu resource, pull a basic configuration from the registry, and start wondering what other parameters should be enabled to make it more secure and efficient. Sure, you could use tools like TLint or TFSec, but `kuzco` saves you time by avoiding the need to dig through the Terraform registry and decipher unclear options. It leverages local LLMs to recommend what **should** be enabled and configured. Simply put, `kuzco` reviews your Terraform and OpenTofu resources, compares them to the provider schema to detect unused parameters, and uses AI to suggest improvements for a more secure, reliable, and optimized setup.

## Demo

<p align="center">
<img width=100% height=100% src="./docs/img/kuzco-demo.gif">
</p>

## Installation

> [!NOTE]
> To use `kuzco`, Ollama must be installed. You can do this by running `brew bundle install` or `brew install ollama`. For more information on customizing Ollama models for tailored Kuzco responses, check out [Customizing Ollama](./docs/Customizing_Ollama.md)

### Go

If you have a functional Go environment, you can install with:

```sh
go install github.com/RoseSecurity/kuzco@latest
```

### Apt

To install packages, you can quickly setup the repository automatically:

```sh
curl -1sLf \
  'https://dl.cloudsmith.io/public/rosesecurity/kuzco/setup.deb.sh' \
  | sudo -E bash
```

Once the repository is configured, you can install with:

```sh
apt install kuzco=<VERSION>
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


Intelligently analyze your Terraform and OpenTofu configurations to receive personalized recommendations for boosting efficiency, security, and performance.

Usage:
  kuzco [flags]
  kuzco [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     Print the CLI version

Flags:
  -a, --address string   IP Address and port to use for the LLM model (ex: http://localhost:11434) (default "http://localhost:11434")
  -f, --file string      Path to the Terraform and OpenTofu file (required)
  -h, --help             help for kuzco
  -m, --model string     LLM model to use for generating recommendations (default "llama3.2")
  -t, --tool terraform   Specifies the configuration tooling for configurations. Valid values include: terraform and `opentofu` (default "terraform")

Use "kuzco [command] --help" for more information about a command.
```
