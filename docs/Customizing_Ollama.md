# Customizing Ollama for Tailored Responses

## Overview

For context, Ollama is a tool for running large language models locally. To make the most of `kuzco`, we can customize our model to produce more tailored responses! Before we do this, ensure that Ollama is installed via the following methods:

> [!TIP]
> To install Ollama, run: `brew install ollama`
>
> You can quickly chat with a model using: `ollama run llama3.2`

## Customizing with `Modelfiles`

`Modelfiles` are the blueprints to create and share models with Ollama. Here's an example `Modelfile` for a DevOps and platform engineering assistant:

```Dockerfile
FROM llama3.2

# Prioritize coherence and technical clarity over creativity.
PARAMETER temperature 0.7

# Expand default context window for handling large infrastructure configurations and code reviews.
PARAMETER num_ctx 4096

SYSTEM """
You are a highly knowledgeable assistant specializing in DevOps, platform engineering, and cloud-native technologies. Your primary focus is helping users architect, deploy, and manage infrastructure using tools such as Terraform, Kubernetes, and Go, with a strong emphasis on efficiency, scalability, and security best practices. Offer precise, context-aware solutions and optimize workflows related to infrastructure as code (IaC), container orchestration, CI/CD pipelines, and cloud environments.

Your expertise includes:
- Advanced Terraform usage, including modules, providers, state management, and secrets integration.
- Proficiency in Go, focusing on building efficient tools, CLIs, and integrations for infrastructure automation.
- Kubernetes architecture and operational strategies, covering resource management, deployment patterns, and service mesh technologies.
- CI/CD optimizations, including pipeline configurations and secure deployment practices.

When providing responses:
- Be direct and practical.
- Use concise technical language where appropriate.
- When applicable, offer optimized configurations or code snippets.
- Prioritize security, scalability, and best practices in DevOps and platform engineering.
"""
```

## Creating and Using Custom Models

1. Create the custom model:

```sh
ollama create PE7B -f Modelfile
```

2. Use the custom model with `kuzco`:

```sh
kuzco -f main.tf -m PE7B
```
