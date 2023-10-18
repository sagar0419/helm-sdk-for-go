package cmd

import (
	"cobra/harbor"
	"os/exec"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// hCred represents the harbor command line options
var hCred = ArtifactRegistry{}

// harborCmd represents the harbor command
var harborPushCmd = NewHarborCmd()

func NewHarborCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "harborPush",
		Short: "Use harborPush to push your Helm Charts to the Harbor Registry.",

		Run: func(cmd *cobra.Command, args []string) {

			// Checking helm
			harbor.CheckingHelm()
			// Harbor Login
			loginHarbor()
			// Packaging Helm Chart
			harbor.HarborPackageHelm(hCred.ChartPath)
			// Pushing Helm Chart
			harbor.HarborPushHelm(hCred.ChartName, hCred.ProjectName)
		},
	}
}

func loginHarbor() {
	command := "helm"
	args := []string{"registry", "login", hCred.HarborURL, "--username", hCred.Username, "--password", hCred.Password}
	cmd := exec.Command(command, args...)
	err := cmd.Run()
	if err != nil {
		logrus.Error("Failed to Login to Harbor \n", err)
		return
	} else {
		logrus.Info("Harbor login successful \n")
	}
}

func init() {
	rootCmd.AddCommand(harborPushCmd)
	HarborPushCmdFlags(harborPushCmd)
}
func HarborPushCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&hCred.HarborURL, "harbor-url", "", "Specify Harbor URL/Hostname to used to login into Harbor Helm registry")
	cmd.Flags().StringVar(&hCred.Username, "user-name", "", "Specify Harbor username to used to login into Harbor Helm registry")
	cmd.Flags().StringVar(&hCred.Password, "password", "", "Specify Harbor password to used to login into Harbor Helm registry")
	cmd.Flags().StringVar(&hCred.ChartName, "chartName", "", "Specify ChartName/Package Name that you want to push to Harbor. Please use the same name that you have used in Chart.yaml.")
	cmd.Flags().StringVar(&hCred.ProjectName, "projectName", "", "Specify the Name of the project where you want to store your Helm Chart/Package.")
	cmd.Flags().StringVar(&hCred.ChartPath, "chartPath", "./", "Specify Path of the directory where Helm chart is stored. By default code will use the current directory.")

	enforceFlags := []string{
		"harbor-url",
		"user-name",
		"password",
		"chartName",
		"projectName",
	}

	for _, flagName := range enforceFlags {
		cmd.MarkFlagRequired(flagName)
	}
}
