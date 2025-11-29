/*
Copyright 2025 aushacker.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"os"

	"github.com/aushacker/app-generator-plugin/api/types"
	"sigs.k8s.io/kustomize/kyaml/fn/framework"
	"sigs.k8s.io/kustomize/kyaml/fn/framework/command"
	"sigs.k8s.io/kustomize/kyaml/kio"

	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func main() {
	config := new(types.AppGenerator)
	fn := func(items []*yaml.RNode) ([]*yaml.RNode, error) {
		items = append(items, makeNode(config.Kind))
		return items, nil
	}
	p := framework.SimpleProcessor{Config: config, Filter: kio.FilterFunc(fn)}
	cmd := command.Build(p, command.StandaloneDisabled, false)
	command.AddGenerateDockerfile(cmd)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func makeNode(kind string) *yaml.RNode {
	rn, _ := yaml.Parse(fmt.Sprintf(`
apiVersion: v1
kind: %s
metadata:
  name: foo`, kind))
	return rn
}
