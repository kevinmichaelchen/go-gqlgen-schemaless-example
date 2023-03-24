PORT = 8081

# These flags can make the Makefile more robust and prevent errors from
# occurring.
# The -e flag causes the shell to exit immediately if any command returns a
# non-zero exit status (i.e. indicates an error).
# The -u flag causes the shell to treat unset variables as an error and exit
# immediately.
SHELL += -eu

RESET = \033[0m
YELLOW = \033[33m
GREEN = \033[32m
BLUE = \033[0;34m

.PHONY: help
help : Makefile
	@sed -n 's/^##//p' $<

## run                    : runs the application
.PHONY: run
run:
	PORT=$(PORT) go run main.go

## gen-server             : auto-generates GraphQL handler schema + types
.PHONY: gen-server
gen-server:
	go run github.com/99designs/gqlgen generate
