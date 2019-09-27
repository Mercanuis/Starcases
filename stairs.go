package main

import "fmt"

func main()  {
	fmt.Printf("For 1 step: %d\n", solveStairs(1))
	fmt.Printf("For 3 steps: %d\n", solveStairs(3))
}

func solveStairs(steps int) int {
	memo := make([]int, steps + 1)

	for i := range memo {
		memo[i] = 0
	}

	memo[0] = 1
	memo[1] = 1

	return solveStairsRecursive(steps, memo)
}

func solveStairsRecursive(stepsLeft int, memo []int) int {
	fmt.Printf("steps left: %d\n", stepsLeft)
	if stepsLeft < 0 {
		return 0;
	} else if memo[stepsLeft] > 0 {
		fmt.Printf("found a hit in memo -- %d\n", memo[stepsLeft])
		return memo[stepsLeft]
	} else {
		memo[stepsLeft] = solveStairsRecursive(stepsLeft - 3, memo) + solveStairsRecursive(stepsLeft -2, memo) + solveStairsRecursive(stepsLeft -1, memo)
		return memo[stepsLeft]
	}
}
