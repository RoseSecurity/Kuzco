//go:build generate

package tools

import (
	_ "github.com/hashicorp/copywrite"
)

// Generate copyright headers
//go:generate go run github.com/hashicorp/copywrite headers -d .. --config ../.copywrite.hcl
