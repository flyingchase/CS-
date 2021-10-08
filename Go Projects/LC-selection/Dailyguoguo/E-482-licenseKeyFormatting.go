package Dailyguoguo

import (
	"strings"
)

func licenseKeyFormatting(s string, k int) string {
	s = strings.Replace(s, "-", "", -1)
	n, counter, res := len(s), 0, ""
	for i := n - 1; i >= 0; i-- {
		res = string(s[i]) + res
		counter++
		if counter%k == 0 && i != 0 {
			res = "-" + res
		}
	}
	s = strings.ToUpper(res)
	return s
}
func liceseKeyFormatting2(s string, k int) string {
	s = strings.ReplaceAll(s, "-", "")
	n, res := len(s), []string{}
	if n == 1 {
		return strings.ToUpper(s)
	}
	left := n % k
	if left != 0 {
		res = append(res, s[0:left])
	}
	for i := left; i < n; i += k {
		res = append(res, s[i:i+k])
	}
	ans := strings.Join(res, "-")
	return ans
}
