package kata

import "fmt"

func Greet(name string) string {
    s := fmt.Sprintf("Hello, %v how are you doing today?", name)
    return s

}
