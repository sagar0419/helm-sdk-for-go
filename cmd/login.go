package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/registry"
)

var chart = ChartSpec{}
var repo = RepoCreds{}

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		logrus.Info("Trying to login to the Helm Registry.")
		// Login Harbor registry
		err := HarborLogin(chart, repo, true)
		// Printing Error if any
		if err != nil {
			logrus.Fatal("logging fail \n", err)
			return err
		} else {
			logrus.Info("Logged in")
			return nil
		}
	},
}

// Harbor registry Login Func
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

func init() {
	rootCmd.AddCommand(loginCmd)
	LoginCmdFlags(loginCmd)
}

func LoginCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&chart.Repository, "harbor-url", "", "Specify Harbor URL/Hostname to used to login into Harbor Helm registry")
	cmd.Flags().StringVar(&repo.Username, "user-name", "", "Specify Harbor username to used to login into Harbor Helm registry")
	cmd.Flags().StringVar(&repo.Password, "password", "", "Specify Harbor password to used to login into Harbor Helm registry")
	enforceFlags := []string{
		"harbor-url",
		"user-name",
		"password",
	}

	for _, flagName := range enforceFlags {
		cmd.MarkFlagRequired(flagName)
	}
}
