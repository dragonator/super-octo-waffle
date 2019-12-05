export class Commit {
  Message: string;
  Author:  string;
  Date:    string;
  Hash:    string;
}

export class Repo {
  Name:              string;
  NameWithOwner:     string;
  Readme:            string;
  PackageJSON:       string;
  DefaultBranchName: string;
  Commits:         Commit[];
}
