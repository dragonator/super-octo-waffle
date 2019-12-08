package types

type blobText struct {
	Blob struct{ Text string } `graphql:"... on Blob"`
}

type repositoryGraphQL struct {
	Name             string
	NameWithOwner    string
	LicenseInfo      struct{ Name string }
	Releases         struct{ TotalCount int32 }
	DefaultBranchRef struct{ Name string }
	//Collaborators    struct{ TotalCount int32 }
	Refs struct{ TotalCount int32 } `graphql:"refs(refPrefix: \"refs/heads/\")"`
	HEAD struct {
		Commit struct {
			History struct{ TotalCount int32 } `graphql:"history(first: 1)"`
		} `graphql:"... on Commit"`
	} `graphql:"HEAD: object(expression: \"HEAD\")"`
}

type PinnedRepositoriesGraphQL struct {
	Organization struct {
		PinnedItems struct {
			TotalCount int32
			Nodes      []struct {
				Repository repositoryGraphQL `graphql:"... on Repository"`
			}
		} `graphql:"pinnedItems(first: 10, types: [REPOSITORY, GIST])"`
	} `graphql:"organization(login: $organization)"`
}

type commitsTargetGraphQL struct {
	Commit struct {
		History struct {
			Nodes []struct {
				MessageHeadline string
				Oid             string
				AuthoredDate    string
				Author          struct{ Name string }
			}
		} `graphql:"history(first: 20)"`
	} `graphql:"... on Commit"`
}

type RepositoryGraphQL struct {
	Repository struct {
		Name             string
		NameWithOwner    string
		Readme           blobText `graphql:"readme: object(expression: \"HEAD:README.md\")"`
		PackageJSON      blobText `graphql:"package_json: object(expression: \"HEAD:package.json\")"`
		DefaultBranchRef struct {
			Name   string
			Target commitsTargetGraphQL
		}
		Refs struct {
			Nodes []struct {
				Ref struct {
					Name string
				} `graphql:"... on Ref"`
			}
		} `graphql:"refs(refPrefix: \"refs/heads/\", first: 50)"`
	} `graphql:"repository(owner: $organization, name: $repository)"`
}
