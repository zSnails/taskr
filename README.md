[![GPLv3 License](https://img.shields.io/badge/License-GPL%20v3-yellow.svg)](https://opensource.org/licenses/)

# Taskr

Taskr is a simple command line tool to create reminders *or tasks* and keep track of them.

I built this as a personal tool for keeping track of future events, I tend to forget everything so yeah, this has been really useful for me.


## Installation

You can either install it via `go install` or just download the binary file

This will install the latest version

```bash
go install github.com/zSnails/taskr@latest
```
    
## Usage/Examples

```bash
taskr new <date> <description>
# this will create a new task
# for example
taskr new "2022-04-16 7:0:0" "Go to my grandma's birthday"

taskr delete <id>
# this will delete a task with the id <id>
# for example
taskr delete 5

taskr all
# this will show all tasks, taskr by default only shows future tasks
```

