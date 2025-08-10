package renamer

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Returns trailingPath, filename and file extention
func ParseFileName(fullPath string) (trailingPath, filename, extension string) {
	re := regexp.MustCompile(`(^.*\/|^)([^\/]+?)(\.[^\.]+)?$`)
	groups := re.FindStringSubmatch(fullPath)
	trailingPath = groups[1]
	filename = groups[2]
	extension = groups[3]
	return
}

// Replaces %n, %0n, ... with number
func Enumerate(filename string, n int) string {
	// enumeration
	re := regexp.MustCompile(`%(0*)n`)
	replaced := re.ReplaceAllStringFunc(filename, func(s string) string {
		groups := re.FindStringSubmatch(s)
		digits := len(groups[1]) + 1

		res := FormatDigitsStr(strconv.Itoa(n), digits)

		return res
	})
	return replaced
}

func InsertDate(filename string) string {
	re := regexp.MustCompile(`%(Y{2,4}|M{2}|D{2}|h{2}|m{2}|s{2})`)
	replaced := re.ReplaceAllStringFunc(filename, func(s string) string {
		group := re.FindStringSubmatch(s)
		now := time.Now()
		strings.Join([]string{"D", "d"}, "")

		switch group[1] {
		case "YY":
			return string([]rune(strconv.Itoa(now.Year()))[2:])
		case "YYYY":
			return strconv.Itoa(now.Year())
		case "MM":
			return FormatDigitsStr(strconv.Itoa(int(now.Month())), 2)
		case "DD":
			return FormatDigitsStr(strconv.Itoa(now.Day()), 2)
		case "hh":
			return FormatDigitsStr(strconv.Itoa(now.Hour()), 2)
		case "mm":
			return FormatDigitsStr(strconv.Itoa(now.Minute()), 2)
		case "ss":
			return FormatDigitsStr(strconv.Itoa(now.Second()), 2)
		default:
			return s
		}
	})
	return replaced
}
