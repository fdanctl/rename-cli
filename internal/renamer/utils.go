package renamer

func FormatDigitsStr(s string, d int) string {
	if len(s) >= d {
		return s
	}
	return FormatDigitsStr("0"+s, d)
}
