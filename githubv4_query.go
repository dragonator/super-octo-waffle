package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type repository struct {
	Name        string
	LicenseInfo struct {
		Name string
	}
}

type basic_query struct {
	Organization struct {
		PinnedItems struct {
			TotalCount int32
			Edges      []struct {
				Node struct {
					Repository repository `graphql:"... on Repository"`
				}
			}
		} `graphql:"pinnedItems(first: 10, types: [REPOSITORY, GIST])"`
	} `graphql:"organization(login: $organization)"`
}

func main() {
	token_source := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), token_source)
	client := githubv4.NewClient(httpClient)
	query := basic_query{}

	variables := map[string]interface{}{
		"organization": githubv4.String("vmware"),
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		panic(err)
	}

	fmt.Println(query.Organization.PinnedItems.TotalCount)
	fmt.Println(query.Organization.PinnedItems.Edges)
}
