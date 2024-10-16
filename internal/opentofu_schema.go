package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

func ExtractOpenTofuProviderSchema(rootDir string) (ProviderSchema, error) {
	schema := ProviderSchema{
		ResourceTypes: make(map[string]map[string]interface{}),
	}

	initCmd := exec.Command("tofu", "init")
	initCmd.Dir = rootDir
	var initStderr bytes.Buffer
	initCmd.Stderr = &initStderr
	if err := initCmd.Run(); err != nil {
		return schema, fmt.Errorf("error running tofu init: %v\nStderr: %s", err, initStderr.String())
	}

	// Now run tofu providers schema -json
	cmd := exec.Command("tofu", "providers", "schema", "-json")
	cmd.Dir = rootDir
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	output, err := cmd.Output()
	if err != nil {
		return schema, fmt.Errorf("error running tofu providers schema -json: %v\nStderr: %s", err, stderr.String())
	}

	var providerData map[string]interface{}
	if err := json.Unmarshal(output, &providerData); err != nil {
		return schema, fmt.Errorf("error unmarshaling provider schema JSON: %v", err)
	}

	if providerSchemas, ok := providerData["provider_schemas"].(map[string]interface{}); ok {
		for _, provider := range providerSchemas {
			if providerMap, ok := provider.(map[string]interface{}); ok {
				if resourceSchemas, ok := providerMap["resource_schemas"].(map[string]interface{}); ok {
					for resType, attributes := range resourceSchemas {
						if attrMap, ok := attributes.(map[string]interface{}); ok {
							schema.ResourceTypes[resType] = attrMap
						}
					}
				}
			}
		}
	}

	return schema, nil
}