# go-vote

A command line application for handling elections.

Protocols are derived mainly from the contents and footnotes of [this series](https://en.wikipedia.org/wiki/Electoral_system) on electoral systems.

**Example:**
```golang
go-vote serve --candidate alice --candidate bob --protocol simpleMajority
```

# Contributing

This repository is officially hosted on [Sourcehut](https://git.sr.ht/~bmaccini/go-vote), mirrored on [Github](https://github.com/benjaminmaccini/go-vote). All issue
tracking is handled on Sourcehut.

## Quickstart

Install:
- [go](https://go.dev/dl/)
- [golangci-lint](https://golangci-lint.run/) - CI linter, but local
- [python](https://www.python.org/downloads/)
- [sqlite](https://www.sqlite.org/index.html) - Primary persistent storage
- [sqlc](https://docs.sqlc.dev/en/stable/overview/install.html) - For generating Go structs from SQL
- [hurl](https://hurl.dev/docs/installation.html) - Tests for server routes
- [rust-lang](https://www.rust-lang.org/tools/install)
- [openapi-to-hurl](https://crates.io/crates/openapi-to-hurl) - Auto-generates `hurl` tests
	- Note: You may need to install `libxml2-dev`
- [dbmate](https://github.com/amacneil/dbmate) - Tool for handling migrations
- [entr](https://eradman.com/entrproject/) - For the `make watch` command
- [ag](https://github.com/ggreer/the_silver_searcher) - For the `make watch` command

Type `make` to list available dev targets.

## Contributing

### Workflows

Different things that need to be developed and how to do them

1. SQL Queries
	a. Edit `db/query.sql` and add it.
	b. Run `make db/generate`
2. SQL Migrations
	a. `dbmate new name_of_my_migration`
	b. Edit file in `db/migrations/<whatever_new_file_was_created>.sql`
	c. `make db/migrate`
3. API Endpoint
	a. Add API request and response payload to `pkg/web/api.go`
	b. Add a new handler to `pkg/web/handlers.go` and its path under `pkg/web/server.go`
	c. Update the OpenAPI documentation (it's helpful to use an LLM for this) at `docs/schema.yaml`
	d. Run the documentation server to confirm.
	e. Generate and run the API tests using `make test/http`
4. Adding a protocol
	a. Follow `pkg/protocol/simple.go` as a template.
	b. Make sure to add tests and add to the `ProtocolCommandMap`.

### Style

- General
	- If possible, don't add new packages. Of course, if it's an extra convienent package, use it.
	- Core library of voting protocols that persist the results to a SQLite database, wrapped by a web service. Additional clients
	can "plug-in" to this via their own library in `/pkg`.
	- Err on the side of global utilities that other parts of the library can use
- Testing
	- Unit tests shouldn't be redundant. Rather than testing like an onion, test like sausage links.
	- HTTP tests are generated from the Open API spec
- DB
	- The core data model should be protocol agnostic and bashed into whatever data structure makes the most
	sense for calculating results.
	- Primary keys should be UUIDs. Based on what you [want](https://web.archive.org/web/20240914010055/https://www.ntietz.com/blog/til-uses-for-the-different-uuid-versions/) to do with them.
	UUID v7 is a good choice.

## TODO

- [ ] update API tests
- [ ] Ranked choice
- [ ] Multi member constituency methods
- [ ] Concurrent elections

### TODO Later

- [ ] Add [cryptographic tallying](https://web.archive.org/web/20200331140611/http://security.hsr.ch/msevote/seminar-papers/HS09_Homomorphic_Tallying_with_Paillier.pdf)
- [ ] Digital Signatures (multisignature?)
- [ ] Verifiable server state (version the same between server-client? Non-malicious server accepting votes?)
- [ ] Audit tables for SQLite
