export class Token {
  access_token: string;
  id_token:     string;
  expires_at:   string;

  constructor(accessToken, idToken, expiresAt) {
    this.access_token = accessToken;
    this.id_token     = idToken;
    this.expires_at   = expiresAt;
  }
}

export class User {
  username:  string;
  //user_id:   string;
  token?:    Token;

  constructor(username, token) {
    this.username = username;
    //this.user_id  = user_id;
    this.token    = token;
  }
}
