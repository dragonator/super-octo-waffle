import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../environments/environment';

@Injectable()
export class PinnedReposService {
  constructor(private httpClient: HttpClient) {}

  getPinnedReposList(pinnedRepos: PinnedRepos) {
    return this.httpClient.get(environment.gateway + '/pinnedItems/' + pinnedRepos.organizationName);
  }
}

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
