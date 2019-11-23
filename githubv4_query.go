package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

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

func main() {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), tokenSource)
	client := githubv4.NewClient(httpClient)
	query := BasicQuery{}

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
