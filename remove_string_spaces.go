package kata

func NoSpace(word string) string {
  str := ""
  for i := 0; i < len(word); i++ {
    if word[i] != ' ' {
      str += string(word[i])
      }
    }
  return str
  }
