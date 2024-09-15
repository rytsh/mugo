package external

func Helper() map[string]interface{} {
	return map[string]interface{}{
		"nothing": nothing,
	}
}

func nothing(v ...any) string {
	return ""
}
