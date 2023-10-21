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

var Dir string

// packageCmd represents the package command
var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "Use this command to 'zip/Create Package' the Helm folder",

	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("package called")

		path, err := login.ChartPackage(Dir)
		if err != nil {
			logrus.Fatal("Error occured while packaging helm: \n", err)
			return err
		} else {
			logrus.Info("Packaged helm chart path: \n", path)
			return nil
		}
	},
}

func init() {
	rootCmd.AddCommand(packageCmd)
	packageCmdFlags(packageCmd)
}

func packageCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&Dir, "dir", "", "Specify Directory Path where Chart.yaml file of Helm chart is present.")
	enforceFlags := []string{
		"dir",
	}
	for _, flag := range enforceFlags {
		cmd.MarkFlagRequired(flag)
	}
}
