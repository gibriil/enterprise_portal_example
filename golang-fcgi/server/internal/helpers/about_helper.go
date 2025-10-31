package helpers

import (
	"fmt"
	"regexp"
	"runtime/debug"
	"time"

	"github.com/gibriil/enterprise_portal_example/internal"
)

type About struct {
	Version       string
	CopyrightYear int
	CanConnect    bool `json:"ValidModernCampusConnection"`
}

type VersionControlSystem struct {
	Type            string
	Revision        string
	RevisionTime    string
	AheadOfRevision bool
}

func GetVersionNumber() string {
	build, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}

	regex := regexp.MustCompile(`(^v\d+\.\d+\.\d+(?:\-rc\d+)?)`)

	version := regex.FindString(build.Main.Version)

	vcs := GetVersionControlInformation()

	if vcs.AheadOfRevision {
		version = fmt.Sprintf("%s:%s", version, vcs.Revision)
	}

	return version
}

func GetVersionControlInformation() *VersionControlSystem {
	build, ok := debug.ReadBuildInfo()

	if !ok {
		panic("No Build Information Available")
	}

	var VCS VersionControlSystem

	for _, setting := range build.Settings {
		switch setting.Key {
		case "vcs":
			VCS.Type = setting.Value
		case "vcs.revision":
			VCS.Revision = setting.Value
		case "vcs.time":
			VCS.RevisionTime = setting.Value
		case "vcs.modified":
			if setting.Value == "true" {
				VCS.AheadOfRevision = true
			} else {
				VCS.AheadOfRevision = false
			}
		}
	}

	return &VCS
}

func GetAboutInformation(app *internal.Application) *About {
	year, _, _ := time.Now().Date()

	return &About{
		Version:       GetVersionNumber(),
		CopyrightYear: year,
	}
}
