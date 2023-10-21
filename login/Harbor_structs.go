package login

type RepoDetails struct {
	Username string
	Password string
}

type ChartDetails struct {
	Repository string
	Name       string
	URL        string
	Tag        string
	TagPaths   string
}
