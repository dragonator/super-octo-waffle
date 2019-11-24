package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"

	"github.com/dragonator/super-octo-waffle/types"
)

func FetchPinnedItemsHandler(context *gin.Context) {
	variables := map[string]interface{}{
		"organization": githubv4.String(context.Param("organization")),
	}

	client := createGithubClient(context)
	query := types.OrganizationsPinnedItems{}
	sendQuery(client, context, &query, variables)
}

func FetchRepositoryDataHandler(context *gin.Context) {
	variables := map[string]interface{}{
		"organization": githubv4.String(context.Param("organization")),
		"repository":   githubv4.String(context.Param("repository")),
	}

	client := createGithubClient(context)
	query := types.RepositoryData{}
	sendQuery(client, context, &query, variables)
}

func createGithubClient(context *gin.Context) *githubv4.Client {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context, tokenSource)

	return githubv4.NewClient(httpClient)
}

func sendQuery(client *githubv4.Client, context *gin.Context, query interface{}, variables map[string]interface{}) {
	err := client.Query(context, query, variables)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	}
	context.JSON(http.StatusOK, query)
}
