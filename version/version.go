package version

//go:generate go run ./update/.

import (
	"fmt"
)

var (
	// URL is the git URL for the repository
	URL = "github.com/cybriq/log"
	// GitRef is the gitref, as in refs/heads/branchname
	GitRef = "refs/heads/main"
	// GitCommit is the commit hash of the current HEAD
	GitCommit = "16bacb218695186e0fa83e20bbbce3bb0e89e619"
	// BuildTime stores the time when the current binary was built
	BuildTime = "2021-06-12T09:27:55+03:00"
	// Tag lists the Tag on the build, adding a + to the newest Tag if the commit is
	// not that commit
	Tag = "v0.0.5"
	// PathBase is the path base returned from runtime caller
	PathBase = "/home/loki/src/github.com/cybriq/log/"
	// Major is the major number from the tag
	Major = 0
	// Minor is the minor number from the tag
	Minor = 0
	// Patch is the patch version number from the tag
	Patch = 5
	// Meta is the extra arbitrary string field from Semver spec
	Meta = ""
)

// Get returns a pretty printed version information string
func Get() string {
	return fmt.Sprint(
		"\nRepository Information\n",
		"\tGit repository: "+URL+"\n",
		"\tBranch: "+GitRef+"\n",
		"\tCommit: "+GitCommit+"\n",
		"\tBuilt: "+BuildTime+"\n",
		"\tTag: "+Tag+"\n",
		"\tMajor:", Major, "\n",
		"\tMinor:", Minor, "\n",
		"\tPatch:", Patch, "\n",
		"\tMeta: ", Meta, "\n",
	)
}
