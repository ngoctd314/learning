package leetcode

func isPalindrome(s string) bool {
	conv := func(r byte) (byte, bool) {
		if r >= 65 && r <= 90 {
			r += 32
			return r, true
		}
		if r >= 97 && r <= 122 {
			return r, true
		}
		if r >= 48 && r <= 57 {
			return r, true
		}

		return 0, false
	}
	i, j := 0, len(s)-1
	for i < j {
		ri, ok1 := conv(s[i])
		if !ok1 {
			i++
			continue
		}

		rj, ok2 := conv(s[j])
		if !ok2 {
			j--
			continue
		}

		if ri != rj {
			return false
		}
		i++
		j--
	}

	return true
}
