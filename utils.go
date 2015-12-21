package advent

func Permutations(s []string) [][]string {
	var output [][]string

	var heapPermute func(s []string, n int)

	heapPermute = func(s []string, n int) {
		if n == 1 {
			tmp := make([]string, len(s))
			copy(tmp, s)
			output = append(output, tmp)
		} else {
			for i := 0; i < n; i++ {
				heapPermute(s, n-1)
				if n%2 == 1 {
					s[0], s[n-1] = s[n-1], s[0]
				} else {
					s[i], s[n-1] = s[n-1], s[i]
				}
			}
		}
	}
	heapPermute(s, len(s))
	return output
}
