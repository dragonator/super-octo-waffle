package types

type Repository struct {
	Name        string
	LicenseInfo struct {
		Name string
	}
}

type BasicQuery struct {
	Organization struct {
		PinnedItems struct {
			TotalCount int32
			Edges      []struct {
				Node struct {
					Repository Repository `graphql:"... on Repository"`
				}
			}
		} `graphql:"pinnedItems(first: 10, types: [REPOSITORY, GIST])"`
	} `graphql:"organization(login: $organization)"`
}
