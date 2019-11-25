package handlers

import (
	"github.com/dragonator/super-octo-waffle/types"
)

func initializePinnedRepositories(
	query types.PinnedRepositoriesGraphQL) *types.PinnedRepositories {

	repositories := []types.PinnedRepository{}
	for _, node := range query.Organization.PinnedItems.Nodes {
		repositories = append(repositories, types.PinnedRepository{
			Name:              node.Repository.Name,
			NameWithOwner:     node.Repository.NameWithOwner,
			LicenseName:       node.Repository.LicenseInfo.Name,
			ContributorsCount: node.Repository.Collaborators.TotalCount,
			ReleasesCount:     node.Repository.Releases.TotalCount,
			HEADCommitsCount:  node.Repository.HEAD.Commit.History.TotalCount,
			BranchesCount:     node.Repository.Refs.TotalCount,
		})
	}

	return &types.PinnedRepositories{
		TotalCount:   query.Organization.PinnedItems.TotalCount,
		Repositories: repositories,
	}
}

func initializeRepository(
	query types.RepositoryGraphQL) *types.Repository {

	commits := []types.Commit{}
	commitHistory := query.Repository.DefaultBranchRef.Target.Commit.History
	for _, commit := range commitHistory.Nodes {
		commits = append(commits, types.Commit{
			Message: commit.MessageHeadline,
			Author:  commit.Author.Name,
			Date:    commit.AuthoredDate,
			Hash:    commit.Oid,
		})
	}

	return &types.Repository{
		Name:              query.Repository.Name,
		NameWithOwner:     query.Repository.NameWithOwner,
		Readme:            query.Repository.Readme.Blob.Text,
		PackageJSON:       query.Repository.PackageJSON.Blob.Text,
		DefaultBranchName: query.Repository.DefaultBranchRef.Name,
		Commits:           commits,
	}
}
