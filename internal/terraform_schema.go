// Copyright (c) RoseSecurity
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type ProviderSchema struct {
	ResourceTypes map[string]map[string]any
}

func ExtractTerraformProviderSchema(rootDir string) (ProviderSchema, error) {
	schema := ProviderSchema{
		ResourceTypes: make(map[string]map[string]any),
	}

	// Run terraform init first
	initCmd := exec.Command("terraform", "init")
	initCmd.Dir = rootDir
	var initStderr bytes.Buffer
	initCmd.Stderr = &initStderr
	if err := initCmd.Run(); err != nil {
		return schema, fmt.Errorf("error running terraform init: %v\nStderr: %s", err, initStderr.String())
	}

	// Now run terraform providers schema -json
	cmd := exec.Command("terraform", "providers", "schema", "-json")
	cmd.Dir = rootDir
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	output, err := cmd.Output()
	if err != nil {
		return schema, fmt.Errorf("error running terraform providers schema -json: %v\nStderr: %s", err, stderr.String())
	}

	var providerData map[string]any
	if err := json.Unmarshal(output, &providerData); err != nil {
		return schema, fmt.Errorf("error unmarshaling provider schema JSON: %v", err)
	}

	if providerSchemas, ok := providerData["provider_schemas"].(map[string]any); ok {
		for _, provider := range providerSchemas {
			if providerMap, ok := provider.(map[string]any); ok {
				if resourceSchemas, ok := providerMap["resource_schemas"].(map[string]any); ok {
					for resType, attributes := range resourceSchemas {
						if attrMap, ok := attributes.(map[string]any); ok {
							schema.ResourceTypes[resType] = attrMap
						}
					}
				}
			}
		}
	}

	return schema, nil
}
