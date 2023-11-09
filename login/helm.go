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
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
)

var chart = ChartDetails{}

func ChartPackage(dir string) (string, error) {
	packageC := action.NewPackage()
	packageC.DependencyUpdate = true
	var values map[string]interface{}
	path, err := packageC.Run(dir, values)
	return path, err
}

func HelmPush(spec ChartDetails, helmChart string) error {
	actionConfig := new(action.Configuration)
	HelmPush := action.NewPushWithOpts(action.WithPushConfig(actionConfig))
	HelmPush.Settings = &cli.EnvSettings{}
	_, err := HelmPush.Run(helmChart, spec.Repository)
	return err
}
