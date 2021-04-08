package glc

func add(n int) int {
	if n < 0 {
		return 0
	} else {
		return n
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
