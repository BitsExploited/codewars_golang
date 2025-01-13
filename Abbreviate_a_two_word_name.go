package kata

import "strings"

func AbbrevName(name string) string{
  //your code here
  split := strings.Split(name, " ")
  return strings.ToUpper(string(split[0][0])) + "." + strings.ToUpper(string(split[1][0]))
}
