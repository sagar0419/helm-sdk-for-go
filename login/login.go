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
	"net/url"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/registry"
)

type HelmClient struct {
	pullClient    *action.Pull
	packageClient *action.Package
	pushClient    *action.Push
	loginClient   *action.RegistryLogin
	installChart  *action.Install
}

func (hc *HelmClient) Login(spec ChartSpec, creds RepoCreds, insecure bool) error {
	ociURL := spec.URL
	if spec.URL == "" {
		ociURL = spec.Repository
	}
	if !registry.IsOCI(ociURL) {
		return nil
	}
	parsedURL, err := url.Parse(ociURL)
	if err != nil {
		return errors.Wrap(err, errFailedToParseURL)
	}
	var out strings.Builder

	opts := action.WithInsecure(insecure)
	err = hc.loginClient.Run(&out, parsedURL.Host, creds.Username, creds.Password, opts)

	logrus.Info(out.String())

	return errors.Wrap(err, errFailedToLogin)
}
