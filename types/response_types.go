package types

type PinnedRepository struct {
	Name              string
	NameWithOwner     string
	LicenseName       string
	ContributorsCount int32
	ReleasesCount     int32
	HEADCommitsCount  int32
	BranchesCount     int32
}

type PinnedRepositories struct {
	OrganizationName string
	TotalCount       int32
	Repositories     []PinnedRepository
}

type Commit struct {
	Message string
	Author  string
	Date    string
	Hash    string
}

type Repository struct {
	Name              string
	NameWithOwner     string
	Readme            string
	PackageJSON       string
	DefaultBranchName string
	Commits           []Commit
}
