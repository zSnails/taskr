$GIT_COMMIT = git rev-list -1 HEAD
$GIT_USERNAME = git config --get user.name
$PROGRAM_VERSION = "0.2.1"

go build `
    -ldflags `
        "-s -w -X main.gitCommit=$($GIT_COMMIT) -X main.buildUser=$($GIT_USERNAME) -X main.programVersion=$($PROGRAM_VERSION)"
