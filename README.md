#Go auth0 Request

A simple wrapper to make requests to auth0 protected APIs muche simpler. Inspired by [auth0-curl](https://github.com/Gleetr/auth0-curl) (author: [Gleetr](https://github.com/Gleetr)).
The lib used to perform great HTTP(s) requests is [goreq](https://github.com/franela/goreq) (author: [franela](https://github.com/franela)).

##Usage

```go
res, err := goauth0Request.Get{
    Domain:     "https://<auth0-subdomain>.auth0.com/oauth/ro",
    Client_id:  "<auth0 apps client id>",
    Username:   "<username, e.g. john@smith.com>",
    Password:   "<password>",
    Connection: "Username-Password-Authentication",
    Grant_type: "password",
    Scope:      "openid email",
    Uri:        "<http://protected-url.com>",
}.Do()

if err != nil {
    log.Fatal(err)
}

var jsonParsed JSONStruct
res.Body.FromJsonTo(&jsonParsed)

```

##License

MIT
