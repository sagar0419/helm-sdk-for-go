package cmd

import (
	"fmt"
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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")

		// login()
		login(chart, repo, true)
	},
}

func login(spec ChartSpec, creds RepoCreds, insecure bool) error {
	rc, err := registry.NewClient()
	if err != nil {
		return err
	}
	client := new(action.RegistryLogin)
	actionConfig := new(action.Configuration)
	actionConfig.RegistryClient = rc

	// actionConfig := new(action.Configuration)
	login := action.NewRegistryLogin(actionConfig)
	// client := new(action.RegistryLogin)
	fmt.Println(client)

	opts := action.WithInsecure(insecure)
	err = login.Run(os.Stdout, chart.Repository, repo.Username, repo.Password, opts)
	// err := client.Run(os.Stdout, chart.Repository, repo.Username, repo.Password, opts)
	if err != nil {
		logrus.Fatal("logging fail", err)
		return err
	} else {
		logrus.Info("Logged in")
		return nil
	}
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
