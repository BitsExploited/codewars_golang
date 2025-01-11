package kata

func Digitize(n int) []int {
  if n == 0{
    return []int{0}
  }
  var result []int
  
  for n > 0 {
    digits := n % 10
    result = append(result, digits)
    n /= 10
  }
  return result
}
