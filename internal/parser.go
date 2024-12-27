// Copyright (c) RoseSecurity
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
)

type Resource struct {
	Type       string
	Name       string
	Attributes map[string]string
}

func ParseConfigurationFile(file string) ([]Resource, error) {
	var resources []Resource
	parser := hclparse.NewParser()

	hclFile, diag := parser.ParseHCLFile(file)
	if diag.HasErrors() {
		return nil, fmt.Errorf("error parsing file %s: %s", file, diag.Error())
	}

	resourceSchema := &hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{
				Type:       "resource",
				LabelNames: []string{"type", "name"},
			},
		},
	}

	content, _, diags := hclFile.Body.PartialContent(resourceSchema)
	if diags.HasErrors() {
		return nil, fmt.Errorf("error decoding HCL in file %s: %s", file, diags.Error())
	}

	for _, block := range content.Blocks {
		if len(block.Labels) < 2 {
			continue
		}
		resourceType := block.Labels[0]
		resourceName := block.Labels[1]
		attributes := make(map[string]string)

		bodyContent, _, _ := block.Body.PartialContent(&hcl.BodySchema{
			Attributes: []hcl.AttributeSchema{},
		})

		for attrName, attr := range bodyContent.Attributes {
			value, diag := attr.Expr.Value(nil)
			if diag.HasErrors() {
				fmt.Printf("Error evaluating attribute %s in file %s: %s\n", attrName, file, diag.Error())
				continue
			}
			attributes[attrName] = value.AsString()
		}

		resources = append(resources, Resource{
			Type:       resourceType,
			Name:       resourceName,
			Attributes: attributes,
		})
	}

	return resources, nil
}
