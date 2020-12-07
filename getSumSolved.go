package main

func getSum(array []int) int {
    sum := 0
    for i := 0; i < len(array); i++ {
        sum += array[i]
    }
    return sum

}
func main() {
    testCases := [][]int{{1, 2, 3, 4, 5}, {1, -2, -3, 4, 5}, {2}, {0}, {}}
    testCasesAnswers := []int{15, 5, 2, 0, 0}
    correct := true
    for i := 0; i < len(testCases); i++ {
        if getSum(testCases[i]) != testCasesAnswers[i] {
            println("Expected: ", testCasesAnswers[i], "Actual: ", getSum(testCases[i]))
            correct = false
        }
    }
    if correct {
        println("Success!")
    }
}