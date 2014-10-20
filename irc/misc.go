package irc

import (
	"strings"
)

func merge_modes(curmode string, newmode string) string {
	ret := curmode
	// Start in "add" mode
	mode_add := true
	for _, ch := range []byte(newmode) {
		if ch == '+' {
			mode_add = true
		} else if ch == '-' {
			mode_add = false
		} else {
			if mode_add {
				// Add flags
				if strings.IndexByte(ret, ch) < 0 {
					ret += string(ch)
				}
			} else {
				// Remove flags
				if strings.IndexByte(ret, ch) >= 0 {
					ret = strings.Replace(ret, string(ch), "", -1)
				}
			}
		}
	}
	// Add leading + if needed
	if !strings.HasPrefix(ret, "+") {
		ret = "+" + ret
	}
	return ret
}
