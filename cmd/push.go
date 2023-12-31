/*
Copyright © 2023 Sagar Parmar <sagar.rajput27@live.com>
*/
package cmd

import (
	"fmt"
	"helm-sdk-for-go/login"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var url = login.ChartDetails{}
var chartName string

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Use this command to push the helm chart on to the helm registry",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("push called")
		err := login.HelmPush(url, chartName)
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
	cmd.Flags().StringVar(&url.Repository, "url", "oci://sagar.parmar.io/demo", "Use registry OCI URL to push the helm chart on the Helm registry")
	cmd.Flags().StringVar(&chartName, "chartName", "", "Provide the name of the chart that you want to push to the helm repo.")

	requiredFlags := []string{
		"chartName",
	}
	for _, flags := range requiredFlags {
		cmd.MarkFlagRequired(flags)
	}

}
