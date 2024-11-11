## kuzco fix

Diagnose configuration errors

### Synopsis

This command analyzes and diagnoses Terraform configuration errors

```
kuzco fix [flags]
```

### Examples

```
kuzco fix -f path/to/config.tf -t terraform
```

### Options

```
  -a, --address string   IP Address and port to use for the LLM model (ex: http://localhost:11434) (default "http://localhost:11434")
  -f, --file string      Path to the Terraform and OpenTofu file (required)
  -h, --help             help for fix
  -m, --model string     LLM model to use for generating recommendations (default "llama3.2")
  -t, --tool terraform   Specifies the configuration tooling for configurations. Valid values include: terraform and `opentofu` (default "terraform")
```

### SEE ALSO

* [kuzco](kuzco.md)	 - Intelligently analyze your Terraform and OpenTofu configurations

