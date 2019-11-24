package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"

	"github.com/dragonator/super-octo-waffle/types"
)

func FetchPinnedItemsHandler(c *gin.Context) {
	organization := c.Param("organization")
	variables := map[string]interface{}{
		"organization": githubv4.String(organization),
	}

	client := createGithubClient(c)
	query := types.BasicQuery{}

	err := client.Query(c, &query, variables)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, query)
}

func createGithubClient(c *gin.Context) *githubv4.Client {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(c, tokenSource)

	return githubv4.NewClient(httpClient)
}
