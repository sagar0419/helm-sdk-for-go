package cmd

import (
	"helm-sdk-for-go/login"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var chart = login.ChartSpec{}
var repo = login.RepoCreds{}

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		logrus.Info("Trying to login to the Helm Registry.")

		// Login Harbor registry
		err := login.HarborLogin(login.ChartSpec(chart), login.RepoCreds{Username: repo.Username, Password: repo.Password}, true)
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
