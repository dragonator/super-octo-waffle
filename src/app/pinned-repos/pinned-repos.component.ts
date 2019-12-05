import { Component, OnInit, OnDestroy } from '@angular/core';
import { ActivatedRoute }               from '@angular/router';

import { API }         from '@app/_services/api.service';
import { PinnedRepos } from '@app/_models';

@Component({
  selector:    'app-pinned-repos',
  templateUrl: './pinned-repos.component.html',
  styleUrls:  ['./pinned-repos.component.scss']
})
export class PinnedReposComponent implements OnInit, OnDestroy {
  public pinnedRepos: PinnedRepos;
  private sub: any;

  constructor(private route: ActivatedRoute, private api: API) { }

  ngOnInit() {
    this.sub = this.route.params.subscribe(params => {
      this.api.fetchPinnedRepos(params['organization'])
        .subscribe(data => {
          this.pinnedRepos = data;
        })
    });
  }

  ngOnDestroy() {
    this.sub.unsubscribe();
  }
}
