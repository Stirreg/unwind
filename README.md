# Unwind

*Unwind your day.*

## Usage

- Have go installed on your machine: https://golang.org/doc/install
- Go to your `$GOPATH`.
- Run `go get github.com/Stirreg/unwind` in your shell.
- Run program using `./bin/unwind-cli [title] [content]` e.g.: `./bin/unwind-cli "Some title" "Some long piece of content."`. You can omit the `./` if your `$GOPATH/bin` is in your path. On windows add `.exe` to the end of the program name and use back slashes instead of forward slashes.
- OR
- Run program using `./bin/unwind` and visit port 8080 on your localhost, use the first url segement to query for some content e.g.: `127.0.0.1:8080/2018-03-02`

Todo:

- [x] Set up most basic program imagineable.
- [x] Decide on system gui or web interface.
- [ ] Add interface to make this actually useful.
- [ ] Provide binaries somewhere (or host website with interface somewhere).
- [ ] Expand todo list.
