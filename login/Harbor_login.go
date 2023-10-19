/*
Copyright The Helm Authors.

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

package login

import (
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/registry"
)

var chart = ChartSpec{}

var repo = RepoCreds{}

func HarborLogin(spec ChartSpec, creds RepoCreds, insecure bool) error {
	rc, err := registry.NewClient()
	if err != nil {
		return err
	}
	actionConfig := new(action.Configuration)
	actionConfig.RegistryClient = rc
	login := action.NewRegistryLogin(actionConfig)
	opts := action.WithInsecure(insecure)
	err = login.Run(os.Stdout, chart.Repository, repo.Username, repo.Password, opts)
	return err
}

func ChartPackage(dir string) (string, error) {
	packageC := action.NewPackage()
	packageC.DependencyUpdate = true
	var values map[string]interface{}
	path, err := packageC.Run(dir, values)
	return path, err
}
