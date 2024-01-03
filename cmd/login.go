/*
Copyright © 2023 Sagar Parmar <sagar.rajput27@live.com>
*/
package cmd

import (
	"helm-sdk-for-go/login"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/registry"
)

var chart = login.ChartDetails{}
var repo = login.RepoDetails{}

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Use this subcommand to log into Helm private registry.",
	RunE: func(cmd *cobra.Command, args []string) error {
		logrus.Info("Trying to login to the Helm Registry.")

		// Login Helm registry
		err := HelmLogin(chart, repo)
		// Printing Error if any
		if err != nil {
			logrus.Fatal("Logging fail.\n", err)
			return err
		} else {
			logrus.Info("Logged in ")
			return nil
		}
	},
}

func HelmLogin(login.ChartDetails, login.RepoDetails) error {
	conf, err := registry.NewClient()
	if err != nil {
		return err
	}
	actionConfig := new(action.Configuration)
	actionConfig.RegistryClient = conf
	login := action.NewRegistryLogin(actionConfig)
	insecure := true
	opts := action.WithInsecure(insecure)
	err = login.Run(os.Stdout, chart.Repository, repo.Username, repo.Password, opts)
	return err
}
func init() {
	rootCmd.AddCommand(loginCmd)
	LoginCmdFlags(loginCmd)
}

func LoginCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&chart.Repository, "helm-url", "sagar.parmar.io", "Specify Helm registry URL/Hostname to used to login into Helm registry")
	cmd.Flags().StringVar(&repo.Username, "user-name", "", "Specify Helm registry username to used to login into Helm registry.")
	cmd.Flags().StringVar(&repo.Password, "password", "", "Specify Helm registry password to used to login into Helm registry.")
	enforceFlags := []string{
		"helm-url",
		"user-name",
		"password",
	}

	for _, flagName := range enforceFlags {
		cmd.MarkFlagRequired(flagName)
	}
}
