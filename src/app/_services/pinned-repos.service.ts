import { Injectable }  from '@angular/core';
import { HttpClient }  from '@angular/common/http';
import { environment } from '@environments/environment';

@Injectable()
export class PinnedReposService {
  constructor(private httpClient: HttpClient) {}

  getPinnedReposList(pinnedRepos: PinnedRepos) {
    return this.httpClient.get(environment.gateway + '/pinnedRepos/' + pinnedRepos.organizationName);
  }
}
