package kata

func Feast(beast string, dish string) bool {
  // your code here
  blength := len(beast)
  dlength := len(dish)
  
  if beast[0] == dish[0] && beast[blength - 1] == dish[dlength - 1] {
    return true
  }
  return false
}
