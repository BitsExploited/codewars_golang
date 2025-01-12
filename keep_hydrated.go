package kata
import "math"
func Litres(time float64) int {
  var litres float64
  litres = time * 0.5
  return int(math.Floor(litres))  
}
