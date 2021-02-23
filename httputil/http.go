package httputil

import "regexp"

func CheckMobileAgent(agent string) (isMobile bool) {

	reg := regexp.MustCompile(`(?i:(mozilla|windows nt))`)

	if len(reg.FindAllString(agent, -1)) == 0 {
		isMobile = true
	}
	return
}
