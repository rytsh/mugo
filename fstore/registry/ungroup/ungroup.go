package ungroup

func Ungroup() map[string]interface{} {
	return map[string]interface{}{
		"nothing": Nothing,
	}
}

func Nothing(v ...any) string {
	return ""
}
