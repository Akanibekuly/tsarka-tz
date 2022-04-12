package utils

func FindSubstr(s string) string {
	slow, fast := 0, 0
	max, start, end := -1, 0, 0
	for fast < len(s) && slow < len(s) {
		dict := make(map[byte]int)
		fast = slow
		for fast < len(s) {
			if _, ok := dict[s[fast]]; ok {
				break
			}
			dict[s[fast]] = fast
			fast++
		}

		if fast == len(s) {
			if fast-slow > max {
				start, end = slow, fast
			}
			break
		}

		if fast-slow > max {
			max = fast - slow
			start, end = slow, fast
		}

		slow = dict[s[fast]] + 1
	}

	return s[start:end]
}
