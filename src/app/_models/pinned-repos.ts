export class PinnedRepo {
	name:              string;
	nameWithOwner:     string;
	licenseName:       string;
	contributorsCount: number;
	releasesCount:     number;
	HEADCommitsCount:  number;
	branchesCount:     number;
}

export class PinnedRepos {
  organizationName: string;
  totalCount: number;
  repositories: Array<PinnedRepo>;
}
