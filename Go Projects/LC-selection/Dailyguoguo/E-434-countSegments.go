package Dailyguoguo

func countSegments(s string) int {
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			continue
		}
		count++
	}
	return count
}
