package utils

import "regexp"

func StringsFromRegex(in string, r string) [][]string {

	regex := regexp.MustCompile(r)

	lines := Lines(in)
	ret := make([][]string, len(lines))
	for i, line := range lines {
		ret[i] = regex.FindStringSubmatch(line)
	}

	return ret
}
