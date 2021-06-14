# Git Audit

Git Audit tool computing changelogs from public GitHub repositories.
This tool can be either run as a client, calling the GitHub API directly,
or as a gRPC server.

> Note: This tool uses public GitHub API which is rate limited to 60 requests per hour.

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

#### Examples

Display the latest commit from the
[octocat/Hello-World](https://github.com/octocat/Hello-World) repository:

```bash
$ git-audit client --owner octocat --repository Hello-World
```

Output:

```text
changelog for github.com/octocat/Hello-World from HEAD
7fd1a60b01f91b314f59955a4e4d4e80d8edf11d Merge pull request #6 from Spaceghost/patch-1
```

This can be equivalently called by appending `--sha HEAD` flag to the previous command:

```bash
$ git-audit client --owner octocat --repository Hello-World --sha HEAD
```

Output:

```text
changelog for github.com/octocat/Hello-World from HEAD
7fd1a60b01f91b314f59955a4e4d4e80d8edf11d Merge pull request #6 from Spaceghost/patch-1
```

Show the full commit message of the latest commit:

```bash
$ git-audit client --owner octocat --repository Hello-World --verbose
```

Output:

```text
changelog for github.com/octocat/Hello-World from HEAD
7fd1a60b01f91b314f59955a4e4d4e80d8edf11d Merge pull request #6 from Spaceghost/patch-1
	New line at end of file.
```

To get the complete changelog of the repository,
supply the commit hash of the first commit.
Initial commit of this repository has hash `553c2077f0edc3d5dc5d17262f6aa498e69d6f8e`.

```bash
$ git-audit client --owner octocat --repository Hello-World --sha 553c2077f0edc3d5dc5d17262f6aa498e69d6f8e
```

Output:

```text
changelog for github.com/octocat/Hello-World from 553c2077f0edc3d5dc5d17262f6aa498e69d6f8e
762941318ee16e59dabbacb1b4049eec22f0d303 New line at end of file. --Signed off by Spaceghost
7fd1a60b01f91b314f59955a4e4d4e80d8edf11d Merge pull request #6 from Spaceghost/patch-1
553c2077f0edc3d5dc5d17262f6aa498e69d6f8e first commit
```

### Running the gRPC Server

```bash
git-audit server --port <PORT>
```

#### Examples

Run the gRPC server on port 8080:

```bash
git-audit server --port 8080
```

Example of calling this gRPC server from a client can be found at
[`/examples/grpc_client`](https://github.com/jsfpdn/git-audit/blob/main/examples/grpc_client/main.go).
To get the latest commit merged to the default branch in the official Go repository,
run the following command.

```bash
$ go run examples/grpc_client/main.go localhost:8080 golang go HEAD
```

The output should look something like this:

```text
owner:"golang" repo:"go" SHA:"HEAD"
326ea438bb579a2010e38e00f515a04344ff96b0 cmd/compile: rewrite a, b = f() to use temporaries when type not identical
	If any of the LHS expressions of an OAS2FUNC are not identical to the
	respective function call results, escape analysis mishandles the
	implicit conversion, causes memory corruption.

	Instead, we should insert autotmps like we already do for f(g()) calls
	and return g() statements.

	Fixes #46725

	Change-Id: I71a08da0bf1a03d09a023da5b6f78fb37a4a4690
	Reviewed-on: https://go-review.googlesource.com/c/go/+/327651
	Trust: Cuong Manh Le <cuong.manhle.vn@gmail.com>
	Run-TryBot: Cuong Manh Le <cuong.manhle.vn@gmail.com>
	TryBot-Result: Go Bot <gobot@golang.org>
	Reviewed-by: Matthew Dempsky <mdempsky@google.com>
```

#### Running the gRPC Server in a Container

gRPC server can be ran in a Docker container.
Run the following commands to build a docker image
and run a docker container with the gRPC server listening on port 8080.

```bash
make docker
docker run --rm -p 8080:8080 josefpodanyml/git-audit server --port 8080
```

Try running the gRPC client from above.

```bash
$ go run examples/grpc_client/main.go localhost:8080 golang go HEAD
```

The output should look something like this:

```text
owner:"golang" repo:"go" SHA:"HEAD"
326ea438bb579a2010e38e00f515a04344ff96b0 cmd/compile: rewrite a, b = f() to use temporaries when type not identical
	If any of the LHS expressions of an OAS2FUNC are not identical to the
	respective function call results, escape analysis mishandles the
	implicit conversion, causes memory corruption.

	Instead, we should insert autotmps like we already do for f(g()) calls
	and return g() statements.

	Fixes #46725

	Change-Id: I71a08da0bf1a03d09a023da5b6f78fb37a4a4690
	Reviewed-on: https://go-review.googlesource.com/c/go/+/327651
	Trust: Cuong Manh Le <cuong.manhle.vn@gmail.com>
	Run-TryBot: Cuong Manh Le <cuong.manhle.vn@gmail.com>
	TryBot-Result: Go Bot <gobot@golang.org>
	Reviewed-by: Matthew Dempsky <mdempsky@google.com>
```

## Deploying the gRPC Server

```bash
k apply -f ./deployment/deployment.yaml
k apply -f ./deployment/service.yaml
```
