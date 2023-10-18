package login

type RepoCreds struct {
	Username string
	Password string
}

type ChartSpec struct {
	// Repository: Helm repository URL, required if ChartSpec.URL not set
	Repository string `json:"repository,omitempty"`
	// Name of Helm chart, required if ChartSpec.URL not set
	Name string `json:"name,omitempty"`
	// Name of Service
	ServiceName string `json:"simpleName,omitempty"`
	// AWS OCI Name compliant
	RepoName string `json:"awsRepoName,omitempty"`
	// Version to override the helm charts version to
	VersionOverride string `json:"versionOverride,omitempty"`
	// Validate if Version exists
	VersionExists bool `json:"versionExists,omitempty"`
	// URL to chart package (typically .tgz), optional and overrides others fields in the spec
	URL string `json:"url,omitempty"`
	// Docker Image Tag
	Tag string `json:"tag,omitempty"`
	// Tag paths to override in values.yaml
	TagPaths string `json:"tagPaths,omitempty"`
	// PullSecretRef is reference to the secret containing credentials to helm repository
	// PullSecretRef xpv1.SecretReference `json:"pullSecretRef,omitempty"`
}

const (
	errFailedToCheckIfLocalChartExists = "failed to check if cached chart file exists"
	errFailedToPullChart               = "failed to pull chart"
	errFailedToLoadChart               = "failed to load chart"
	errUnexpectedDirContentTmpl        = "expected 1 .tgz chart file, got [%s]"
	errFailedToParseURL                = "failed to parse URL"
	errFailedToLogin                   = "failed to login to registry"
	errUnexpectedOCIUrlTmpl            = "url not prefixed with oci://, got [%s]"
)
