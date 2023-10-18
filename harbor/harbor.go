package harbor

import (
	"os/exec"

	"github.com/sirupsen/logrus"
)

func CheckingHelm() {
	helmInstall := "helm"
	cmd := exec.Command(helmInstall, "version")
	err := cmd.Run()
	if err != nil {
		logrus.Fatal("Helm is not available on the machine or an error occurred. \n")
		return
	} else {
		logrus.Info("Found helm on the machine. \n")
	}
}

func HarborPackageHelm(chartPath string) {
	command := "helm"
	args := []string{"package", chartPath}
	cmd := exec.Command(command, args...)
	err := cmd.Run()
	if err != nil {
		logrus.Fatal("Helm Chart not found on the given path. \n")
		return
	} else {
		logrus.Info("Packaging the Helm chart. \n")
	}
}

func HarborPushHelm(chartName, projectName string) {
	command := "helm"
	args := []string{"push", chartName, projectName}
	cmd := exec.Command(command, args...)
	err := cmd.Run()
	if err != nil {
		logrus.Fatal("Helm ChartName or the ProjectName is not correct. \n")
		return
	} else {
		logrus.Info("Pushing the Helm chart to the Harbor Project. \n")
	}
}
