// Package version records versioning information about this module.
package version

var Version string

// String formats the semantic version string for this module.
func String() string {
	return Version
}
