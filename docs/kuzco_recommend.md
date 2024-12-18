## kuzco recommend

Intelligently analyze your Terraform and OpenTofu configurations

### Synopsis

Intelligently analyze your Terraform and OpenTofu configurations to receive personalized recommendations for boosting efficiency, security, and performance.

```
kuzco recommend [flags]
```

### Options

```
  -a, --address string   IP Address and port to use for the LLM model (ex: http://localhost:11434) (default "http://localhost:11434")
  -f, --file string      Path to the Terraform and OpenTofu file (required)
  -h, --help             help for recommend
  -m, --model string     LLM model to use for generating recommendations (default "llama3.2")
  -p, --prompt string    User prompt for guiding the response format of the LLM model
  -t, --tool terraform   Specifies the configuration tooling for configurations. Valid values include: terraform and `opentofu` (default "terraform")
```

### SEE ALSO

* [kuzco](kuzco.md)	 - Intelligently analyze your Terraform and OpenTofu configurations

