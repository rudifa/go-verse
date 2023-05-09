# Verse ...

## Create a go + cobra project

- check preconditiond: go and cobra-cli installed

```
verse % go version
go version go1.20.3 darwin/amd64

verse % cobra-cli version
Error: unknown command "version" for "cobra-cli"
```

- create the project skeleton

```
golang % mkdir verse && cd verse

verse % go mod init verse

verse % cobra-cli init
Your Cobra application is ready at
/Users/rudifarkas/GitHub/golang/verse
```

- look

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
A longer description that spans multiple lines and likely contains
...
```

- create the git repo

```
verse % echo '# see also https://github.com/github/gitignore/blob/main/community/Golang/Go.AllowList.gitignore' > .gitignore
verse % echo '.DS_Store' >> .gitignore

verse % git init
Initialized empty Git repository in /Users/rudifarkas/GitHub/golang/verse/.git/

```

```

```

```

```

```

```
