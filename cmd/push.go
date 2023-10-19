/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"helm-sdk-for-go/login"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var url = login.ChartSpec{}
var chartName string

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("push called")
		err := login.HarborPush(url, chartName)
		if err != nil {
			logrus.Fatal("An error accoured unable to push the repo.\n")
			return err
		} else {
			logrus.Info("Chart is pushed to the helm registry. \n", url.Repository)
			return nil
		}
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
	helmPushFlag(pushCmd)
}
func helmPushFlag(cmd *cobra.Command) {
	cmd.Flags().StringVar(&url.Repository, "url", "oci://sagar.parmar.io/helm", "Use registry OCI URL to push the helm chart on the Helm registry")
	cmd.Flags().StringVar(&chartName, "chartName", "", "Provide the name of the chart that you want to push to the helm repo.")

	requiredFlags := []string{
		"chartName",
	}
	for _, flags := range requiredFlags {
		cmd.MarkFlagRequired(flags)
	}

}
