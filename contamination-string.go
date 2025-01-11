package kata


func Contamination(text, char string) string {
  if text == "" || char == "" {
return ""
}
  result := ""
  for range text {
result += char
}
  return result
}
