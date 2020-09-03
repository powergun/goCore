# go core constructs

## devops notes

build a component (package)

`go build '<package_name>'`

(each source file defines the package it belongs to)

run unit tests in a package:

`cd core; go test github.com/powergun/goCore/typings`

note I have to compose a path from pwd to the location of the package
`typings`

## structure of a project

### multiple binaries per project

source: <https://ieftimov.com/post/golang-package-multiple-binaries/>

example: goNetwork/http_/foo_server

to install the foo_server binary:

`go install github.com/powergun/goNetwork/http_/foo_server`

this program is free to import any functionality from `github.com/`
namespace