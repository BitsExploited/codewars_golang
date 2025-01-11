package kata
import "strconv"

func Derive(coefficient, exponent int) string {
  coeff := coefficient * exponent
  expo := exponent - 1
  
  result := strconv.Itoa(coeff) + "x^" + strconv.Itoa(expo)
  
  return result
}
