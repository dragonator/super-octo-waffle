package handlers

import (
	"encoding/json"
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

	client := githubv4.NewClient(createAuthorizedClient(context, os.Getenv("GITHUB_TOKEN")))
	query := types.PinnedRepositoriesQuery{}
	err := client.Query(context, &query, variables)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	}

	repositories := []types.PinnedRepository{}
	for _, node := range query.Organization.PinnedItems.Nodes {
		repositories = append(repositories, types.PinnedRepository{
			Name:              node.Repository.Name,
			NameWithOwner:     node.Repository.NameWithOwner,
			LicenseName:       node.Repository.LicenseInfo.Name,
			ContributorsCount: 0, //TODO
			ReleasesCount:     node.Repository.Releases.TotalCount,
			DefaultBranchName: node.Repository.DefaultBranchRef.Name,
			CommitsCount:      node.Repository.HEAD.Commit.History.TotalCount,
		})
	}

	pinnedRepositories := types.PinnedRepositories{
		OrganizationName: context.Param("organization"),
		TotalCount:       query.Organization.PinnedItems.TotalCount,
		Repositories:     repositories,
	}

	context.JSON(http.StatusOK, pinnedRepositories)
}

func DownloadCommitPatchHandler(context *gin.Context) {
	req := prepareCommitPatchRequest(context)
	client := &http.Client{}

	resp, req_err := client.Do(req)
	if req_err != nil {
		context.JSON(http.StatusInternalServerError, req_err)
	}

	body, read_err := ioutil.ReadAll(resp.Body)
	if read_err != nil {
		context.JSON(http.StatusInternalServerError, read_err)
	}

	//patchJSON, marshal_err := json.Marshal(body)
	//if marshal_err != nil {
	//		context.JSON(http.StatusInternalServerError, marshal_err)
	//}

	context.JSON(http.StatusOK, string(body))
}

func FetchRepositoryDataHandler(context *gin.Context) {
	commits := []types.Commit{}
	for _, commit := range requestCommitsInJSON(context) {
		commits = append(commits, types.Commit{
			Message:      commit.Commit.Message,
			Author:       commit.Committer.Name,
			AuthoredDate: commit.Committer.Date,
		})
	}

	variables := map[string]interface{}{
		"organization": githubv4.String(context.Param("organization")),
		"repository":   githubv4.String(context.Param("repository")),
	}

	client := githubv4.NewClient(createAuthorizedClient(context, os.Getenv("GITHUB_TOKEN")))
	query := types.RepositoryQuery{}
	err := client.Query(context, &query, variables)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	}

	repository := types.Repository{
		Name:          query.Repository.Name,
		NameWithOwner: query.Repository.NameWithOwner,
		Readme:        query.Repository.Readme.Blob.Text,
		PackageJSON:   query.Repository.PackageJSON.Blob.Text,
		Commits:       commits,
	}

	context.JSON(http.StatusOK, repository)
}

func createAuthorizedClient(context *gin.Context, token string) *http.Client {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	return oauth2.NewClient(context, tokenSource)
}

func requestCommitsInJSON(context *gin.Context) []types.UnmarshalCommitScheme {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits",
		context.Param("organization"),
		context.Param("repository"))

	resp, response_err := http.Get(url)
	if response_err != nil {
		context.JSON(http.StatusInternalServerError, response_err)
	}

	body, read_err := ioutil.ReadAll(resp.Body)
	if read_err != nil {
		context.JSON(http.StatusInternalServerError, read_err)
	}

	var unmarshalled_commits []types.UnmarshalCommitScheme
	unmarshal_err := json.Unmarshal([]byte(body), &unmarshalled_commits)
	if unmarshal_err != nil {
		context.JSON(http.StatusInternalServerError, unmarshal_err)
	}

	return unmarshalled_commits
}

func prepareCommitPatchRequest(context *gin.Context) *http.Request {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits/%s",
		context.Param("organization"),
		context.Param("repository"),
		context.Param("sha"))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	}
	req.Header.Add("Accept", "application/vnd.github.VERSION.patch")

	return req
}
