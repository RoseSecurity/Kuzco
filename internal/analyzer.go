// Copyright (c) RoseSecurity
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"fmt"
	"path/filepath"
	"strings"

	log "github.com/charmbracelet/log"
)

func Run(filePath, tool, model, prompt, addr string) error {
	if !(strings.HasSuffix(filePath, ".tf") || strings.HasSuffix(filePath, ".tofu")) {
		return fmt.Errorf("invalid file extension: %s", filePath)
	}

	resources, err := ParseConfigurationFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to parse HCL configuration: %w", err)
	}

	if len(resources) == 0 {
		log.Warnf("No resources found in file: %s", filePath)
		return nil
	}

	dir := filepath.Dir(filePath)
	var providerSchema ProviderSchema

	switch tool {
	case "terraform":
		providerSchema, err = ExtractTerraformProviderSchema(dir)
	case "opentofu":
		providerSchema, err = ExtractOpenTofuProviderSchema(dir)
	default:
		return fmt.Errorf("unsupported tool: %s", tool)
	}

	if err != nil {
		return fmt.Errorf("error extracting provider schema: %w", err)
	}

	if err := printDiff(resources, providerSchema, model, tool, prompt, addr); err != nil {
		return fmt.Errorf("error printing differences: %w", err)
	}

	return nil
}

func printDiff(resources []Resource, schema ProviderSchema, model, tool, prompt, addr string) error {
	for _, resource := range resources {
		possibleAttrs, ok := schema.ResourceTypes[resource.Type]
		if !ok {
			log.Warnf("Warning: Resource type %s not found in schema.", resource.Type)
			continue
		}

		usedAttrs := resource.Attributes
		unusedAttrs := findUnusedAttributes(usedAttrs, possibleAttrs)

		if len(unusedAttrs) > 0 {
			recommendations, err := GetRecommendations(resource.Type, unusedAttrs, model, tool, prompt, addr)
			if err != nil {
				return fmt.Errorf("error fetching recommendations for resource %s: %w", resource.Name, err)
			}
			PrettyPrint(recommendations)
		} else {
			log.Infof("Resource: %s (Type: %s) - All attributes are used.", resource.Name, resource.Type)
		}
	}
	return nil
}

func findUnusedAttributes(usedAttrs map[string]any, possibleAttrs map[string]any) []string {
	var unusedAttrs []string
	for attr := range possibleAttrs {
		if _, used := usedAttrs[attr]; !used {
			unusedAttrs = append(unusedAttrs, attr)
		}
	}
	return unusedAttrs
}
