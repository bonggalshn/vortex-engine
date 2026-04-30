package util

func ListContains(stringList []string, item string) bool {
	for _, element := range stringList {
		if element == item {
			return true
		}
	}
	return false
}
