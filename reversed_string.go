package kata

func Solution(word string) string {
  str := ""
  length := len(word)
  runes := []rune(word)
  
  for i := length - 1; i >= 0; i-- {
    str += string(runes[i])
  }
  
  return str
}
