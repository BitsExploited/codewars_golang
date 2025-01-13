package kata

func GetGrade(a,b,c int) rune {
    // Code here
  score := (a + b + c) / 3
  if score >= 0 && score < 60 {
        return 'F'
    } else if score >= 60 && score < 70 {
        return 'D'
    } else if score >= 70 && score < 80 {
        return 'C'
    } else if score >= 80 && score < 90 {
        return 'B'
    } else {
        return 'A'
    }
}
