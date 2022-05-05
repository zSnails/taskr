$GIT_COMMIT = git rev-parse --short HEAD
$GIT_USERNAME = git config --get user.name
$MODULE = "github.com/zSnails/taskr"
$PROGRAM_VERSION = "2.1.0"

$LD_FLAGS = "-s -w -X '$($MODULE)/internal/command.CommitHash=$($GIT_COMMIT)' -X '$($MODULE)/internal/command.BuildUser=$($GIT_USERNAME)' -X '$($MODULE)/internal/command.Version=$($PROGRAM_VERSION)'"
echo $LD_FLAGS

go build -ldflags="$($LD_FLAGS)"
