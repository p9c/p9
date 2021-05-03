package version

//go:generate go run ./update/.

import (
	"fmt"
)

var (

	// URL is the git URL for the repository
	URL = "github.com/p9c/p9"
	// GitRef is the gitref, as in refs/heads/branchname
	GitRef = "refs/heads/main"
	// GitCommit is the commit hash of the current HEAD
	GitCommit = "40a7d8327e70dca9576a7257ad320106d62ee72f"
	// BuildTime stores the time when the current binary was built
	BuildTime = "2021-05-03T14:08:30+02:00"
	// Tag lists the Tag on the build, adding a + to the newest Tag if the commit is
	// not that commit
	Tag = "v0.0.2+"
	// PathBase is the path base returned from runtime caller
	PathBase = "/home/loki/src/github.com/p9c/p9/"
	// Major is the major number from the tag
	Major = 0
	// Minor is the minor number from the tag
	Minor = 0
	// Patch is the patch version number from the tag
	Patch = 2
	// Meta is the extra arbitrary string field from Semver spec
	Meta = ""
)

// Get returns a pretty printed version information string
func Get() string {
	return fmt.Sprint(
		"\nRepository Information\n"+
		"\tGit repository: "+URL+"\n",
		"\tBranch: "+GitRef+"\n"+
		"\tCommit: "+GitCommit+"\n"+
		"\tBuilt: "+BuildTime+"\n"+
		"\tTag: "+Tag+"\n",
		"\tMajor:", Major, "\n",
		"\tMinor:", Minor, "\n",
		"\tPatch:", Patch, "\n",
		"\tMeta: ", Meta, "\n",
	)
}
