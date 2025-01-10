package kata


func CloseCompare(a, b, margin float64) int {
  diff := a - b
	if diff < 0 {
		diff = -diff
	}

	if diff <= margin {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}
