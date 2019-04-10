package os

import "strings"

//GetName .
func GetName(ua string) string {
	var os string
	if isAndroid(ua) {
		os = "android"
	} else if isIphone(ua) {
		os = "iphone os"
	} else if isMac(ua) {
		os = "mac os x"
	} else if isIpad(ua) {
		os = "ipad"
	} else if isWindows(ua) {
		os = "windows nt"
	} else if isLinux(ua) {
		os = "linux"
	} else if isNokia(ua) {
		os = "nokia"
	} else {
		os = "other"
	}
	return os
}

//是否为安卓.
func isAndroid(ua string) bool {
	return strings.Contains(ua, "android")
}

//是否为iphone.
func isIphone(ua string) bool {
	return strings.Contains(ua, "iphone os")
}

//是否为mac.
func isMac(ua string) bool {
	return strings.Contains(ua, "mac os x")
}

//是否为windows.
func isWindows(ua string) bool {
	return strings.Contains(ua, "windows nt")
}

//是否为ipad.
func isIpad(ua string) bool {
	return strings.Contains(ua, "ipad")
}

//是否为linux.
func isLinux(ua string) bool {
	isLinux := strings.Contains(ua, "linux")
	if isLinux && !isAndroid(ua) {
		return true
	}
	return false
}

func isNokia(ua string) bool {
	return strings.Contains(ua, "nokia")
}
