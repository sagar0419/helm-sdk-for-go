/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"helm-sdk-for-go/login"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
)

var url = login.ChartSpec{}
var chartName string

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("push called")
		err := HarborPush(url, chartName)
		if err != nil {
			logrus.Fatal("An error accoured unable to push the repo.\n")
			return err
		} else {
			logrus.Info("Chart is pushed to the helm registry. \n", url)
			return nil
		}
	},
}

func HarborPush(spec login.ChartSpec, helmChart string) error {
	actionConfig := new(action.Configuration)
	harborPush := action.NewPushWithOpts(action.WithPushConfig(actionConfig))
	harborPush.Settings = &cli.EnvSettings{}
	_, err := harborPush.Run(helmChart, spec.Repository)
	return err
}

func init() {
	rootCmd.AddCommand(pushCmd)
	helmPushFlag(pushCmd)
}
func helmPushFlag(cmd *cobra.Command) {
	cmd.Flags().StringVar(&url.Repository, "url", "ci://harbor.bender.rocks/helm", "Use registry OCI URL to push the helm chart on the Helm registry")
	cmd.Flags().StringVar(&chartName, "chartName", "", "Provide the name of the chart that you want to push to the helm repo.")

	requiredFlags := []string{
		"chartName",
	}
	for _, flags := range requiredFlags {
		cmd.MarkFlagRequired(flags)
	}

}
