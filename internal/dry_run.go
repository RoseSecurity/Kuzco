// Copyright (c) RoseSecurity
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"fmt"
	"path/filepath"
	"strings"

	log "github.com/charmbracelet/log"
)

// DryRun parses a single .tf or .tofu file, compares declared attributes and
// nested blocks against the provider schema, and returns a slice of attribute
// or block names that are defined in the schema but absent in the file.
func DryRun(filePath, tool string) ([]string, error) {
	if !isSupportedFile(filePath) {
		return nil, fmt.Errorf("unsupported file extension: %s", filePath)
	}

	// ── Parse configuration ────────────────────────────────────────────────────
	resources, err := ParseConfigurationFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("parse error: %w", err)
	}
	if len(resources) == 0 {
		log.Warnf("no resources found in %s", filePath)
		return nil, nil
	}

	// ── Load provider schema ───────────────────────────────────────────────────
	dir := filepath.Dir(filePath)

	var schema ProviderSchema
	switch strings.ToLower(tool) {
	case "terraform":
		schema, err = ExtractTerraformProviderSchema(dir)
	case "opentofu":
		schema, err = ExtractOpenTofuProviderSchema(dir)
	default:
		return nil, fmt.Errorf("unsupported tool: %s (expected terraform|opentofu)", tool)
	}
	if err != nil {
		return nil, fmt.Errorf("provider schema extraction error: %w", err)
	}

	// ── Diff used vs. possible attributes/blocks ───────────────────────────────
	return findUnused(schema, resources), nil
}

// isSupportedFile returns true for .tf and .tofu files.
func isSupportedFile(p string) bool {
	return strings.HasSuffix(p, ".tf") || strings.HasSuffix(p, ".tofu")
}

// findUnused returns a flattened list of schema attributes/blocks that are not
// present in any of the parsed resources.
func findUnused(schema ProviderSchema, resources []Resource) []string {
	var unused []string

	for _, r := range resources {
		possible, ok := schema.ResourceTypes[r.Type]
		if !ok {
			log.Warnf("no provider schema for resource type %s", r.Type)
			continue
		}

		if len(r.Attributes)+len(r.UsedBlocks) == len(possible) {
			continue // all defined items are present
		}

		unused = append(unused, diff(r.Attributes, r.UsedBlocks, possible)...)
	}

	return unused
}

// diff returns items present in possible but absent in usedAttrs / usedBlocks.
func diff(
	usedAttrs map[string]interface{},
	usedBlocks map[string]bool,
	possible map[string]interface{},
) []string {
	var missing []string

	for name, schemaVal := range possible {
		// Nested blocks are represented by map[string]interface{} in the schema.
		if _, isBlock := schemaVal.(map[string]interface{}); isBlock {
			if !usedBlocks[name] {
				missing = append(missing, name)
			}
			continue
		}

		if _, ok := usedAttrs[name]; !ok {
			missing = append(missing, name)
		}
	}

	return missing
}
