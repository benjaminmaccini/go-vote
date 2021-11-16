# go-vote

_Note that this is in pre-alpha. If you want to get involved reach out to me_

A command line application for handling elections. Written in Go.

Protocols are derived mainly from the contents and footnotes of [this series](https://en.wikipedia.org/wiki/Electoral_system)
on electoral systems.

# Contributing

Raise an issue or a pull-request or reach out to me about something.

## Adding/Editing commands
Uses [Cobra](https://github.com/spf13/cobra). After adding/editing run `go install`,
make sure `alias go-vote="$GOPATH/bin/go-vote"` is also set.

## Testing

- `curl -X POST -H "Content-Type: application/json" -d '{"candidate": {"name": "bob"}, "value": 1}' 127.0.0.1:1234/electionId`
- `curl -X GET 127.0.0.1:1234/6ad7c653-edb3-465b-9a76-a3cf82a4212a/results`

