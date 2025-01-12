package kata

func CountPositivesSumNegatives(numbers []int) []int {
  count := 0
  add := 0
  for _, i := range numbers {
    if i > 0 {
      count++ // adding teh positive numbers
    } else {
      add += i // adding the negative numbers
    }
  }
  return []int{count, add} // returning as per the question
}
