import { Injectable }  from '@angular/core';
import { HttpClient }  from '@angular/common/http';

import { environment }       from '@environments/environment';
import { PinnedRepos, Repo } from '@app/_models';

@Injectable()
export class API {
  constructor(private httpClient: HttpClient) {}

  public fetchPinnedRepos(owner: string) {
    return this.httpClient.get<PinnedRepos>(environment.gateway + '/api/pinnedRepos/' + owner);
  }

  public fetchRepo(owner: string, repo: string) {
    return this.httpClient.get<Repo>(environment.gateway + '/api/repo/' + owner + '/' + repo);
  }
}
