package kata

func SquareSum(numbers []int) int {
  sum := 0
  for _, j := range numbers {
    sum += j * j
  }
    return sum
}
