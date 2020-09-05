package basics

import "testing"

// go build -ldflags="-X main.Version 2.0.0" program_name
// similar to cmake's set_compiler_definition() technique
// in the real world, this version string can come from
// the build pipeline and git
var Version = "1.0.0"

func TestSomething(t *testing.T) {

}
