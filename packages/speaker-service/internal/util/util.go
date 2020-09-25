package util

func SetIfNotEmpty(m map[string]string, key, val string) {
	if val != "" {
		m[key] = val
	}
}