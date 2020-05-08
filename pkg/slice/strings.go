package slice

func StringContains(s string, slice []string) bool {
	for _, s2 := range slice {
		if s == s2 {
			return true
		}
	}
	return false
}
