package main

import (
  "fmt"
  "os"
  "context"

  "github.com/shurcooL/githubv4"
  "golang.org/x/oauth2"
)

type repository struct {
  Name string
  LicenseInfo struct {
    Name string
  }
}

type basic_query struct {
  Organization struct {
    PinnedItems struct {
      TotalCount int32
      Edges []struct {
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
  client     := githubv4.NewClient(httpClient)
  query      := basic_query{}

  variables := map[string]interface{} {
    "organization": githubv4.String("vmware"),
  }

  err := client.Query(context.Background(), &q, variables)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  fmt.Println(q.Organization.PinnedItems.TotalCount)
  fmt.Println(q.Organization.PinnedItems.Edges)
}
