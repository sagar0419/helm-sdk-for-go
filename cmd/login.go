/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
)

var chart = ChartSpec{}

var repo = RepoCreds{}

// var settings *cli.EnvSettings

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")

		login()
	},
}

// func login(spec ChartSpec, repo RepoCreds, insecure bool) error {
// 	actionConfig := new(action.Configuration)
// 	settings = cli.New()
// 	if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), os.Getenv("HELM_DRIVER"), debug); err != nil {
// 		log.Fatal(err)
// 	}

// 	opts := action.WithInsecure(insecure)
// 	debug()
// 	loginClient := action.NewRegistryLogin(actionConfig)
// 	fmt.Println(loginClient)
// 	err := loginClient.Run(nil, chart.Repository, repo.Username, repo.Password, opts)
// 	if err != nil {
// 		logrus.Fatal("Unable to login to the helm registry. ", err)
// 		return err
// 	} else {
// 		logrus.Info("Logged into the helm repo.")
// 		return nil
// 	}
// }

func login() error {
	actionConfig := new(action.Configuration)
	settings := cli.New()
	if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		logrus.Fatal(err)
	}

	client := action.NewRegistryLogin(actionConfig)
	insecure := true
	opts := action.WithInsecure(insecure)

	// err := action.NewRegistryLogin(actionConfig).Run(os.Stdout, chart.Repository, repo.Username, repo.Password, opts)
	err := client.Run(os.Stdout, chart.Repository, repo.Username, repo.Password, opts)
	if err != nil {
		logrus.Fatal("logging failes", err)
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
