[![License][license-badge]][license-link]
[![Go Reference][godoc-badge]][godoc-link]
[![Go Report Card][goreportcard-badge]][goreportcard-link]
[![Actions][github-actions-badge]][github-actions-link]
[![Releases][github-release-badge]][github-release-link]

# Go Build Metadata

✨ Common application metadata for use with go build and ldflags

## Motivations

The Go toolchain has a feature where you can specify the value for a named
variable at `go build`-time on the command line, and that value will be embedded
into that variable in the resulting binary.

This feature is primarily used for embedding a version string or git SHA, so
that the provenance of the resulting binary can be traced back to the exact ref
from which it was produced.

See [this blog post from Cloudflare](https://blog.cloudflare.com/setting-go-variables-at-compile-time)
for an introduction.

This library provides three things:

1. A number of appropriately named variables that can accept values when
   running `go build`.
2. The automatic conversion of string values into complex types (like
   `url.URL`) where appropriate.
3. Public functions for retrieving these typed values.

This library also has the capacity to accept a wide variety of relevant build
values, and uses a short import path so that they can be succinctly configured
as part of a `go build` command.

## Installing

This library can be installed by running:

```shell
go get -u jdk.sh/meta
```

## Usage

### Example

Create a `main.go` file containing the following:

```go
package main

import (
    "fmt"
    "jdk.sh/meta"
)

func main() {
    fmt.Println(
        "version", meta.Version(),
        "built on", meta.Date(),
    )
}
```

Build your `main.go` into a binary using `go build`. Notice the usage of
`-ldflags`, specifying values for the variables `jdk.sh/meta.version` and
`jdk.sh/meta.date`:

```shell
go build -ldflags "\
    -X 'jdk.sh/meta.version=v1.2.3' \
    -X 'jdk.sh/meta.date=$(date -R)' \
  " main.go
```

Finally, execute the program and observe that the given values are now "baked"
into the resulting binary:

```shell
./main
version v1.2.3 built on 2019-08-23 18:00:00 +0000 UTC
```

### Variables

| Name                      | Purpose                                                                                                                                                                                        |
| ------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `jdk.sh/meta.author`      | The name of the application author. May contain their name, email address, or optionally both.                                                                                                 |
| `jdk.sh/meta.author_url`  | URL for the application author. Typically links to the author's personal homepage or Github profile.                                                                                           |
| `jdk.sh/meta.copyright`   | The copyright for the application. Typically the name if the author or organization, sometimes prefixed with a year or year range.                                                             |
| `jdk.sh/meta.date`        | The time that the application was built. Supports several common formats.                                                                                                                      |
| `jdk.sh/meta.desc`        | Description for the application. Typically a longer statement describing what the application does.                                                                                            |
| `jdk.sh/meta.dev`         | The development status for the application. An application in development mode may indicate that it's using experimental or untested features, and should be used with caution.                |
| `jdk.sh/meta.docs`        | URL for application documentation. Typically links to a page where a user can find technical documentation.                                                                                    |
| `jdk.sh/meta.license`     | The license identifier for the application. Should not the full license body, but one of the identifiers from https://spdx.org/licenses, so that the type of license can be easily determined. |
| `jdk.sh/meta.license_url` | URL for the application license. Typically links to a page where the verbatim license body is available.                                                                                       |
| `jdk.sh/meta.name`        | The name of the application. Typically named the same as the binary, or for display in an error or help message.                                                                               |
| `jdk.sh/meta.note`        | An arbitrary message for the application. Can be used to store a message about the build environment, release, etc.                                                                            |
| `jdk.sh/meta.sha`         | Git SHA that was used to build the application. A 40 character "long" SHA should be provided.                                                                                                  |
| `jdk.sh/meta.src`         | URL for the application source code. Typically links to a repository where a user can browse or clone the source code.                                                                         |
| `jdk.sh/meta.title`       | The title of the application. Typically a full or non-abbreviated form of the application name.                                                                                                |
| `jdk.sh/meta.url`         | URL for the application homepage. Typically links to a page where a user can learn more about the application.                                                                                 |
| `jdk.sh/meta.version`     | The version slug for the application. The value can be used to point back to a specific tag or release. Supports semver, see https://semver.org.                                               |

## License

This code is distributed under the [MIT License][license-link], see [LICENSE.txt][license-file] for more information.

[github-actions-badge]:  https://github.com/joshdk/meta/workflows/Build/badge.svg
[github-actions-link]:   https://github.com/joshdk/meta/actions
[github-release-badge]:  https://img.shields.io/github/release/joshdk/meta/all.svg
[github-release-link]:   https://github.com/joshdk/meta/releases
[godoc-badge]:           https://pkg.go.dev/badge/jdk.sh/meta.svg
[godoc-link]:            https://pkg.go.dev/jdk.sh/meta
[goreportcard-badge]:    https://goreportcard.com/badge/jdk.sh/meta
[goreportcard-link]:     https://goreportcard.com/report/jdk.sh/meta
[license-badge]:         https://img.shields.io/badge/license-MIT-green.svg
[license-file]:          https://github.com/joshdk/meta/blob/master/LICENSE.txt
[license-link]:          https://opensource.org/licenses/MIT
