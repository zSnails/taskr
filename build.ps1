$GIT_COMMIT = git rev-parse --short HEAD
$GIT_USERNAME = git config --get user.name
$MODULE = "github.com/zSnails/taskr"
$PROGRAM_VERSION = "2.2.1"

$env:CGO_ENABLED = 1
$LD_FLAGS = "-s -w -X '$($MODULE)/internal/command.CommitHash=$($GIT_COMMIT)' -X '$($MODULE)/internal/command.BuildUser=$($GIT_USERNAME)' -X '$($MODULE)/internal/command.Version=$($PROGRAM_VERSION)'"
$TAGS = "osusergroup,netgo,sqlite_omit_load_extension"
go build -tags=$TAGS -trimpath -ldflags="$($LD_FLAGS)"
