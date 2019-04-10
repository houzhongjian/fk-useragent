package os

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

//GetVersion .
func GetVersion(ua string) string {
	osName := GetName(ua)
	if osName == "other" {
		return "other"
	}

	if osName == "nokia" {
		//单独处理nokia.
	}

	reg := fmt.Sprintf("%s [\\d.]+", osName)
	r, err := regexp.Compile(reg)
	if err != nil {
		log.Printf("err:%+v\n", err)
		return ""
	}
	s := r.FindString(ua)

	s = strings.Replace(s, osName, "", -1)
	return s
}
