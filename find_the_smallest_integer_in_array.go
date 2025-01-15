package kata

func SmallestIntegerFinder(numbers []int) int {
    if len(numbers) == 0 {
        
        return 0 
    }

    smallest := numbers[0] // Assuming the first number in the smallest number

    for _, num := range numbers {
        if num < smallest {
            smallest = num
        }
    }

    return smallest
}
