// Copyright (c) RoseSecurity
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"errors"
)

var (
	errInvalidFileExtension    = errors.New("the provided file must have a .tf or .tofu extension")
	errParsingConfig           = errors.New("error parsing configuration file")
	errProviderExtraction      = errors.New("error extracting provider schema")
	errUnsupportedTool         = errors.New("unsupported tool, supported tools are 'terraform' and 'opentofu'")
	errPrintingDifferences     = errors.New("error printing differences")
	errFetchingRecommendations = errors.New("error getting recommendations")
	errIdentifyingUnusedAttrs  = errors.New("error identifying unused attributes")
	errMakingRequest           = errors.New("error making request")
	errDecodingResponse        = errors.New("error decoding response")
	errDecodingHcl             = errors.New("error decoding HCL in file")
)
