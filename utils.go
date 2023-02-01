package dockerfileparser

import "strings"

func parseLine(s string) (k, v string) {
	ss := strings.Split(s, " ")
	k = ss[0]
	v = strings.Join(ss[1:], " ")
	return
}
