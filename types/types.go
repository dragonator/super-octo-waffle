package types

type blobText struct {
	Blob struct{ Text string } `graphql:"... on Blob"`
}

type refsCommitsCount struct {
	TotalCount int32
	Nodes      []struct {
		Target struct {
			Commit struct {
				History struct{ TotalCount int32 } `graphql:"history(first: 1)"`
			} `graphql:"... on Commit"`
		}
	}
}

type repository struct {
	Name             string
	NameWithOwner    string
	LicenseInfo      struct{ Name string }
	Releases         struct{ TotalCount int32 } `graphql:"releases(last: 5)"`
	DefaultBranchRef struct{ Name string }
	Refs             refsCommitsCount `graphql:"refs(first: 100, refPrefix: \"refs/heads/\")"`
}

type OrganizationsPinnedItems struct {
	Organization struct {
		PinnedItems struct {
			TotalCount int32
			Nodes      []struct {
				Repository repository `graphql:"... on Repository"`
			}
		} `graphql:"pinnedItems(first: 10, types: [REPOSITORY, GIST])"`
	} `graphql:"organization(login: $organization)"`
}

type RepositoryData struct {
	Repository struct {
		DefaultBranchRef struct {
			DefaultBranch struct {
				Commit struct {
					History struct {
						Nodes []struct {
							MessageHeadline string
							AuthoredDate    string
							Author          struct{ Name string }
						}
					}
				} `graphql:"... on Commit"`
			} `graphql:"default_branch: target"`
		}
		Readme      blobText `graphql:"readme: object(expression: \"HEAD:README.md\")"`
		PackageJSON blobText `graphql:"package_json: object(expression: \"HEAD:package.json\")"`
	} `graphql:"repository(owner: $organization, name: $repository)"`
}
