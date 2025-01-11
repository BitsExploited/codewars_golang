package kata

import "fmt"

func MultiTable(number int) string {
  var result string
  for i := 1; i <= 10; i++ {
    result += fmt.Sprintf("%d * %d = %d", i, number, i * number)
    
    if i < 10 {
      result += "\n"
    }
  }
  return result
}
