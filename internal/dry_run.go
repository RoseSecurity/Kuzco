// Copyright (c) RoseSecurity
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

// DryRun checks the provided file for unused attributes based on a provider schema.
func DryRun(filePath, tool string) ([]string, error) {
	if !(strings.HasSuffix(filePath, ".tf") || strings.HasSuffix(filePath, ".tofu")) {
		return nil, fmt.Errorf("the provided file must have a .tf or .tofu extension")
	}

	resources, err := ParseConfigurationFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error parsing configuration file: %v", err)
	}

	// Placeholder schema, which would typically be loaded based on the provider type
	var providerSchema ProviderSchema
	dir := filepath.Dir(filePath) // Ensure `dir` is correctly derived

	switch tool {
	case "terraform":
		providerSchema, err = ExtractTerraformProviderSchema(dir)
		if err != nil {
			return nil, fmt.Errorf("error extracting Terraform provider schema: %v", err)
		}
	case "opentofu":
		providerSchema, err = ExtractOpenTofuProviderSchema(dir)
		if err != nil {
			return nil, fmt.Errorf("error extracting OpenTofu provider schema: %v", err)
		}
	default:
		return nil, fmt.Errorf("unsupported tool: %s. Supported tools are 'terraform' and 'opentofu'", tool)
	}

	// Identify and return unused attributes
	unusedAttrs, err := testPossibleAttributes(resources, providerSchema, tool)
	if err != nil {
		return nil, fmt.Errorf("error identifying unused attributes: %v", err)
	}

	return unusedAttrs, nil
}

// testPossibleAttributes checks each resource for unused attributes and returns them.
func testPossibleAttributes(resources []Resource, schema ProviderSchema, tool string) ([]string, error) {
	var unusedAttrs []string
	for _, resource := range resources {
		if possibleAttrs, ok := schema.ResourceTypes[resource.Type]; ok {
			usedAttrs := resource.Attributes
			unusedAttrsForResource := testFindUnusedAttributes(usedAttrs, possibleAttrs)
			// Collect unused attributes
			unusedAttrs = append(unusedAttrs, unusedAttrsForResource...)
		} else {
			fmt.Printf("No schema found for resource type %s. Skipping unused attribute check.\n", resource.Type)
		}
	}
	return unusedAttrs, nil
}

// testFindUnusedAttributes identifies unused attributes by comparing used and possible attributes.
func testFindUnusedAttributes(usedAttrs map[string]string, possibleAttrs map[string]any) []string {
	// Mirror logic in analyzer's findUnusedAttributes: only consider names
	// under block.attributes and block.block_types.
	validNames := make(map[string]struct{})

	if blockAny, ok := possibleAttrs["block"]; ok {
		if block, ok := blockAny.(map[string]any); ok {
			if attrsAny, ok := block["attributes"]; ok {
				if attrsMap, ok := attrsAny.(map[string]any); ok {
					for name := range attrsMap {
						validNames[name] = struct{}{}
					}
				}
			}
			if blockTypesAny, ok := block["block_types"]; ok {
				if btMap, ok := blockTypesAny.(map[string]any); ok {
					for name := range btMap {
						validNames[name] = struct{}{}
					}
				}
			}
		}
	}

	var unused []string
	for name := range validNames {
		if _, used := usedAttrs[name]; !used {
			unused = append(unused, name)
		}
	}
	sort.Strings(unused)

	return unused
}
