package types

type refs struct {
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
	Refs             refs `graphql:"refs(first: 100, refPrefix: \"refs/heads/\")"`
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
