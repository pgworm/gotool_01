// Package gotool is a library of utility functions used to implement the standard "Go" tool provided
// as a convenience to developers who want to write tools with similar semantics.
package gotool

import "go/build"

// Export functions here to make it easier to keep the implementations up to date with upstream.

var DefaultContext = Context{
	BuildContext: build.Default,
}

type Context struct {
	BuildContext build.Context
}

// ImportPaths returns the import paths to use for the given command line.
//
// The path "all" is expanded to all packages in $GOPATH and $GOROOT.
// The path "std" is expanded to all packages in the Go standard library.
// The path "cmd" is expanded to all Go standard commands.
// The string "..." is treated as a wildcard within a path.
// Relative import paths are not converted to full import paths.
// If args is empty, a single element "." is returned.
func (c *Context) ImportPaths(args []string) []string {
	return c.importPaths(args)
}

// ImportPaths returns the import paths to use for the given command line
// using default context.
//
// The path "all" is expanded to all packages in $GOPATH and $GOROOT.
// The path "std" is expanded to all packages in the Go standard library.
// The path "cmd" is expanded to all Go standard commands.
// The string "..." is treated as a wildcard within a path.
// Relative import paths are not converted to full import paths.
// If args is empty, a single element "." is returned.
func ImportPaths(args []string) []string {
	return DefaultContext.importPaths(args)
}
