package kata

func Summation(n int) int {
    // the sleeper must awaken!
  count := 0
  for i := 1; i <= n; i++ {
    count += i
  }
  return count
}
