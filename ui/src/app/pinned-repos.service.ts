import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../environments/environment';

@Injectable()
export class PinnedReposService {
  constructor(private httpClient: HttpClient) {}

  getPinnedReposList(pinnedRepos: PinnedRepos) {
    return this.httpClient.get(environment.gateway + '/pinnedItems/' + PinnedRepos.organizationName);
  }
}

type PinnedRepo struct {
	name              string;
	nameWithOwner     string;
	licenseName       string;
	contributorsCount int;
	releasesCount     int;
	HEADCommitsCount  int;
	branchesCount     int;
}

export class PinnedRepos {
  organizationName: string;
  totalCount: int;
  repositories: Array<PinnedRepo>;
}
