package misc

import "strings"

func GetJsonName(tag string, fallback string) string {
	switch tag {
	case "-":
	case "":
		return fallback
	default:
		idx := 0
		if idx = strings.Index(tag, ","); idx < 0 {
			idx = len(tag)
		} else if idx == 0 {
			return fallback
		}
		return tag[:idx]
	}
	return ""
}
