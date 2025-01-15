package kata

import "strings"

func ReverseWords(str string) string {
  list := strings.Split(str, " ")
  for i, j := 0, len(list) - 1; i < j; i, j = i + 1, j - 1 {
    list[i], list[j] = list[j], list[i]
  }
  return strings.Join(list, " ") // reverse those words
}
