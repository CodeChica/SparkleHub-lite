# Sparkles

This is a web application that allows you to say nice things about other people
called "Sparkles".

A single HTML file can be found in `./public/index.html` that is the home page
for the web application. A live version of this application can be found
[here][production].

## Getting Started

1. Install [Golang][golang].
2. Run `./script/server` from a terminal.
3. Open a browser to [http://localhost:8080](http://localhost:8080).

## HTTP API

### POST /sparkles.json

Use this endpoint to create a new sparkle.

```bash
$ ./script/sparkle @monalisa for helping me with my project!
HTTP/1.1 201 Created
Access-Control-Allow-Origin: *
Content-Type: application/json
Date: Wed, 13 Oct 2021 15:52:49 GMT
Content-Length: 68

{"sparklee":"@monalisa","reason":"for helping me with my project!"}
```

### GET /sparkles.json

Use this endpoint to get a list of all the sparkles.

```bash
$ ./script/sparkles | jq '.'
[
  {
    "sparklee": "@monalisa",
    "reason": "for helping me with my project!"
  }
]
```

[golang]: https://golang.org/doc/install
[production]: https://sparklehub.herokuapp.com
