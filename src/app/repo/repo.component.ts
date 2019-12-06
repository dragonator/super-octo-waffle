import { Component, OnInit, OnDestroy } from '@angular/core';
import { ActivatedRoute }               from '@angular/router';

import { API }  from '@app/_services/api.service';
import { Repo } from '@app/_models'

@Component({
  selector:    'app-repo',
  templateUrl: './repo.component.html',
  styleUrls:  ['./repo.component.scss']
})
export class RepoComponent implements OnInit, OnDestroy {
  public owner: string;
  public repo: Repo;
  private sub: any;

  constructor(private route: ActivatedRoute, private api: API) { }

  public downloadPatch(repoWithOwner: string, commit: string) {
    this.api.fetchCommitPatch(repoWithOwner, commit)
      .subscribe((response: any) => {
        let filename = 'patch_for_commit_' + commit;
        let blob = new Blob([response], {type: response.type});
        let downloadLink = document.createElement('a');
        downloadLink.href = window.URL.createObjectURL(blob);
        downloadLink.setAttribute('download', filename);
        document.body.appendChild(downloadLink);
        downloadLink.click();
      }
    )
  }

  ngOnInit() {
    this.sub = this.route.params.subscribe(params => {
      this.api.fetchRepo(params['organization'], params['repo'])
        .subscribe(data => {
          this.owner = params['organization'];
          this.repo  = data;
        })
    });
  }

  ngOnDestroy() {
    this.sub.unsubscribe();
  }

}
