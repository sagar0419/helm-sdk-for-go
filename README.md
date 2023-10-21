This repo Contains Go Code for the following things: -
- Login into a private helm.
- Create a helm package.
- Push the helm chart to the private helm repo

To run this code execute the follwoing command, it will show you all the available options :- 
`go run main.go`

Once you execute the above command you will get an output like this :-

```
You can use this command to login into private helm repo, create a package of the helm directory and push it to the helm registry.

Usage:
  helm-sdk-for-go [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  login       Use this sub command to login into Helm private registry.
  package     Use this command to 'zip/Create Package' the Helm folder
  push        Use this command to push the helm chart on to the helm registry

Flags:
  -h, --help     help for helm-sdk-for-go
  -t, --toggle   Help message for toggle

Use "helm-sdk-for-go [command] --help" for more information about a command.
```

Use your desired option to do the Login, Push and Pull Operations.