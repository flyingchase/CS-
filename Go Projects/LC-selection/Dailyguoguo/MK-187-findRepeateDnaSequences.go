package Dailyguoguo

// 使用 set 存储出现过的字符串，遍历保存
func findRepeatedDnaSequences(s string) []string {
	res := make([]string, 0)
	if len(s) == 0 {
		return res
	}
	set := make(map[string]int, 0)
	size := len(s)
	for i := 0; i+10 <= size; i++ {
		sequence := s[i : i+10]
		set[sequence] += 1
		count := set[sequence]
		if count == 2 {
			res = append(res, sequence)
		}
	}
	return res
}

// Rabin-Karp 算法
func findRepeatedDnaSequences2(s string) []string {
	if len(s) < 10 {
		return nil
	}
	res, dict := []string{}, make(map[int64]struct{})
	var base, prime, head, hash int64
	base, prime, head = 4, 1000000007, 1
	dna := map[byte]int64{'A': 0, 'C': 1, 'G': 2, 'T': 3}
	for i := 0; i < 9; i++ {
		head = (head * base) % prime
	}
	hash = 0
	for i := 0; i < 10; i++ {
		hash = (hash*base + dna[s[i]]) % prime
	}
	for i := 0; i <= len(s)-10; i++ {
		_, ok := dict[hash]
		if ok && !isDup(res, s[i:i+10]) {
			res = append(res, s[i:i+10])
		} else if !ok {
			dict[hash] = struct{}{}
		}
		if i+10 < len(s) {
			hash = (base * (hash + prime - dna[s[i]]*head) + dna[s[i+10]]) % prime
		}
	}
	return res
}
func isDup(res []string, s string) bool {
	for i := 0; i < len(res); i++ {
		if res[i] == s {
			return true
		}
	}
	return false
}
