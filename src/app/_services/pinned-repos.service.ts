import { Injectable }  from '@angular/core';
import { HttpClient }  from '@angular/common/http';

import { environment } from '@environments/environment';
import { PinnedRepos } from '@app/_models';

@Injectable()
export class PinnedReposService {
  constructor(private httpClient: HttpClient) {}

  getPinnedReposList(owner: string) {
    return this.httpClient.get<PinnedRepos>(environment.gateway + '/api/pinnedRepos/' + owner);
  }
}
