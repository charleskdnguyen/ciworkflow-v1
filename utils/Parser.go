package utils

func parseString(v, dv string) string {
	if v != "" {
		return v
	}

	return dv
}
