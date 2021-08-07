# Sparkles

This is a web application that allows you to say nice things about other people
called "Sparkles".

A single HTML file can be found in `./public/index.html` that is the home page
for the web application.

## Getting Started

1. Install [Golang][golang].
2. Run `./script/server` from a terminal.
3. Open a browser to [http://localhost:8080](http://localhost:8080).

## HTTP API

### POST /sparkles

Use this endpoint to create a new sparkle.

```bash
モ curl http://localhost:8080/sparkles -d 'body=@you+for+working+hard'
```

### GET /sparkles.json

Use this endpoint to get a list of all the sparkles.

```bash
モ curl http://localhost:8080/sparkles.json | jq '.'
{
  "sparkles": [
    {
      "Sparklee": {
        "Name": "@you"
      },
      "Reason": "for working hard"
    }
  ]
}
```

[golang]: https://golang.org/doc/install
