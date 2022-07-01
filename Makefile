GIT_COMMIT = $(shell git rev-parse --short HEAD)
GIT_USERNAME = $(shell git config --get user.name)
MODULE = github.com/zSnails/taskr
VERSION = 2.2.1

LD_FLAGS = -s -w -X '$(MODULE)/internal/command.CommitHash=$(GIT_COMMIT)' -X '$(MODULE)/internal/command.BuildUser=$(GIT_USERNAME)' -X '$(MODULE)/internal/command.Version=$(VERSION)'
TAGS = osuusergroup,netgo,sqlite_omit_load_extension

build: main.go internal/**/*.go
	go build -tags=$(TAGS) -trimpath -ldflags="$(LD_FLAGS)"

install: build
	cp ./taskr $(GOBIN)
