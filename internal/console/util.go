package console

import "strings"

func containsStringSlice(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getStringInBetween(value string, a string, b string) string {
	firstSplit := strings.Split(value, a)
	if len(firstSplit) > 1 {
		secondSplit := strings.Split(firstSplit[1], b)
		if len(secondSplit) > 0 {
			return strings.TrimSpace(secondSplit[0])
		}
	}

	return ""
}
