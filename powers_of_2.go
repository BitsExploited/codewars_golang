package kata

func PowersOfTwo(n int) []uint64 {
  result := make([]uint64, n + 1)
  
  for i := 0; i <= n; i++ {
    result[i] = 1 << i
  }
  return result
}
