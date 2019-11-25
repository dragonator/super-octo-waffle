package types

type PinnedRepository struct {
	Name              string
	NameWithOwner     string
	LicenseName       string
	ContributorsCount int32
	ReleasesCount     int32
	DefaultBranchName string
	CommitsCount      int32
}

type PinnedRepositories struct {
	OrganizationName string
	TotalCount       int32
	Repositories     []PinnedRepository
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
