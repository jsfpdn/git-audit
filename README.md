# git-audit-tool

Git Audit tool computing changelogs from public GitHub repositories.
This tool can be either run as a client, calling the GitHub API directly,
or as a gRPC server.

## Install

```bash
git clone https://github.com/jsfpdn/git-audit.git
cd git-audit
make install
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
similar to `git log --pretty=oneline`.
To see whole commit messages, use the `--verbose` (`-v`) flag.

### Running the gRPC Server

```bash
git-audit server --port <PORT>
```

Example of calling this gRPC server from a client can be found at `/examples/grpc_client`.
Try running `go run examples/grpc_client/main.go <server_address> golang go HEAD`
to get the latest commit merged to the default branch in the official Go repository.

#### Running the gRPC Server in a Container

```bash
make docker
docker run --rm -p 8080:8080 josefpodanyml/git-audit server --port 8080
```

## Deploying the gRPC Server

```bash
k apply -f ./deployment/deployment.yaml
k apply -f ./deployment/service.yaml
```
