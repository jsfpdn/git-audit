# git-audit-tool

Git Audit tool computing changelogs from public GitHub repositories. 

## Install

```bash
go get github.com/jsfpdn/git-audit
```

## Usage

Show help with `git-audit help`.
To see help for individual subcommands,
run `git-audit help client` or `git-audit help server`.

### Using CLI Client

```bash
git-audit client --owner <OWNER> --repository <REPOSITORY>
```

By default, commits are pretty-printed on a single line,
like when running `git log --pretty=oneline`.
To see complete commit messages, the `--verbose` (`-v`) flag.

### Running the gRPC Server

```bash
git-audit server --port <PORT>
```


## TODOs

- [] decrease Docker image size