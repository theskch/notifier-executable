# Notifier executable
Notifier executable is a small example project demonstrating how notifier is to be used. After every interval it reads `stdin` from start and sends the messages to the passed url. Url and interval are configurable using flags. `url` is mandatory while `interval` is set to 5s by default.

## Usage
Build the project using `go build` command.
See usage using `notify --help`