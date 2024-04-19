/*
Time complexity != Time taken, It is the rate at which time taken increases wrt input size
The rate at which the time, required to run a code, changes with respect to the input size.

Big O notation (upper-bound) is used to represent time complexity (worst) -> O(n)
Theta notation (average) -> θ(n)
Omega notation (best, lower-boundn) -> Ω(n)
    1. Compute worst case TC,
    2. Ignore constants,
    3. Ignore lower order terms

Space complexity = Auxiliary space (to solve problem) + Input space (to store input)

1 second = 10^8
*/

package main

import "fmt"

func timeComplexity() {
	var n int // Initialize n (you can set its value as needed)

	// First nested loop (O(n^2))
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d %d\n", i, j)
		}
	}

	// Second nested loop (O(n^2/2 + n/2))
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%d %d\n", i, j)
		}
	}
}
