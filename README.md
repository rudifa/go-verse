# Verse ... the genessis of a cli application

## Create a go + cobra project

- preconditions: are go and cobra-cli installed?

```
verse % go version
go version go1.20.3 darwin/amd64

verse % cobra-cli
Cobra is a CLI library for Go that empowers applications.
...
```

- recommended: vscode, with extension `Go for Visual Studio Code` installed

- create the project skeleton

```
golang % mkdir verse && cd verse

verse % go mod init verse

verse % cobra-cli init
Your Cobra application is ready at
/Users/rudifarkas/GitHub/golang/verse
```

- look at what we have

```
verse % tree
.
├── cmd
│   └── root.go
├── LICENSE
├── go.mod
├── go.sum
└── main.go

2 directories, 5 files
```

- run (builds an executable in memory and runs it)

```
verse % go run .
A longer description that spans multiple lines and ...
```

- build an executable and run it

```
verse % go build .

verse % ./verse
A longer description that spans multiple lines and ...
```

- create the git repo

```
verse % echo '# see also https://github.com/github/gitignore/blob/main/community/Golang/Go.AllowList.gitignore' > .gitignore
verse % echo '.DS_Store' >> .gitignore

verse % git init
Initialized empty Git repository in /Users/rudifarkas/GitHub/golang/verse/.git/

verse % git add .
verse % git commit -m 'initial commit'
[main (root-commit) d5e8577] initial commit
...
```

- add the first command

```
verse % cobra-cli add future
future created at /Users/rudifarkas/GitHub/golang/verse

verse % go run .
A longer description that spans multiple lines and ...
...
Usage:
  verse [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  future      A brief description of your command
  help        Help about any command

Flags:
  -h, --help     help for verse
  -t, --toggle   Help message for toggle

Use "verse [command] --help" for more information about a command.

verse % go run . future
future called
```

- look again

```
verse % tree
.
├── cmd
│   ├── future.go
│   └── root.go
├── LICENSE
├── README.md
├── go.mod
├── go.sum
├── main.go
└── verse
```

- in `cmd/root.go` add the 2 lines (optional \*)

```
func Execute() {
    cobra.EnablePrefixMatching = true                  // allow abbreviations
    rootCmd.CompletionOptions.DisableDefaultCmd = true // disable default completion command
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}
```

(\*) my preferences: with `allows...`, it suffices to type just the first few letters of a command name; with `disables...` the app will not offer to generate shell scripts for installing command completions in the shell environment`

- now edit the help texts in `root.go`and `future.go` to something appropriate, then run

```
verse % go run . help
This is a go application that uses cobra to add CLI functionality.
...
Available Commands:
  future      A look into the future
  help        Help about any command

verse % go run . help fut
Future is now, go for it.
...
```

- commit

```
git add .
verse % git commit -m 'add command future and update help messages'
```

- now add a package to which to delegate the actual work for the command `future`

```
verse % mkdir pkg
verse % mkdir pkg/future
verse % touch pkg/future/peek.go
```

- add code to `pkg/future/peek.go`

```
package future

import (
    "os/exec"
)

func OpenOpenAIPage() error {
    cmd := exec.Command("open", "https://openai.com")
    return cmd.Run()
}
```

- in `cmd/future.go` call `future.OpenOpenAIPage()`

```
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("future called")
        future.OpenOpenAIPage()
    },
```

When you save above change, the vscode go extension inserts for you

```
import (
    "fmt"
    "verse/pkg/future"

    "github.com/spf13/cobra"
)
```

- run it and enjoy the view

```
verse % go run . fut
```

- look yet again at the project structure

```
verse % tree
.
├── cmd
│   ├── future.go
│   └── root.go
├── pkg
│   └── future
│       └── peek.go
├── LICENSE
├── README.md
├── go.mod
├── go.sum
├── main.go
└── verse
```

Normally you don't edit files `go.mod` `go.sum`, managed by the go tools. Their contents reflect the packages imported into the project, directly or as depenedencies.

- commit, create a github repo and push, update LICENSE, ...

There we go.
