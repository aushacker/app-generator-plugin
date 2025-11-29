// Copyright 2025 aushacker
// SPDX-License-Identifier: Apache-2.0

package types

type AppGenerator struct {
	ApiVersion string `yaml:"apiVersion" json:"apiVersion"`
	Kind       string `yaml:"kind" json:"kind"`
	Value      string `yaml:"value" json:"value"`
}
