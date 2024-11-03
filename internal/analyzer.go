package internal

import (
	"fmt"
	"path/filepath"
	"strings"
)

func Run(filePath, tool, model, prompt, addr string) error {
	if !(strings.HasSuffix(filePath, ".tf") || strings.HasSuffix(filePath, ".tofu")) {
		return fmt.Errorf("the provided file must have a .tf or .tofu extension")
	}

	resources, err := ParseConfigurationFile(filePath)
	if err != nil {
		return fmt.Errorf("error parsing configuration file: %v", err)
	}

	dir := filepath.Dir(filePath)
	var providerSchema ProviderSchema

	switch tool {
	case "terraform":
		providerSchema, err = ExtractTerraformProviderSchema(dir)
		if err != nil {
			return fmt.Errorf("error extracting Terraform provider schema: %v", err)
		}
	case "opentofu":
		providerSchema, err = ExtractOpenTofuProviderSchema(dir)
		if err != nil {
			return fmt.Errorf("error extracting OpenTofu provider schema: %v", err)
		}
	default:
		return fmt.Errorf("unsupported tool: %s. Supported tools are 'terraform' and 'opentofu'", tool)
	}

	if err := printDiff(resources, providerSchema, model, tool, prompt, addr); err != nil {
		return fmt.Errorf("error printing differences: %v", err)
	}

	return nil
}

func printDiff(resources []Resource, schema ProviderSchema, model, tool, prompt, addr string) error {
	for _, resource := range resources {
		if possibleAttrs, ok := schema.ResourceTypes[resource.Type]; ok {
			usedAttrs := resource.Attributes
			unusedAttrs := findUnusedAttributes(usedAttrs, possibleAttrs)

			if len(unusedAttrs) > 0 {
				// Get recommendations based on unused attributes
				recommendations, err := GetRecommendations(resource.Type, unusedAttrs, model, tool, prompt, addr)
				if err != nil {
					return fmt.Errorf("error getting recommendations for resource %s: %v", resource.Name, err)
				}
				PrettyPrint(recommendations)
			} else {
				fmt.Printf("Resource: %s (Type: %s) - All attributes are used.\n\n", resource.Name, resource.Type)
			}
		} else {
			fmt.Printf("Warning: Resource type %s not found in schema.\n", resource.Type)
		}
	}
	return nil
}

func findUnusedAttributes(usedAttrs map[string]string, possibleAttrs map[string]interface{}) []string {
	var unusedAttrs []string
	for attr := range possibleAttrs {
		if _, used := usedAttrs[attr]; !used {
			unusedAttrs = append(unusedAttrs, attr)
		}
	}
	return unusedAttrs
}
