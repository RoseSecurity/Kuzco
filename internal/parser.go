// Copyright (c) RoseSecurity
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/zclconf/go-cty/cty"
)

// Resource represents a Terraform/OpenTofu resource block.
type Resource struct {
	Type       string          // e.g. "aws_instance"
	Name       string          // e.g. "web"
	Attributes map[string]any  // flattened attribute map
	UsedBlocks map[string]bool // set of nested-block types encountered
}

// ParseConfigurationFile extracts every `resource` block from a *.tf or *.tftpl
// file, returning a slice of Resource objects with type, name, attributes, and
// the set of nested block types used under each resource.
func ParseConfigurationFile(path string) ([]Resource, error) {
	parser := hclparse.NewParser()
	hclFile, diags := parser.ParseHCLFile(path)
	if diags.HasErrors() {
		return nil, fmt.Errorf("HCL parse error: %s", diags.Error())
	}

	// Top-level schema – grab all resource blocks regardless of provider.
	resourceSchema := &hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{Type: "resource", LabelNames: []string{"type", "name"}},
		},
	}

	rootContent, _, diags := hclFile.Body.PartialContent(resourceSchema)
	if diags.HasErrors() {
		return nil, fmt.Errorf("HCL decode error: %s", diags.Error())
	}

	evalCtx := &hcl.EvalContext{Variables: map[string]cty.Value{}}
	var resources []Resource

	for _, b := range rootContent.Blocks {
		if len(b.Labels) < 2 {
			continue // malformed block – skip
		}

		rType, rName := b.Labels[0], b.Labels[1]

		// Allow everything inside – attributes and nested blocks.
		body := &hcl.BodySchema{}
		content, _, diags := b.Body.PartialContent(body)
		if diags.HasErrors() {
			return nil, fmt.Errorf("error decoding %s.%s: %s", rType, rName, diags.Error())
		}

		attrs := make(map[string]any)
		for key, attr := range content.Attributes {
			val, diag := attr.Expr.Value(evalCtx)
			if diag.HasErrors() || !val.IsKnown() {
				continue // skip unevaluable / unknown values
			}
			attrs[key] = convertCty(val)
		}

		used := make(map[string]bool)
		collectNestedBlockTypes(content.Blocks, used)

		resources = append(resources, Resource{
			Type:       rType,
			Name:       rName,
			Attributes: attrs,
			UsedBlocks: used,
		})
	}

	return resources, nil
}

// collectNestedBlockTypes recursively records every nested block type.
func collectNestedBlockTypes(blocks hcl.Blocks, dst map[string]bool) {
	for _, b := range blocks {
		dst[b.Type] = true
		sub, _, _ := b.Body.PartialContent(&hcl.BodySchema{}) // walk deeper
		collectNestedBlockTypes(sub.Blocks, dst)
	}
}

// convertCty flattens a cty.Value into native Go types.
func convertCty(v cty.Value) any {
	if !v.IsKnown() || v.IsNull() {
		return nil
	}

	switch {
	case v.Type().Equals(cty.String):
		return v.AsString()
	case v.Type().Equals(cty.Number):
		n, _ := v.AsBigFloat().Float64()
		return n
	case v.Type().Equals(cty.Bool):
		return v.True()
	case v.Type().IsTupleType() || v.Type().IsListType() || v.Type().IsSetType():
		var out []any
		it := v.ElementIterator()
		for it.Next() {
			_, elem := it.Element()
			out = append(out, convertCty(elem))
		}
		return out
	case v.Type().IsMapType() || v.Type().IsObjectType():
		out := make(map[string]any)
		it := v.ElementIterator()
		for it.Next() {
			key, val := it.Element()
			out[key.AsString()] = convertCty(val)
		}
		return out
	default:
		return v.GoString()
	}
}
