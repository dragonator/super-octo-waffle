import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../environments/environment';

@Injectable()
export class PinnedRepositoriesService {
  constructor(private httpClient: HttpClient) {}

  getPinnedRepositoriesList(pinnedRepositories: PinnedRepositories) {
    return this.httpClient.get(environment.gateway + '/pinnedItems/' + PinnedRepositories.organizationName);
  }
}

type PinnedRepository struct {
	name              string;
	nameWithOwner     string;
	licenseName       string;
	contributorsCount int;
	releasesCount     int;
	HEADCommitsCount  int;
	branchesCount     int;
}

export class PinnedRepositories {
  organizationName: string;
  totalCount: int;
  repositories: Array<PinnedRepository>;
}
