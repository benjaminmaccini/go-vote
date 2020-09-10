# go-vote

A command line application for handling elections. Written in Go.

Protocols are derived mainly from the contents and footnotes of [this series](https://en.wikipedia.org/wiki/Electoral_system)
on electoral systems.

# Contributing

Raise an issue or a pull-request or reach out to me about something.

## Adding/Editing commands
Uses [Cobra](https://github.com/spf13/cobra). After adding/editing run `go install`,
make sure `alias go-vote="$GOPATH/bin/go-vote"` is also set.

## Testing

- `make test` or `go test ./...` for testing
- `make watch` is useful for development (requires [entr](https://github.com/clibs/entr)).

