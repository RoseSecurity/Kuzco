# Kuzco

<p align="center">
<img width=100% height=100% src="./docs/img/kuzco-logo.png">
</p>

<p align="center">
  <em>Enhance your Terraform configurations with intelligent analysis powered by local LLMs</em>
</p>

## Introduction

Kuzco is an intelligent tool that analyzes your Terraform configurations and provides personalized recommendations to improve efficiency, security, and performance. By leveraging powerful machine learning models, Kuzco evaluates your infrastructure as code and suggests optimizations tailored to your specific use case.

## Installation

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
