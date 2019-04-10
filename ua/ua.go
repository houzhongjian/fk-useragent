package ua

import (
	"strings"

	"uaparse.com/uaparse/ua/os"
)

//UserAgentResult .
type UserAgentResult struct {
	Status    string `json:"status"`
	Os        string `json:"os"`
	OsVersion string `json:"os_version"`
}

//Parse .
func Parse(ua string) UserAgentResult {
	ua = strings.Replace(strings.ToLower(ua), "_", ".", -1)

	return parse(ua)
}

func parse(ua string) UserAgentResult {
	res := UserAgentResult{
		Os:        os.GetName(ua),
		Status:    "success",
		OsVersion: os.GetVersion(ua),
	}
	return res
}

func isMobile(ua string) bool {
	return strings.Contains(ua, "mobile")
}
