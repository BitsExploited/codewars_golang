package kata

func Between(a, b int) []int {
      // your code here
  var arr []int
  for i := a; i <= b; i++ {
    arr = append(arr, i)
  }
  return arr
}
