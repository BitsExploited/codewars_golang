package kata

func Solution(a, b string) string {
    if len(a) > len(b) {
        return b + a + b
    } else if len(b) > len(a) {
        return a + b + a
    }
    // If the lengths are equal
    return a + b + a
}
