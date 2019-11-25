package handlers

import (
	"fmt"
	"io/ioutil"
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

	client := createAuthorizedGithubClient(context)
	query := types.PinnedRepositoriesGraphQL{}
	err := client.Query(context, &query, variables)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	pinnedRepositories := initializePinnedRepositories(query)
	pinnedRepositories.OrganizationName = context.Param("organization")

	context.JSON(http.StatusOK, pinnedRepositories)
}

func DownloadCommitPatchHandler(context *gin.Context) {
	req, err := prepareCommitPatchRequest(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	client := &http.Client{}

	resp, req_err := client.Do(req)
	if req_err != nil {
		context.JSON(http.StatusInternalServerError, req_err)
		return
	}

	body, read_err := ioutil.ReadAll(resp.Body)
	if read_err != nil {
		context.JSON(http.StatusInternalServerError, read_err)
		return
	}

	context.JSON(http.StatusOK, string(body))
}

func FetchRepositoryDataHandler(context *gin.Context) {
	variables := map[string]interface{}{
		"organization": githubv4.String(context.Param("organization")),
		"repository":   githubv4.String(context.Param("repository")),
	}

	client := createAuthorizedGithubClient(context)
	query := types.RepositoryGraphQL{}
	err := client.Query(context, &query, variables)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	repository := initializeRepository(query)

	context.JSON(http.StatusOK, repository)
}

func createAuthorizedGithubClient(context *gin.Context) *githubv4.Client {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	return githubv4.NewClient(oauth2.NewClient(context, tokenSource))
}

func prepareCommitPatchRequest(context *gin.Context) (*http.Request, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits/%s",
		context.Param("organization"),
		context.Param("repository"),
		context.Param("sha"))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.github.VERSION.patch")

	return req, err
}
