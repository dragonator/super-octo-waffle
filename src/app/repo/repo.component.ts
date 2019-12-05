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

  public selectTab(tabName) {
    var tabs = document.getElementsByClassName("tab");
    for (var i = 0; i < tabs.length; i++) {
      tabs[i].style.display = "none";
    }
    document.getElementById(tabName).style.display = "block";
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
