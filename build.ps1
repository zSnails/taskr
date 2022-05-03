$GIT_COMMIT = git rev-list -1 HEAD
$GIT_USERNAME = git config --get user.name
$PROGRAM_VERSION = "1.0.0"

go build `
    -ldflags `
        "-s -w -X main.gitCommit=$($GIT_COMMIT) -X main.buildUser=$($GIT_USERNAME) -X main.programVersion=$($PROGRAM_VERSION)"
