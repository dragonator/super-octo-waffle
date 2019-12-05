export class PinnedRepo {
  Name:              string;
  NameWithOwner:     string;
  LicenseName:       string;
  ContributorsCount: number;
  ReleasesCount:     number;
  HEADCommitsCount:  number;
  BranchesCount:     number;
}

export class PinnedRepos {
  OrganizationName: string;
  TotalCount: number;
  Repositories: PinnedRepo[];
}
