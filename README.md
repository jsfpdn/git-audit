# git-audit-tool

Git Audit tool computing changelogs from public GitHub repositories. 

## Install

```bash
go get github.com/jsfpdn/git-audit
```

## Usage

### Using CLI Client

```bash
git-audit client --owner <OWNER> --repository <REPOSITORY>
```

### Running the gRPC Server

```bash
git-audit server --port <PORT>
```


## TODOs

- [] decrease Docker image size