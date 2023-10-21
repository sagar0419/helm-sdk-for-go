package login

type RepoDetails struct {
	Username string
	Password string
}

type ChartDetails struct {
	Repository string
	Name       string
	URL        string
	// Docker Image Tag
	Tag string `json:"tag,omitempty"`
	// Tag paths to override in values.yaml
	TagPaths string `json:"tagPaths,omitempty"`
	// PullSecretRef is reference to the secret containing credentials to helm repository
	// PullSecretRef xpv1.SecretReference `json:"pullSecretRef,omitempty"`
}
