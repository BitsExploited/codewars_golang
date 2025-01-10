package kata

func Well(x []string) string {
    count := 0
    for _, i := range x {
        if i == "good" {
            count++
        }
    }
    if count == 0 {
        return "Fail!"
    } else if count == 1 || count == 2 {
        return "Publish!"
    } else {
        return "I smell a series!"
    }
}
