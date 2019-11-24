package types

type PinnedRepositories struct {
	OrganizationName string
	TotalCount       int32
	Repositories     []struct {
		Name              string
		NameWithOwner     string
		LicenseName       string
		ContributorsCount int32
		ReleasesCount     int32
		DefaultBranchName string
		CommitsCount      int32
	}
}

type Commit struct {
	Message      string
	Author       string
	AuthoredDate string
}

type Repository struct {
	Name          string
	NameWithOwner string
	Readme        string
	PackageJSON   string
	Commits       []Commit
}
